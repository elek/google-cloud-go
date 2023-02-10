// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: google/firestore/admin/v1/database.proto

package adminpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of the database.
// See https://cloud.google.com/datastore/docs/firestore-or-datastore for
// information about how to choose.
//
// Mode changes are only allowed if the database is empty.
type Database_DatabaseType int32

const (
	// The default value. This value is used if the database type is omitted.
	Database_DATABASE_TYPE_UNSPECIFIED Database_DatabaseType = 0
	// Firestore Native Mode
	Database_FIRESTORE_NATIVE Database_DatabaseType = 1
	// Firestore in Datastore Mode.
	Database_DATASTORE_MODE Database_DatabaseType = 2
)

// Enum value maps for Database_DatabaseType.
var (
	Database_DatabaseType_name = map[int32]string{
		0: "DATABASE_TYPE_UNSPECIFIED",
		1: "FIRESTORE_NATIVE",
		2: "DATASTORE_MODE",
	}
	Database_DatabaseType_value = map[string]int32{
		"DATABASE_TYPE_UNSPECIFIED": 0,
		"FIRESTORE_NATIVE":          1,
		"DATASTORE_MODE":            2,
	}
)

func (x Database_DatabaseType) Enum() *Database_DatabaseType {
	p := new(Database_DatabaseType)
	*p = x
	return p
}

func (x Database_DatabaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Database_DatabaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_database_proto_enumTypes[0].Descriptor()
}

func (Database_DatabaseType) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_database_proto_enumTypes[0]
}

func (x Database_DatabaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Database_DatabaseType.Descriptor instead.
func (Database_DatabaseType) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_database_proto_rawDescGZIP(), []int{0, 0}
}

// The type of concurrency control mode for transactions.
type Database_ConcurrencyMode int32

const (
	// Not used.
	Database_CONCURRENCY_MODE_UNSPECIFIED Database_ConcurrencyMode = 0
	// Use optimistic concurrency control by default. This mode is available
	// for Cloud Firestore databases.
	Database_OPTIMISTIC Database_ConcurrencyMode = 1
	// Use pessimistic concurrency control by default. This mode is available
	// for Cloud Firestore databases.
	//
	// This is the default setting for Cloud Firestore.
	Database_PESSIMISTIC Database_ConcurrencyMode = 2
	// Use optimistic concurrency control with entity groups by default.
	//
	// This is the only available mode for Cloud Datastore.
	//
	// This mode is also available for Cloud Firestore with Datastore Mode but
	// is not recommended.
	Database_OPTIMISTIC_WITH_ENTITY_GROUPS Database_ConcurrencyMode = 3
)

// Enum value maps for Database_ConcurrencyMode.
var (
	Database_ConcurrencyMode_name = map[int32]string{
		0: "CONCURRENCY_MODE_UNSPECIFIED",
		1: "OPTIMISTIC",
		2: "PESSIMISTIC",
		3: "OPTIMISTIC_WITH_ENTITY_GROUPS",
	}
	Database_ConcurrencyMode_value = map[string]int32{
		"CONCURRENCY_MODE_UNSPECIFIED":  0,
		"OPTIMISTIC":                    1,
		"PESSIMISTIC":                   2,
		"OPTIMISTIC_WITH_ENTITY_GROUPS": 3,
	}
)

func (x Database_ConcurrencyMode) Enum() *Database_ConcurrencyMode {
	p := new(Database_ConcurrencyMode)
	*p = x
	return p
}

func (x Database_ConcurrencyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Database_ConcurrencyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_database_proto_enumTypes[1].Descriptor()
}

func (Database_ConcurrencyMode) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_database_proto_enumTypes[1]
}

func (x Database_ConcurrencyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Database_ConcurrencyMode.Descriptor instead.
func (Database_ConcurrencyMode) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_database_proto_rawDescGZIP(), []int{0, 1}
}

// The type of App Engine integration mode.
type Database_AppEngineIntegrationMode int32

