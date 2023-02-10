// Copyright 2020 Google LLC
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
// source: google/devtools/artifactregistry/v1beta2/file.proto

package artifactregistrypb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The algorithm used to compute the hash.
type Hash_HashType int32

const (
	// Unspecified.
	Hash_HASH_TYPE_UNSPECIFIED Hash_HashType = 0
	// SHA256 hash.
	Hash_SHA256 Hash_HashType = 1
	// MD5 hash.
	Hash_MD5 Hash_HashType = 2
)

// Enum value maps for Hash_HashType.
var (
	Hash_HashType_name = map[int32]string{
		0: "HASH_TYPE_UNSPECIFIED",
		1: "SHA256",
		2: "MD5",
	}
	Hash_HashType_value = map[string]int32{
		"HASH_TYPE_UNSPECIFIED": 0,
		"SHA256":                1,
		"MD5":                   2,
	}
)

func (x Hash_HashType) Enum() *Hash_HashType {
	p := new(Hash_HashType)
	*p = x
	return p
}

func (x Hash_HashType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Hash_HashType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_artifactregistry_v1beta2_file_proto_enumTypes[0].Descriptor()
}

func (Hash_HashType) Type() protoreflect.EnumType {
	return &file_google_devtools_artifactregistry_v1beta2_file_proto_enumTypes[0]
}

func (x Hash_HashType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Hash_HashType.Descriptor instead.
func (Hash_HashType) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescGZIP(), []int{0, 0}
}

// A hash of file content.
type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The algorithm used to compute the hash value.
	Type Hash_HashType `protobuf:"varint,1,opt,name=type,proto3,enum=google.devtools.artifactregistry.v1beta2.Hash_HashType" json:"type,omitempty"`
	// The hash value.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescGZIP(), []int{0}
}

func (x *Hash) GetType() Hash_HashType {
	if x != nil {
		return x.Type
	}
	return Hash_HASH_TYPE_UNSPECIFIED
}

