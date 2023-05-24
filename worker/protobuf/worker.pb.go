// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: worker.proto

package protobuf

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ThumbnailRequestType int32

const (
	ThumbnailRequestType_LIVE    ThumbnailRequestType = 0
	ThumbnailRequestType_VOD     ThumbnailRequestType = 1
	ThumbnailRequestType_SECTION ThumbnailRequestType = 2
)

// Enum value maps for ThumbnailRequestType.
var (
	ThumbnailRequestType_name = map[int32]string{
		0: "LIVE",
		1: "VOD",
		2: "SECTION",
	}
	ThumbnailRequestType_value = map[string]int32{
		"LIVE":    0,
		"VOD":     1,
		"SECTION": 2,
	}
)

func (x ThumbnailRequestType) Enum() *ThumbnailRequestType {
	p := new(ThumbnailRequestType)
	*p = x
	return p
}

func (x ThumbnailRequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThumbnailRequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_worker_proto_enumTypes[0].Descriptor()
}

func (ThumbnailRequestType) Type() protoreflect.EnumType {
	return &file_worker_proto_enumTypes[0]
}

func (x ThumbnailRequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThumbnailRequestType.Descriptor instead.
func (ThumbnailRequestType) EnumDescriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream  uint64               `protobuf:"varint,1,opt,name=stream,proto3" json:"stream,omitempty"`
	Course  uint64               `protobuf:"varint,2,opt,name=course,proto3" json:"course,omitempty"`
	Version string               `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	End     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequest) GetStream() uint64 {
	if x != nil {
		return x.Stream
	}
	return 0
}

func (x *StreamRequest) GetCourse() uint64 {
	if x != nil {
		return x.Course
	}
	return 0
}

func (x *StreamRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *StreamRequest) GetEnd() *timestamp.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job string `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResponse) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

type StreamEndRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID   string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	KeepVod bool   `protobuf:"varint,2,opt,name=keepVod,proto3" json:"keepVod,omitempty"`
}

func (x *StreamEndRequest) Reset() {
	*x = StreamEndRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEndRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEndRequest) ProtoMessage() {}

func (x *StreamEndRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEndRequest.ProtoReflect.Descriptor instead.
func (*StreamEndRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *StreamEndRequest) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *StreamEndRequest) GetKeepVod() bool {
	if x != nil {
		return x.KeepVod
	}
	return false
}

type StreamEndResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamEndResponse) Reset() {
	*x = StreamEndResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEndResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEndResponse) ProtoMessage() {}

func (x *StreamEndResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEndResponse.ProtoReflect.Descriptor instead.
func (*StreamEndResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{3}
}

type GenerateThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             ThumbnailRequestType `protobuf:"varint,1,opt,name=type,proto3,enum=protobuf.ThumbnailRequestType" json:"type,omitempty"`
	Input            string               `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`                               // filepath or url to video
	TimestampSeconds []uint64             `protobuf:"varint,3,rep,packed,name=timestampSeconds,proto3" json:"timestampSeconds,omitempty"` // Only used for SECTION
}

func (x *GenerateThumbnailRequest) Reset() {
	*x = GenerateThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateThumbnailRequest) ProtoMessage() {}

func (x *GenerateThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateThumbnailRequest.ProtoReflect.Descriptor instead.
func (*GenerateThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateThumbnailRequest) GetType() ThumbnailRequestType {
	if x != nil {
		return x.Type
	}
	return ThumbnailRequestType_LIVE
}

func (x *GenerateThumbnailRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *GenerateThumbnailRequest) GetTimestampSeconds() []uint64 {
	if x != nil {
		return x.TimestampSeconds
	}
	return nil
}

type GenerateThumbnailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *GenerateThumbnailResponse) Reset() {
	*x = GenerateThumbnailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateThumbnailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateThumbnailResponse) ProtoMessage() {}