const (
	// Not used.
	Database_APP_ENGINE_INTEGRATION_MODE_UNSPECIFIED Database_AppEngineIntegrationMode = 0
	// If an App Engine application exists in the same region as this database,
	// App Engine configuration will impact this database. This includes
	// disabling of the application & database, as well as disabling writes to
	// the database.
	Database_ENABLED Database_AppEngineIntegrationMode = 1
	// Appengine has no affect on the ability of this database to serve
	// requests.
	Database_DISABLED Database_AppEngineIntegrationMode = 2
)

// Enum value maps for Database_AppEngineIntegrationMode.
var (
	Database_AppEngineIntegrationMode_name = map[int32]string{
		0: "APP_ENGINE_INTEGRATION_MODE_UNSPECIFIED",
		1: "ENABLED",
		2: "DISABLED",
	}
	Database_AppEngineIntegrationMode_value = map[string]int32{
		"APP_ENGINE_INTEGRATION_MODE_UNSPECIFIED": 0,
		"ENABLED":  1,
		"DISABLED": 2,
	}
)

func (x Database_AppEngineIntegrationMode) Enum() *Database_AppEngineIntegrationMode {
	p := new(Database_AppEngineIntegrationMode)
	*p = x
	return p
}

func (x Database_AppEngineIntegrationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Database_AppEngineIntegrationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_database_proto_enumTypes[2].Descriptor()
}

func (Database_AppEngineIntegrationMode) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_database_proto_enumTypes[2]
}

func (x Database_AppEngineIntegrationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Database_AppEngineIntegrationMode.Descriptor instead.
func (Database_AppEngineIntegrationMode) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_database_proto_rawDescGZIP(), []int{0, 2}
}

// A Cloud Firestore Database.
// Currently only one database is allowed per cloud project; this database
// must have a `database_id` of '(default)'.
type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Database.
	// Format: `projects/{project}/databases/{database}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The location of the database. Available databases are listed at
	// https://cloud.google.com/firestore/docs/locations.
	LocationId string `protobuf:"bytes,9,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// The type of the database.
	// See https://cloud.google.com/datastore/docs/firestore-or-datastore for
	// information about how to choose.
	Type Database_DatabaseType `protobuf:"varint,10,opt,name=type,proto3,enum=google.firestore.admin.v1.Database_DatabaseType" json:"type,omitempty"`
	// The concurrency control mode to use for this database.
	ConcurrencyMode Database_ConcurrencyMode `protobuf:"varint,15,opt,name=concurrency_mode,json=concurrencyMode,proto3,enum=google.firestore.admin.v1.Database_ConcurrencyMode" json:"concurrency_mode,omitempty"`
	// The App Engine integration mode to use for this database.
	AppEngineIntegrationMode Database_AppEngineIntegrationMode `protobuf:"varint,19,opt,name=app_engine_integration_mode,json=appEngineIntegrationMode,proto3,enum=google.firestore.admin.v1.Database_AppEngineIntegrationMode" json:"app_engine_integration_mode,omitempty"`
	// Output only. The key_prefix for this database. This key_prefix is used, in combination
	// with the project id ("<key prefix>~<project id>") to construct the
	// application id that is returned from the Cloud Datastore APIs in Google App
	// Engine first generation runtimes.
	//
	// This value may be empty in which case the appid to use for URL-encoded keys
	// is the project_id (eg: foo instead of v~foo).
	KeyPrefix string `protobuf:"bytes,20,opt,name=key_prefix,json=keyPrefix,proto3" json:"key_prefix,omitempty"`
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag string `protobuf:"bytes,99,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *Database) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Database) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Database) GetType() Database_DatabaseType {
	if x != nil {
		return x.Type
	}
	return Database_DATABASE_TYPE_UNSPECIFIED
}

func (x *Database) GetConcurrencyMode() Database_ConcurrencyMode {
	if x != nil {
		return x.ConcurrencyMode
	}
	return Database_CONCURRENCY_MODE_UNSPECIFIED
}

