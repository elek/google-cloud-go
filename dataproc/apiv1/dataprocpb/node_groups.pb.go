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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: google/cloud/dataproc/v1/node_groups.proto

package dataprocpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A request to create a node group.
type CreateNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource where this node group will be created.
	// Format: `projects/{project}/regions/{region}/clusters/{cluster}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The node group to create.
	NodeGroup *NodeGroup `protobuf:"bytes,2,opt,name=node_group,json=nodeGroup,proto3" json:"node_group,omitempty"`
	// Optional. An optional node group ID. Generated if not specified.
	//
	// The ID must contain only letters (a-z, A-Z), numbers (0-9),
	// underscores (_), and hyphens (-). Cannot begin or end with underscore
	// or hyphen. Must consist of from 3 to 33 characters.
	NodeGroupId string `protobuf:"bytes,4,opt,name=node_group_id,json=nodeGroupId,proto3" json:"node_group_id,omitempty"`
	// Optional. A unique ID used to identify the request. If the server receives
	// two
	// [CreateNodeGroupRequest](https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#google.cloud.dataproc.v1.CreateNodeGroupRequests)
	// with the same ID, the second request is ignored and the
	// first [google.longrunning.Operation][google.longrunning.Operation] created
	// and stored in the backend is returned.
	//
	// Recommendation: Set this value to a
	// [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier).
	//
	// The ID must contain only letters (a-z, A-Z), numbers (0-9),
	// underscores (_), and hyphens (-). The maximum length is 40 characters.
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *CreateNodeGroupRequest) Reset() {
	*x = CreateNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeGroupRequest) ProtoMessage() {}

func (x *CreateNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_node_groups_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeGroupRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateNodeGroupRequest) GetNodeGroup() *NodeGroup {
	if x != nil {
		return x.NodeGroup
	}
	return nil
}

func (x *CreateNodeGroupRequest) GetNodeGroupId() string {
	if x != nil {
		return x.NodeGroupId
	}
	return ""
}

func (x *CreateNodeGroupRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

// A request to resize a node group.
type ResizeNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the node group to resize.
	// Format:
	// `projects/{project}/regions/{region}/clusters/{cluster}/nodeGroups/{nodeGroup}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The number of running instances for the node group to maintain.
	// The group adds or removes instances to maintain the number of instances
	// specified by this parameter.
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// Optional. A unique ID used to identify the request. If the server receives
	// two
	// [ResizeNodeGroupRequest](https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#google.cloud.dataproc.v1.ResizeNodeGroupRequests)
	// with the same ID, the second request is ignored and the
	// first [google.longrunning.Operation][google.longrunning.Operation] created
	// and stored in the backend is returned.
	//
	// Recommendation: Set this value to a
	// [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier).
	//
	// The ID must contain only letters (a-z, A-Z), numbers (0-9),
	// underscores (_), and hyphens (-). The maximum length is 40 characters.
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Optional. Timeout for graceful YARN decomissioning. [Graceful
	// decommissioning]
	// (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/scaling-clusters#graceful_decommissioning)
	// allows the removal of nodes from the Compute Engine node group
	// without interrupting jobs in progress. This timeout specifies how long to
	// wait for jobs in progress to finish before forcefully removing nodes (and
	// potentially interrupting jobs). Default timeout is 0 (for forceful
	// decommission), and the maximum allowed timeout is 1 day. (see JSON
	// representation of
	// [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//
	// Only supported on Dataproc image versions 1.2 and higher.
	GracefulDecommissionTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=graceful_decommission_timeout,json=gracefulDecommissionTimeout,proto3" json:"graceful_decommission_timeout,omitempty"`
}

func (x *ResizeNodeGroupRequest) Reset() {
	*x = ResizeNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeNodeGroupRequest) ProtoMessage() {}

func (x *ResizeNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*ResizeNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_node_groups_proto_rawDescGZIP(), []int{1}
}

func (x *ResizeNodeGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResizeNodeGroupRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ResizeNodeGroupRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ResizeNodeGroupRequest) GetGracefulDecommissionTimeout() *durationpb.Duration {
	if x != nil {
		return x.GracefulDecommissionTimeout
	}
	return nil
}

// A request to get a node group .
type GetNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the node group to retrieve.
	// Format:
	// `projects/{project}/regions/{region}/clusters/{cluster}/nodeGroups/{nodeGroup}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetNodeGroupRequest) Reset() {
	*x = GetNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeGroupRequest) ProtoMessage() {}