func (x *GenerateThumbnailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateThumbnailResponse.ProtoReflect.Descriptor instead.
func (*GenerateThumbnailResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateThumbnailResponse) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Workload uint32 `protobuf:"varint,2,opt,name=Workload,proto3" json:"Workload,omitempty"`
	Version  string `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	CPU      string `protobuf:"bytes,4,opt,name=CPU,proto3" json:"CPU,omitempty"`
	Memory   string `protobuf:"bytes,5,opt,name=Memory,proto3" json:"Memory,omitempty"`
	Disk     string `protobuf:"bytes,6,opt,name=Disk,proto3" json:"Disk,omitempty"`
	Uptime   string `protobuf:"bytes,7,opt,name=Uptime,proto3" json:"Uptime,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{9}
}

func (x *HeartbeatRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *HeartbeatRequest) GetWorkload() uint32 {
	if x != nil {
		return x.Workload
	}
	return 0
}

func (x *HeartbeatRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HeartbeatRequest) GetCPU() string {
	if x != nil {
		return x.CPU
	}
	return ""
}

func (x *HeartbeatRequest) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *HeartbeatRequest) GetDisk() string {
	if x != nil {
		return x.Disk
	}
	return ""
}

func (x *HeartbeatRequest) GetUptime() string {
	if x != nil {
		return x.Uptime
	}
	return ""
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{10}
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0x22, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x42, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x65, 0x70, 0x56, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x6b, 0x65, 0x65, 0x70, 0x56, 0x6f, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x90, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x22, 0x2f, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x2d,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xae, 0x01, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x43, 0x50, 0x55, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x50, 0x55, 0x12, 0x16,
	0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x36, 0x0a, 0x14, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x56, 0x4f, 0x44,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32,
	0x9f, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x45, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x99, 0x01, 0x0a, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_worker_proto_goTypes = []interface{}{
	(ThumbnailRequestType)(0),         // 0: protobuf.ThumbnailRequestType
	(*StreamRequest)(nil),             // 1: protobuf.StreamRequest
	(*StreamResponse)(nil),            // 2: protobuf.StreamResponse
	(*StreamEndRequest)(nil),          // 3: protobuf.StreamEndRequest
	(*StreamEndResponse)(nil),         // 4: protobuf.StreamEndResponse
	(*GenerateThumbnailRequest)(nil),  // 5: protobuf.GenerateThumbnailRequest
	(*GenerateThumbnailResponse)(nil), // 6: protobuf.GenerateThumbnailResponse
	(*Status)(nil),                    // 7: protobuf.Status
	(*RegisterRequest)(nil),           // 8: protobuf.RegisterRequest
	(*RegisterResponse)(nil),          // 9: protobuf.RegisterResponse
	(*HeartbeatRequest)(nil),          // 10: protobuf.HeartbeatRequest
	(*HeartbeatResponse)(nil),         // 11: protobuf.HeartbeatResponse
	(*timestamp.Timestamp)(nil),       // 12: google.protobuf.Timestamp
}
var file_worker_proto_depIdxs = []int32{
	12, // 0: protobuf.StreamRequest.end:type_name -> google.protobuf.Timestamp
	0,  // 1: protobuf.GenerateThumbnailRequest.type:type_name -> protobuf.ThumbnailRequestType
	1,  // 2: protobuf.ToWorker.RequestStream:input_type -> protobuf.StreamRequest
	3,  // 3: protobuf.ToWorker.RequestStreamEnd:input_type -> protobuf.StreamEndRequest
	8,  // 4: protobuf.FromWorker.Register:input_type -> protobuf.RegisterRequest
	10, // 5: protobuf.FromWorker.Heartbeat:input_type -> protobuf.HeartbeatRequest
	2,  // 6: protobuf.ToWorker.RequestStream:output_type -> protobuf.StreamResponse
	4,  // 7: protobuf.ToWorker.RequestStreamEnd:output_type -> protobuf.StreamEndResponse
	9,  // 8: protobuf.FromWorker.Register:output_type -> protobuf.RegisterResponse
	11, // 9: protobuf.FromWorker.Heartbeat:output_type -> protobuf.HeartbeatResponse
	6,  // [6:10] is the sub-list for method output_type
	2,  // [2:6] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEndRequest); i {
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
		file_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEndResponse); i {
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
		file_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateThumbnailRequest); i {
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
		file_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateThumbnailResponse); i {
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
		file_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_worker_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		EnumInfos:         file_worker_proto_enumTypes,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}
