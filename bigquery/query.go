// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bigquery

import (
	"golang.org/x/net/context"
	bq "google.golang.org/api/bigquery/v2"
)

// QueryConfig holds the configuration for a query job.
type QueryConfig struct {
	// JobID is the ID to use for the query job. If this field is empty, a job ID
	// will be automatically created.
	JobID string

	// Dst is the table into which the results of the query will be written.
	// If this field is nil, a temporary table will be created.
	Dst *Table

	// The query to execute. See https://cloud.google.com/bigquery/query-reference for details.
	Q string

	// DefaultProjectID and DefaultDatasetID specify the dataset to use for unqualified table names in the query.
	// If DefaultProjectID is set, DefaultDatasetID must also be set.
	DefaultProjectID string
	DefaultDatasetID string

	// TableDefinitions describes data sources outside of BigQuery.
	// The map keys may be used as table names in the query string.
	TableDefinitions map[string]ExternalData

	// CreateDisposition specifies the circumstances under which the destination table will be created.
	// The default is CreateIfNeeded.
	CreateDisposition TableCreateDisposition

	// WriteDisposition specifies how existing data in the destination table is treated.
	// The default is WriteAppend.
	WriteDisposition TableWriteDisposition

	// DisableQueryCache prevents results being fetched from the query cache.
	// If this field is false, results are fetched from the cache if they are available.
	// The query cache is a best-effort cache that is flushed whenever tables in the query are modified.
	// Cached results are only available when TableID is unspecified in the query's destination Table.
	// For more information, see https://cloud.google.com/bigquery/querying-data#querycaching
	DisableQueryCache bool

	// DisableFlattenedResults prevents results being flattened.
	// If this field is false, results from nested and repeated fields are flattened.
	// DisableFlattenedResults implies AllowLargeResults
	// For more information, see https://cloud.google.com/bigquery/docs/data#nested
	DisableFlattenedResults bool

	// AllowLargeResults allows the query to produce arbitrarily large result tables.
	// The destination must be a table.
	// When using this option, queries will take longer to execute, even if the result set is small.
	// For additional limitations, see https://cloud.google.com/bigquery/querying-data#largequeryresults
	AllowLargeResults bool

	// Priority species the priority with which to schedule the query.
	// The default priority is InteractivePriority.
	// For more information, see https://cloud.google.com/bigquery/querying-data#batchqueries
	Priority QueryPriority

	// MaxBillingTier sets the maximum billing tier for a Query.
	// Queries that have resource usage beyond this tier will fail (without
	// incurring a charge). If this field is zero, the project default will be used.
	MaxBillingTier int

	// MaxBytesBilled limits the number of bytes billed for
	// this job.  Queries that would exceed this limit will fail (without incurring
	// a charge).
	// If this field is less than 1, the project default will be
	// used.
	MaxBytesBilled int64

	// UseStandardSQL causes the query to use standard SQL.
	// The default is false (using legacy SQL).
	UseStandardSQL bool
}

// QueryPriority species a priority with which a query is to be executed.
type QueryPriority string

const (
	BatchPriority       QueryPriority = "BATCH"
	InteractivePriority QueryPriority = "INTERACTIVE"
)

// A Query queries data from a BigQuery table. Use Client.Query to create a Query.
type Query struct {
	client *Client
	QueryConfig

	// The query to execute. See https://cloud.google.com/bigquery/query-reference for details.
	// Deprecated: use QueryConfig.Q instead.
	Q string

	// DefaultProjectID and DefaultDatasetID specify the dataset to use for unqualified table names in the query.
	// If DefaultProjectID is set, DefaultDatasetID must also be set.
	DefaultProjectID string // Deprecated: use QueryConfig.DefaultProjectID instead.
	DefaultDatasetID string // Deprecated: use QueryConfig.DefaultDatasetID instead.
}

func (q *QueryConfig) implementsSource()     {}
func (q *QueryConfig) implementsReadSource() {}

func (q *Query) implementsSource()     {}
func (q *Query) implementsReadSource() {}

func (q *QueryConfig) customizeQuerySrc(conf *bq.JobConfigurationQuery) {
	conf.Query = q.Q

	if len(q.TableDefinitions) > 0 {
		conf.TableDefinitions = make(map[string]bq.ExternalDataConfiguration)
	}
	for name, data := range q.TableDefinitions {
		conf.TableDefinitions[name] = data.externalDataConfig()
	}

	if q.DefaultProjectID != "" || q.DefaultDatasetID != "" {
		conf.DefaultDataset = &bq.DatasetReference{
			DatasetId: q.DefaultDatasetID,
			ProjectId: q.DefaultProjectID,
		}
	}

	if tier := int64(q.MaxBillingTier); tier > 0 {
		conf.MaximumBillingTier = &tier
	}
	conf.CreateDisposition = string(q.CreateDisposition)
	conf.WriteDisposition = string(q.WriteDisposition)
	conf.AllowLargeResults = q.AllowLargeResults
	conf.Priority = string(q.Priority)

	f := false
	if q.DisableQueryCache {
		conf.UseQueryCache = &f
	}
	if q.DisableFlattenedResults {
		conf.FlattenResults = &f
		// DisableFlattenResults implies AllowLargeResults.
		conf.AllowLargeResults = true
	}
	if q.MaxBytesBilled >= 1 {
		conf.MaximumBytesBilled = q.MaxBytesBilled
	}
	if q.UseStandardSQL {
		conf.UseLegacySql = false
		conf.ForceSendFields = append(conf.ForceSendFields, "UseLegacySql")
	}

	if q.Dst != nil {
		q.Dst.customizeQueryDst(conf)
	}
}

// Query creates a query with string q.
// The returned Query may optionally be further configured before its Run method is called.
func (c *Client) Query(q string) *Query {
	return &Query{
		client:      c,
		QueryConfig: QueryConfig{Q: q},
	}
}

// Run initiates a query job.
func (q *Query) Run(ctx context.Context) (*Job, error) {
	job := &bq.Job{
		Configuration: &bq.JobConfiguration{
			Query: &bq.JobConfigurationQuery{},
		},
	}
	setJobRef(job, q.JobID, q.client.projectID)

	q.QueryConfig.customizeQuerySrc(job.Configuration.Query)

	// For compatability, allow some legacy fields to be set directly on the query.
	// Even though Query.Run is new, it is called by Query.Read, which may have non-empty deprecated fields.
	// TODO(jba): delete this code when deleting Client.Copy.
	conf := job.Configuration.Query
	if q.Q != "" {
		conf.Query = q.Q
	}
	if q.DefaultProjectID != "" || q.DefaultDatasetID != "" {
		conf.DefaultDataset = &bq.DatasetReference{
			DatasetId: q.DefaultDatasetID,
			ProjectId: q.DefaultProjectID,
		}
	}
	// end of compatability code.

	j, err := q.client.service.insertJob(ctx, job, q.client.projectID)
	if err != nil {
		return nil, err
	}
	j.isQuery = true
	return j, nil
}
