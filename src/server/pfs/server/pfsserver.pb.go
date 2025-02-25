// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: server/pfs/server/pfsserver.proto

package server

import (
	index "github.com/pachyderm/pachyderm/v2/src/internal/storage/fileset/index"
	pfs "github.com/pachyderm/pachyderm/v2/src/pfs"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShardTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs    []string   `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	PathRange *PathRange `protobuf:"bytes,2,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
}

func (x *ShardTask) Reset() {
	*x = ShardTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardTask) ProtoMessage() {}

func (x *ShardTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardTask.ProtoReflect.Descriptor instead.
func (*ShardTask) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{0}
}

func (x *ShardTask) GetInputs() []string {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ShardTask) GetPathRange() *PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

type ShardTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompactTasks []*CompactTask `protobuf:"bytes,1,rep,name=compact_tasks,json=compactTasks,proto3" json:"compact_tasks,omitempty"`
}

func (x *ShardTaskResult) Reset() {
	*x = ShardTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardTaskResult) ProtoMessage() {}

func (x *ShardTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardTaskResult.ProtoReflect.Descriptor instead.
func (*ShardTaskResult) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{1}
}

func (x *ShardTaskResult) GetCompactTasks() []*CompactTask {
	if x != nil {
		return x.CompactTasks
	}
	return nil
}

type PathRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lower string `protobuf:"bytes,1,opt,name=lower,proto3" json:"lower,omitempty"`
	Upper string `protobuf:"bytes,2,opt,name=upper,proto3" json:"upper,omitempty"`
}

func (x *PathRange) Reset() {
	*x = PathRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathRange) ProtoMessage() {}

func (x *PathRange) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathRange.ProtoReflect.Descriptor instead.
func (*PathRange) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{2}
}

func (x *PathRange) GetLower() string {
	if x != nil {
		return x.Lower
	}
	return ""
}

func (x *PathRange) GetUpper() string {
	if x != nil {
		return x.Upper
	}
	return ""
}

type CompactTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs    []string   `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	PathRange *PathRange `protobuf:"bytes,2,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
}

func (x *CompactTask) Reset() {
	*x = CompactTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactTask) ProtoMessage() {}

func (x *CompactTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactTask.ProtoReflect.Descriptor instead.
func (*CompactTask) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{3}
}

func (x *CompactTask) GetInputs() []string {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *CompactTask) GetPathRange() *PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

type CompactTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CompactTaskResult) Reset() {
	*x = CompactTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactTaskResult) ProtoMessage() {}

func (x *CompactTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactTaskResult.ProtoReflect.Descriptor instead.
func (*CompactTaskResult) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{4}
}

func (x *CompactTaskResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ConcatTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs []string `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
}

func (x *ConcatTask) Reset() {
	*x = ConcatTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConcatTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConcatTask) ProtoMessage() {}

func (x *ConcatTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConcatTask.ProtoReflect.Descriptor instead.
func (*ConcatTask) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{5}
}

func (x *ConcatTask) GetInputs() []string {
	if x != nil {
		return x.Inputs
	}
	return nil
}

type ConcatTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConcatTaskResult) Reset() {
	*x = ConcatTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConcatTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConcatTaskResult) ProtoMessage() {}

func (x *ConcatTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConcatTaskResult.ProtoReflect.Descriptor instead.
func (*ConcatTaskResult) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{6}
}

func (x *ConcatTaskResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ValidateTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PathRange *PathRange `protobuf:"bytes,2,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
}

func (x *ValidateTask) Reset() {
	*x = ValidateTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTask) ProtoMessage() {}

func (x *ValidateTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTask.ProtoReflect.Descriptor instead.
func (*ValidateTask) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ValidateTask) GetPathRange() *PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

type ValidateTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First     *index.Index `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Last      *index.Index `protobuf:"bytes,2,opt,name=last,proto3" json:"last,omitempty"`
	Error     string       `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	SizeBytes int64        `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
}

func (x *ValidateTaskResult) Reset() {
	*x = ValidateTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTaskResult) ProtoMessage() {}

func (x *ValidateTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTaskResult.ProtoReflect.Descriptor instead.
func (*ValidateTaskResult) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateTaskResult) GetFirst() *index.Index {
	if x != nil {
		return x.First
	}
	return nil
}

func (x *ValidateTaskResult) GetLast() *index.Index {
	if x != nil {
		return x.Last
	}
	return nil
}

func (x *ValidateTaskResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ValidateTaskResult) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

type PutFileURLTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dst         string   `protobuf:"bytes,1,opt,name=dst,proto3" json:"dst,omitempty"`
	Datum       string   `protobuf:"bytes,2,opt,name=datum,proto3" json:"datum,omitempty"`
	URL         string   `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"`
	Paths       []string `protobuf:"bytes,4,rep,name=paths,proto3" json:"paths,omitempty"`
	StartOffset int64    `protobuf:"varint,5,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	EndOffset   int64    `protobuf:"varint,7,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
}

func (x *PutFileURLTask) Reset() {
	*x = PutFileURLTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileURLTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileURLTask) ProtoMessage() {}

func (x *PutFileURLTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileURLTask.ProtoReflect.Descriptor instead.
func (*PutFileURLTask) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{9}
}

func (x *PutFileURLTask) GetDst() string {
	if x != nil {
		return x.Dst
	}
	return ""
}

func (x *PutFileURLTask) GetDatum() string {
	if x != nil {
		return x.Datum
	}
	return ""
}

func (x *PutFileURLTask) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *PutFileURLTask) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *PutFileURLTask) GetStartOffset() int64 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *PutFileURLTask) GetEndOffset() int64 {
	if x != nil {
		return x.EndOffset
	}
	return 0
}

type PutFileURLTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PutFileURLTaskResult) Reset() {
	*x = PutFileURLTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileURLTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileURLTaskResult) ProtoMessage() {}

func (x *PutFileURLTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileURLTaskResult.ProtoReflect.Descriptor instead.
func (*PutFileURLTaskResult) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{10}
}