func (x *Hash) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Files store content that is potentially associated with Packages or Versions.
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the file, for example:
	// "projects/p1/locations/us-central1/repositories/repo1/files/a%2Fb%2Fc.txt".
	// If the file ID part contains slashes, they are escaped.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The size of the File in bytes.
	SizeBytes int64 `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	// The hashes of the file content.
	Hashes []*Hash `protobuf:"bytes,4,rep,name=hashes,proto3" json:"hashes,omitempty"`
	// The time when the File was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time when the File was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The name of the Package or Version that owns this file, if any.
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *File) GetHashes() []*Hash {
	if x != nil {
		return x.Hashes
	}
	return nil
}

func (x *File) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *File) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *File) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// The request to list files.
type ListFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the repository whose files will be listed. For example:
	// "projects/p1/locations/us-central1/repositories/repo1
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// An expression for filtering the results of the request. Filter rules are
	// case insensitive. The fields eligible for filtering are:
	//
	//   - `name`
	//
	//   - `owner`
	//
	//     An example of using a filter:
	//
	//   - `name="projects/p1/locations/us-central1/repositories/repo1/files/a/b/*"` --> Files with an
	//     ID starting with "a/b/".
	//
	//   - `owner="projects/p1/locations/us-central1/repositories/repo1/packages/pkg1/versions/1.0"` -->
	//     Files owned by the version `1.0` in package `pkg1`.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// The maximum number of files to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous list request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListFilesRequest) Reset() {
	*x = ListFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesRequest) ProtoMessage() {}

func (x *ListFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesRequest.ProtoReflect.Descriptor instead.
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescGZIP(), []int{2}
}

func (x *ListFilesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListFilesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListFilesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFilesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response from listing files.
type ListFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The files returned.
	Files []*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// The token to retrieve the next page of files, or empty if there are no
	// more files to return.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListFilesResponse) Reset() {
	*x = ListFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesResponse) ProtoMessage() {}

func (x *ListFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesResponse.ProtoReflect.Descriptor instead.
func (*ListFilesResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescGZIP(), []int{3}
}

func (x *ListFilesResponse) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *ListFilesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request to retrieve a file.
type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the file to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescGZIP(), []int{4}
}

func (x *GetFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_devtools_artifactregistry_v1beta2_file_proto protoreflect.FileDescriptor

var file_google_devtools_artifactregistry_v1beta2_file_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3a, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x44,
	0x35, 0x10, 0x02, 0x22, 0x8c, 0x03, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x46, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x3a, 0x79, 0xea, 0x41, 0x76, 0x0a, 0x24, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x4e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x7d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6c,
	0x65, 0x7d, 0x22, 0x7e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x90, 0x02, 0x0a,
	0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x42, 0x09, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0xaa, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x32, 0xca, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x32, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescOnce sync.Once
	file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescData = file_google_devtools_artifactregistry_v1beta2_file_proto_rawDesc
)

func file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescGZIP() []byte {
	file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescOnce.Do(func() {
		file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescData)
	})
	return file_google_devtools_artifactregistry_v1beta2_file_proto_rawDescData
}

var file_google_devtools_artifactregistry_v1beta2_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_devtools_artifactregistry_v1beta2_file_proto_goTypes = []interface{}{
	(Hash_HashType)(0),            // 0: google.devtools.artifactregistry.v1beta2.Hash.HashType
	(*Hash)(nil),                  // 1: google.devtools.artifactregistry.v1beta2.Hash
	(*File)(nil),                  // 2: google.devtools.artifactregistry.v1beta2.File
	(*ListFilesRequest)(nil),      // 3: google.devtools.artifactregistry.v1beta2.ListFilesRequest
	(*ListFilesResponse)(nil),     // 4: google.devtools.artifactregistry.v1beta2.ListFilesResponse
	(*GetFileRequest)(nil),        // 5: google.devtools.artifactregistry.v1beta2.GetFileRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_google_devtools_artifactregistry_v1beta2_file_proto_depIdxs = []int32{
	0, // 0: google.devtools.artifactregistry.v1beta2.Hash.type:type_name -> google.devtools.artifactregistry.v1beta2.Hash.HashType
	1, // 1: google.devtools.artifactregistry.v1beta2.File.hashes:type_name -> google.devtools.artifactregistry.v1beta2.Hash
	6, // 2: google.devtools.artifactregistry.v1beta2.File.create_time:type_name -> google.protobuf.Timestamp
	6, // 3: google.devtools.artifactregistry.v1beta2.File.update_time:type_name -> google.protobuf.Timestamp
	2, // 4: google.devtools.artifactregistry.v1beta2.ListFilesResponse.files:type_name -> google.devtools.artifactregistry.v1beta2.File
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_devtools_artifactregistry_v1beta2_file_proto_init() }
func file_google_devtools_artifactregistry_v1beta2_file_proto_init() {
	if File_google_devtools_artifactregistry_v1beta2_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesRequest); i {
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
		file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesResponse); i {
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
		file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
			RawDescriptor: file_google_devtools_artifactregistry_v1beta2_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_artifactregistry_v1beta2_file_proto_goTypes,
		DependencyIndexes: file_google_devtools_artifactregistry_v1beta2_file_proto_depIdxs,
		EnumInfos:         file_google_devtools_artifactregistry_v1beta2_file_proto_enumTypes,
		MessageInfos:      file_google_devtools_artifactregistry_v1beta2_file_proto_msgTypes,
	}.Build()
	File_google_devtools_artifactregistry_v1beta2_file_proto = out.File
	file_google_devtools_artifactregistry_v1beta2_file_proto_rawDesc = nil
	file_google_devtools_artifactregistry_v1beta2_file_proto_goTypes = nil
	file_google_devtools_artifactregistry_v1beta2_file_proto_depIdxs = nil
}