func (x *GetNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*GetNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_node_groups_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodeGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_dataproc_v1_node_groups_proto protoreflect.FileDescriptor

var file_google_cloud_dataproc_v1_node_groups_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x12, 0x21, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x72, 0x6f, 0x63, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x27,
	0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xd2, 0x01, 0x0a, 0x16,
	0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x62, 0x0a, 0x1d,
	0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x1b, 0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x44, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0x54, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa7, 0x06, 0x0a, 0x13, 0x4e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x95,
	0x02, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f,
	0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xb0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x22, 0x37, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x3a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0xda, 0x41, 0x1f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2c, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0xca, 0x41, 0x40, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0xfd, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x43, 0x22, 0x3e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65,
	0x73, 0x69, 0x7a, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x73,
	0x69, 0x7a, 0x65, 0xca, 0x41, 0x40, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0xaa, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x46, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x39, 0x12, 0x37, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x1a, 0x4b, 0xca, 0x41, 0x17, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x42, 0xcd, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76,
	0x31, 0x42, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x63, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x70, 0x62, 0xea, 0x41,
	0x5f, 0x0a, 0x25, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dataproc_v1_node_groups_proto_rawDescOnce sync.Once
	file_google_cloud_dataproc_v1_node_groups_proto_rawDescData = file_google_cloud_dataproc_v1_node_groups_proto_rawDesc
)

func file_google_cloud_dataproc_v1_node_groups_proto_rawDescGZIP() []byte {
	file_google_cloud_dataproc_v1_node_groups_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataproc_v1_node_groups_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataproc_v1_node_groups_proto_rawDescData)
	})
	return file_google_cloud_dataproc_v1_node_groups_proto_rawDescData
}