func (x *Database) GetAppEngineIntegrationMode() Database_AppEngineIntegrationMode {
	if x != nil {
		return x.AppEngineIntegrationMode
	}
	return Database_APP_ENGINE_INTEGRATION_MODE_UNSPECIFIED
}

func (x *Database) GetKeyPrefix() string {
	if x != nil {
		return x.KeyPrefix
	}
	return ""
}

func (x *Database) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

var File_google_firestore_admin_v1_database_proto protoreflect.FileDescriptor

var file_google_firestore_admin_v1_database_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa4, 0x06, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5e, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x7b, 0x0a, 0x1b, 0x61, 0x70, 0x70,
	0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x18, 0x61, 0x70,
	0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x09, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74,
	0x61, 0x67, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x57,
	0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x19, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x46, 0x49, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x53, 0x54, 0x4f, 0x52, 0x45,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x22, 0x77, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f,
	0x4e, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x54, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x45, 0x53, 0x53, 0x49, 0x4d, 0x49, 0x53, 0x54, 0x49, 0x43, 0x10, 0x02, 0x12, 0x21, 0x0a,
	0x1d, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x54, 0x49, 0x43, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x03,
	0x22, 0x62, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x27,
	0x41, 0x50, 0x50, 0x5f, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x3a, 0x52, 0xea, 0x41, 0x4f, 0x0a, 0x21, 0x66, 0x69, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x27, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x7d, 0x52, 0x01, 0x01, 0x42, 0xe1, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x04, 0x47, 0x43,
	0x46, 0x53, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firestore_admin_v1_database_proto_rawDescOnce sync.Once
	file_google_firestore_admin_v1_database_proto_rawDescData = file_google_firestore_admin_v1_database_proto_rawDesc
)

func file_google_firestore_admin_v1_database_proto_rawDescGZIP() []byte {
	file_google_firestore_admin_v1_database_proto_rawDescOnce.Do(func() {
		file_google_firestore_admin_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_admin_v1_database_proto_rawDescData)
	})
	return file_google_firestore_admin_v1_database_proto_rawDescData
}

var file_google_firestore_admin_v1_database_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_firestore_admin_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_firestore_admin_v1_database_proto_goTypes = []interface{}{
	(Database_DatabaseType)(0),             // 0: google.firestore.admin.v1.Database.DatabaseType
	(Database_ConcurrencyMode)(0),          // 1: google.firestore.admin.v1.Database.ConcurrencyMode
	(Database_AppEngineIntegrationMode)(0), // 2: google.firestore.admin.v1.Database.AppEngineIntegrationMode
	(*Database)(nil),                       // 3: google.firestore.admin.v1.Database
}
var file_google_firestore_admin_v1_database_proto_depIdxs = []int32{
	0, // 0: google.firestore.admin.v1.Database.type:type_name -> google.firestore.admin.v1.Database.DatabaseType
	1, // 1: google.firestore.admin.v1.Database.concurrency_mode:type_name -> google.firestore.admin.v1.Database.ConcurrencyMode
	2, // 2: google.firestore.admin.v1.Database.app_engine_integration_mode:type_name -> google.firestore.admin.v1.Database.AppEngineIntegrationMode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_firestore_admin_v1_database_proto_init() }
func file_google_firestore_admin_v1_database_proto_init() {
	if File_google_firestore_admin_v1_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_firestore_admin_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Database); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_admin_v1_database_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_admin_v1_database_proto_goTypes,
		DependencyIndexes: file_google_firestore_admin_v1_database_proto_depIdxs,
		EnumInfos:         file_google_firestore_admin_v1_database_proto_enumTypes,
		MessageInfos:      file_google_firestore_admin_v1_database_proto_msgTypes,
	}.Build()
	File_google_firestore_admin_v1_database_proto = out.File
	file_google_firestore_admin_v1_database_proto_rawDesc = nil
	file_google_firestore_admin_v1_database_proto_goTypes = nil
	file_google_firestore_admin_v1_database_proto_depIdxs = nil
}