func (x *PutFileURLTaskResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFileURLTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL       string         `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	File      *pfs.File      `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	PathRange *pfs.PathRange `protobuf:"bytes,3,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
}

func (x *GetFileURLTask) Reset() {
	*x = GetFileURLTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileURLTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileURLTask) ProtoMessage() {}

func (x *GetFileURLTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileURLTask.ProtoReflect.Descriptor instead.
func (*GetFileURLTask) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{11}
}

func (x *GetFileURLTask) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *GetFileURLTask) GetFile() *pfs.File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *GetFileURLTask) GetPathRange() *pfs.PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

type GetFileURLTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFileURLTaskResult) Reset() {
	*x = GetFileURLTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pfs_server_pfsserver_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileURLTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileURLTaskResult) ProtoMessage() {}

func (x *GetFileURLTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_pfs_server_pfsserver_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileURLTaskResult.ProtoReflect.Descriptor instead.
func (*GetFileURLTaskResult) Descriptor() ([]byte, []int) {
	return file_server_pfs_server_pfsserver_proto_rawDescGZIP(), []int{12}
}

var File_server_pfs_server_pfsserver_proto protoreflect.FileDescriptor

var file_server_pfs_server_pfsserver_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x66, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x66, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x66, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x2a,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x66, 0x73, 0x2f,
	0x70, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x09, 0x53, 0x68, 0x61,
	0x72, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x33,
	0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x66, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x66, 0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x22, 0x37, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x70, 0x70, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x70, 0x70, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x66, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x70,
	0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x66,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x8f, 0x01, 0x0a,
	0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0xa2,
	0x01, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52,
	0x4c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x76, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12,
	0x20, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52,
	0x4c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x39, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64,
	0x65, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x66, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_pfs_server_pfsserver_proto_rawDescOnce sync.Once
	file_server_pfs_server_pfsserver_proto_rawDescData = file_server_pfs_server_pfsserver_proto_rawDesc
)

func file_server_pfs_server_pfsserver_proto_rawDescGZIP() []byte {
	file_server_pfs_server_pfsserver_proto_rawDescOnce.Do(func() {
		file_server_pfs_server_pfsserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_pfs_server_pfsserver_proto_rawDescData)
	})
	return file_server_pfs_server_pfsserver_proto_rawDescData
}

var file_server_pfs_server_pfsserver_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_server_pfs_server_pfsserver_proto_goTypes = []interface{}{
	(*ShardTask)(nil),            // 0: pfsserver.ShardTask
	(*ShardTaskResult)(nil),      // 1: pfsserver.ShardTaskResult
	(*PathRange)(nil),            // 2: pfsserver.PathRange
	(*CompactTask)(nil),          // 3: pfsserver.CompactTask
	(*CompactTaskResult)(nil),    // 4: pfsserver.CompactTaskResult
	(*ConcatTask)(nil),           // 5: pfsserver.ConcatTask
	(*ConcatTaskResult)(nil),     // 6: pfsserver.ConcatTaskResult
	(*ValidateTask)(nil),         // 7: pfsserver.ValidateTask
	(*ValidateTaskResult)(nil),   // 8: pfsserver.ValidateTaskResult
	(*PutFileURLTask)(nil),       // 9: pfsserver.PutFileURLTask
	(*PutFileURLTaskResult)(nil), // 10: pfsserver.PutFileURLTaskResult
	(*GetFileURLTask)(nil),       // 11: pfsserver.GetFileURLTask
	(*GetFileURLTaskResult)(nil), // 12: pfsserver.GetFileURLTaskResult
	(*index.Index)(nil),          // 13: index.Index
	(*pfs.File)(nil),             // 14: pfs_v2.File
	(*pfs.PathRange)(nil),        // 15: pfs_v2.PathRange
}
var file_server_pfs_server_pfsserver_proto_depIdxs = []int32{
	2,  // 0: pfsserver.ShardTask.path_range:type_name -> pfsserver.PathRange
	3,  // 1: pfsserver.ShardTaskResult.compact_tasks:type_name -> pfsserver.CompactTask
	2,  // 2: pfsserver.CompactTask.path_range:type_name -> pfsserver.PathRange
	2,  // 3: pfsserver.ValidateTask.path_range:type_name -> pfsserver.PathRange
	13, // 4: pfsserver.ValidateTaskResult.first:type_name -> index.Index
	13, // 5: pfsserver.ValidateTaskResult.last:type_name -> index.Index
	14, // 6: pfsserver.GetFileURLTask.file:type_name -> pfs_v2.File
	15, // 7: pfsserver.GetFileURLTask.path_range:type_name -> pfs_v2.PathRange
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_server_pfs_server_pfsserver_proto_init() }
func file_server_pfs_server_pfsserver_proto_init() {
	if File_server_pfs_server_pfsserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_pfs_server_pfsserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardTask); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardTaskResult); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathRange); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactTask); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactTaskResult); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConcatTask); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConcatTaskResult); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTask); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTaskResult); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileURLTask); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileURLTaskResult); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileURLTask); i {
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
		file_server_pfs_server_pfsserver_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileURLTaskResult); i {
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
			RawDescriptor: file_server_pfs_server_pfsserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_pfs_server_pfsserver_proto_goTypes,
		DependencyIndexes: file_server_pfs_server_pfsserver_proto_depIdxs,
		MessageInfos:      file_server_pfs_server_pfsserver_proto_msgTypes,
	}.Build()
	File_server_pfs_server_pfsserver_proto = out.File
	file_server_pfs_server_pfsserver_proto_rawDesc = nil
	file_server_pfs_server_pfsserver_proto_goTypes = nil
	file_server_pfs_server_pfsserver_proto_depIdxs = nil
}