var file_google_cloud_dataproc_v1_node_groups_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_dataproc_v1_node_groups_proto_goTypes = []interface{}{
	(*CreateNodeGroupRequest)(nil), // 0: google.cloud.dataproc.v1.CreateNodeGroupRequest
	(*ResizeNodeGroupRequest)(nil), // 1: google.cloud.dataproc.v1.ResizeNodeGroupRequest
	(*GetNodeGroupRequest)(nil),    // 2: google.cloud.dataproc.v1.GetNodeGroupRequest
	(*NodeGroup)(nil),              // 3: google.cloud.dataproc.v1.NodeGroup
	(*durationpb.Duration)(nil),    // 4: google.protobuf.Duration
	(*longrunning.Operation)(nil),  // 5: google.longrunning.Operation
}
var file_google_cloud_dataproc_v1_node_groups_proto_depIdxs = []int32{
	3, // 0: google.cloud.dataproc.v1.CreateNodeGroupRequest.node_group:type_name -> google.cloud.dataproc.v1.NodeGroup
	4, // 1: google.cloud.dataproc.v1.ResizeNodeGroupRequest.graceful_decommission_timeout:type_name -> google.protobuf.Duration
	0, // 2: google.cloud.dataproc.v1.NodeGroupController.CreateNodeGroup:input_type -> google.cloud.dataproc.v1.CreateNodeGroupRequest
	1, // 3: google.cloud.dataproc.v1.NodeGroupController.ResizeNodeGroup:input_type -> google.cloud.dataproc.v1.ResizeNodeGroupRequest
	2, // 4: google.cloud.dataproc.v1.NodeGroupController.GetNodeGroup:input_type -> google.cloud.dataproc.v1.GetNodeGroupRequest
	5, // 5: google.cloud.dataproc.v1.NodeGroupController.CreateNodeGroup:output_type -> google.longrunning.Operation
	5, // 6: google.cloud.dataproc.v1.NodeGroupController.ResizeNodeGroup:output_type -> google.longrunning.Operation
	3, // 7: google.cloud.dataproc.v1.NodeGroupController.GetNodeGroup:output_type -> google.cloud.dataproc.v1.NodeGroup
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_dataproc_v1_node_groups_proto_init() }
func file_google_cloud_dataproc_v1_node_groups_proto_init() {
	if File_google_cloud_dataproc_v1_node_groups_proto != nil {
		return
	}
	file_google_cloud_dataproc_v1_clusters_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeGroupRequest); i {
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
		file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeNodeGroupRequest); i {
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
		file_google_cloud_dataproc_v1_node_groups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeGroupRequest); i {
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
			RawDescriptor: file_google_cloud_dataproc_v1_node_groups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_dataproc_v1_node_groups_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataproc_v1_node_groups_proto_depIdxs,
		MessageInfos:      file_google_cloud_dataproc_v1_node_groups_proto_msgTypes,
	}.Build()
	File_google_cloud_dataproc_v1_node_groups_proto = out.File
	file_google_cloud_dataproc_v1_node_groups_proto_rawDesc = nil
	file_google_cloud_dataproc_v1_node_groups_proto_goTypes = nil
	file_google_cloud_dataproc_v1_node_groups_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeGroupControllerClient is the client API for NodeGroupController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeGroupControllerClient interface {
	// Creates a node group in a cluster. The returned
	// [Operation.metadata][google.longrunning.Operation.metadata] is
	// [NodeGroupOperationMetadata](https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#nodegroupoperationmetadata).
	CreateNodeGroup(ctx context.Context, in *CreateNodeGroupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Resizes a node group in a cluster. The returned
	// [Operation.metadata][google.longrunning.Operation.metadata] is
	// [NodeGroupOperationMetadata](https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#nodegroupoperationmetadata).
	ResizeNodeGroup(ctx context.Context, in *ResizeNodeGroupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Gets the resource representation for a node group in a
	// cluster.
	GetNodeGroup(ctx context.Context, in *GetNodeGroupRequest, opts ...grpc.CallOption) (*NodeGroup, error)
}

type nodeGroupControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeGroupControllerClient(cc grpc.ClientConnInterface) NodeGroupControllerClient {
	return &nodeGroupControllerClient{cc}
}

func (c *nodeGroupControllerClient) CreateNodeGroup(ctx context.Context, in *CreateNodeGroupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.dataproc.v1.NodeGroupController/CreateNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupControllerClient) ResizeNodeGroup(ctx context.Context, in *ResizeNodeGroupRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.dataproc.v1.NodeGroupController/ResizeNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupControllerClient) GetNodeGroup(ctx context.Context, in *GetNodeGroupRequest, opts ...grpc.CallOption) (*NodeGroup, error) {
	out := new(NodeGroup)
	err := c.cc.Invoke(ctx, "/google.cloud.dataproc.v1.NodeGroupController/GetNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeGroupControllerServer is the server API for NodeGroupController service.
type NodeGroupControllerServer interface {
	// Creates a node group in a cluster. The returned
	// [Operation.metadata][google.longrunning.Operation.metadata] is
	// [NodeGroupOperationMetadata](https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#nodegroupoperationmetadata).
	CreateNodeGroup(context.Context, *CreateNodeGroupRequest) (*longrunning.Operation, error)
	// Resizes a node group in a cluster. The returned
	// [Operation.metadata][google.longrunning.Operation.metadata] is
	// [NodeGroupOperationMetadata](https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#nodegroupoperationmetadata).
	ResizeNodeGroup(context.Context, *ResizeNodeGroupRequest) (*longrunning.Operation, error)
	// Gets the resource representation for a node group in a
	// cluster.
	GetNodeGroup(context.Context, *GetNodeGroupRequest) (*NodeGroup, error)
}

// UnimplementedNodeGroupControllerServer can be embedded to have forward compatible implementations.
type UnimplementedNodeGroupControllerServer struct {
}

func (*UnimplementedNodeGroupControllerServer) CreateNodeGroup(context.Context, *CreateNodeGroupRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeGroup not implemented")
}
func (*UnimplementedNodeGroupControllerServer) ResizeNodeGroup(context.Context, *ResizeNodeGroupRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeNodeGroup not implemented")
}
func (*UnimplementedNodeGroupControllerServer) GetNodeGroup(context.Context, *GetNodeGroupRequest) (*NodeGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeGroup not implemented")
}

func RegisterNodeGroupControllerServer(s *grpc.Server, srv NodeGroupControllerServer) {
	s.RegisterService(&_NodeGroupController_serviceDesc, srv)
}

func _NodeGroupController_CreateNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupControllerServer).CreateNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dataproc.v1.NodeGroupController/CreateNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupControllerServer).CreateNodeGroup(ctx, req.(*CreateNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupController_ResizeNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupControllerServer).ResizeNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dataproc.v1.NodeGroupController/ResizeNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupControllerServer).ResizeNodeGroup(ctx, req.(*ResizeNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupController_GetNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupControllerServer).GetNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dataproc.v1.NodeGroupController/GetNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupControllerServer).GetNodeGroup(ctx, req.(*GetNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeGroupController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dataproc.v1.NodeGroupController",
	HandlerType: (*NodeGroupControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNodeGroup",
			Handler:    _NodeGroupController_CreateNodeGroup_Handler,
		},
		{
			MethodName: "ResizeNodeGroup",
			Handler:    _NodeGroupController_ResizeNodeGroup_Handler,
		},
		{
			MethodName: "GetNodeGroup",
			Handler:    _NodeGroupController_GetNodeGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dataproc/v1/node_groups.proto",
}
