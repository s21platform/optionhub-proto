// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.26.1
// source: optionhub.proto

package optionhub_proto

import (
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

type EmptyOptionhub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyOptionhub) Reset() {
	*x = EmptyOptionhub{}
	mi := &file_optionhub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyOptionhub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyOptionhub) ProtoMessage() {}

func (x *EmptyOptionhub) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyOptionhub.ProtoReflect.Descriptor instead.
func (*EmptyOptionhub) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_optionhub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{1}
}

func (x *Record) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Record) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type GetByNameIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetByNameIn) Reset() {
	*x = GetByNameIn{}
	mi := &file_optionhub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByNameIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByNameIn) ProtoMessage() {}

func (x *GetByNameIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByNameIn.ProtoReflect.Descriptor instead.
func (*GetByNameIn) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{2}
}

func (x *GetByNameIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetByNameOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options []*Record `protobuf:"bytes,1,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *GetByNameOut) Reset() {
	*x = GetByNameOut{}
	mi := &file_optionhub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByNameOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByNameOut) ProtoMessage() {}

func (x *GetByNameOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByNameOut.ProtoReflect.Descriptor instead.
func (*GetByNameOut) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{3}
}

func (x *GetByNameOut) GetOptions() []*Record {
	if x != nil {
		return x.Options
	}
	return nil
}

// message request
type GetByIdIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the row in the db
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdIn) Reset() {
	*x = GetByIdIn{}
	mi := &file_optionhub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByIdIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdIn) ProtoMessage() {}

func (x *GetByIdIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdIn.ProtoReflect.Descriptor instead.
func (*GetByIdIn) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIdIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// message response
type GetByIdOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetByIdOut) Reset() {
	*x = GetByIdOut{}
	mi := &file_optionhub_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByIdOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdOut) ProtoMessage() {}

func (x *GetByIdOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdOut.ProtoReflect.Descriptor instead.
func (*GetByIdOut) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{5}
}

func (x *GetByIdOut) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetByIdOut) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// message response for all active values
type GetAllOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Record `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *GetAllOut) Reset() {
	*x = GetAllOut{}
	mi := &file_optionhub_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOut) ProtoMessage() {}

func (x *GetAllOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOut.ProtoReflect.Descriptor instead.
func (*GetAllOut) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllOut) GetValues() []*Record {
	if x != nil {
		return x.Values
	}
	return nil
}

// message request
type AddIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// value that will be added
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddIn) Reset() {
	*x = AddIn{}
	mi := &file_optionhub_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddIn) ProtoMessage() {}

func (x *AddIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddIn.ProtoReflect.Descriptor instead.
func (*AddIn) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{7}
}

func (x *AddIn) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// message response
type AddOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of added value
	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddOut) Reset() {
	*x = AddOut{}
	mi := &file_optionhub_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOut) ProtoMessage() {}

func (x *AddOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOut.ProtoReflect.Descriptor instead.
func (*AddOut) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{8}
}

func (x *AddOut) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddOut) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// message request
type SetByIdIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SetByIdIn) Reset() {
	*x = SetByIdIn{}
	mi := &file_optionhub_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetByIdIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetByIdIn) ProtoMessage() {}

func (x *SetByIdIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetByIdIn.ProtoReflect.Descriptor instead.
func (*SetByIdIn) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{9}
}

func (x *SetByIdIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// message response
type SetByIdOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of value that will be changed
	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetByIdOut) Reset() {
	*x = SetByIdOut{}
	mi := &file_optionhub_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetByIdOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetByIdOut) ProtoMessage() {}

func (x *SetByIdOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetByIdOut.ProtoReflect.Descriptor instead.
func (*SetByIdOut) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{10}
}

func (x *SetByIdOut) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetByIdOut) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// message request
type DeleteByIdIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// value with such id will be marked as inactive
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteByIdIn) Reset() {
	*x = DeleteByIdIn{}
	mi := &file_optionhub_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteByIdIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIdIn) ProtoMessage() {}

func (x *DeleteByIdIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByIdIn.ProtoReflect.Descriptor instead.
func (*DeleteByIdIn) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteByIdIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// message response
type DeleteByIdOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// flag that says if deletion was successful
	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteByIdOut) Reset() {
	*x = DeleteByIdOut{}
	mi := &file_optionhub_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteByIdOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIdOut) ProtoMessage() {}

func (x *DeleteByIdOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByIdOut.ProtoReflect.Descriptor instead.
func (*DeleteByIdOut) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteByIdOut) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_optionhub_proto protoreflect.FileDescriptor

var file_optionhub_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x68, 0x75, 0x62, 0x22, 0x2e, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x1d, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x49,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2e, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4f, 0x75,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f,
	0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xab, 0x07, 0x0a, 0x10, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x68, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x73, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x73, 0x12, 0x0f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x68, 0x75, 0x62, 0x1a, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x4f, 0x73, 0x12, 0x06,
	0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x22,
	0x00, 0x12, 0x26, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x4f, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0a,
	0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x53, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f,
	0x75, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x12, 0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x12, 0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x62, 0x62, 0x79, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x0c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x1a, 0x0d, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x29,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x48, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e,
	0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x1a, 0x0d,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12,
	0x29, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07,
	0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x1a, 0x0d,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12,
	0x28, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x68, 0x75, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_optionhub_proto_rawDescOnce sync.Once
	file_optionhub_proto_rawDescData = file_optionhub_proto_rawDesc
)

func file_optionhub_proto_rawDescGZIP() []byte {
	file_optionhub_proto_rawDescOnce.Do(func() {
		file_optionhub_proto_rawDescData = protoimpl.X.CompressGZIP(file_optionhub_proto_rawDescData)
	})
	return file_optionhub_proto_rawDescData
}

var file_optionhub_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_optionhub_proto_goTypes = []any{
	(*EmptyOptionhub)(nil), // 0: EmptyOptionhub
	(*Record)(nil),         // 1: Record
	(*GetByNameIn)(nil),    // 2: GetByNameIn
	(*GetByNameOut)(nil),   // 3: GetByNameOut
	(*GetByIdIn)(nil),      // 4: GetByIdIn
	(*GetByIdOut)(nil),     // 5: GetByIdOut
	(*GetAllOut)(nil),      // 6: GetAllOut
	(*AddIn)(nil),          // 7: AddIn
	(*AddOut)(nil),         // 8: AddOut
	(*SetByIdIn)(nil),      // 9: SetByIdIn
	(*SetByIdOut)(nil),     // 10: SetByIdOut
	(*DeleteByIdIn)(nil),   // 11: DeleteByIdIn
	(*DeleteByIdOut)(nil),  // 12: DeleteByIdOut
}
var file_optionhub_proto_depIdxs = []int32{
	1,  // 0: GetByNameOut.options:type_name -> Record
	1,  // 1: GetAllOut.values:type_name -> Record
	2,  // 2: OptionhubService.GetOsBySearchName:input_type -> GetByNameIn
	4,  // 3: OptionhubService.GetOsByID:input_type -> GetByIdIn
	0,  // 4: OptionhubService.GetAllOs:input_type -> EmptyOptionhub
	7,  // 5: OptionhubService.AddOs:input_type -> AddIn
	9,  // 6: OptionhubService.SetOsByID:input_type -> SetByIdIn
	11, // 7: OptionhubService.DeleteOsByID:input_type -> DeleteByIdIn
	2,  // 8: OptionhubService.GetWorkPlaceBySearchName:input_type -> GetByNameIn
	4,  // 9: OptionhubService.GetWorkPlaceById:input_type -> GetByIdIn
	7,  // 10: OptionhubService.AddWorkPlace:input_type -> AddIn
	2,  // 11: OptionhubService.GetStudyPlaceBySearchName:input_type -> GetByNameIn
	4,  // 12: OptionhubService.GetStudyPlaceById:input_type -> GetByIdIn
	7,  // 13: OptionhubService.AddStudyPlace:input_type -> AddIn
	2,  // 14: OptionhubService.GetHobbyBySearchName:input_type -> GetByNameIn
	4,  // 15: OptionhubService.GetHobbyById:input_type -> GetByIdIn
	7,  // 16: OptionhubService.AddHobby:input_type -> AddIn
	2,  // 17: OptionhubService.GetSkillBySearchName:input_type -> GetByNameIn
	4,  // 18: OptionhubService.GetSkillById:input_type -> GetByIdIn
	7,  // 19: OptionhubService.AddSkill:input_type -> AddIn
	2,  // 20: OptionhubService.GetCityBySearchName:input_type -> GetByNameIn
	4,  // 21: OptionhubService.GetCityById:input_type -> GetByIdIn
	7,  // 22: OptionhubService.AddCity:input_type -> AddIn
	3,  // 23: OptionhubService.GetOsBySearchName:output_type -> GetByNameOut
	5,  // 24: OptionhubService.GetOsByID:output_type -> GetByIdOut
	6,  // 25: OptionhubService.GetAllOs:output_type -> GetAllOut
	8,  // 26: OptionhubService.AddOs:output_type -> AddOut
	10, // 27: OptionhubService.SetOsByID:output_type -> SetByIdOut
	12, // 28: OptionhubService.DeleteOsByID:output_type -> DeleteByIdOut
	3,  // 29: OptionhubService.GetWorkPlaceBySearchName:output_type -> GetByNameOut
	5,  // 30: OptionhubService.GetWorkPlaceById:output_type -> GetByIdOut
	8,  // 31: OptionhubService.AddWorkPlace:output_type -> AddOut
	3,  // 32: OptionhubService.GetStudyPlaceBySearchName:output_type -> GetByNameOut
	5,  // 33: OptionhubService.GetStudyPlaceById:output_type -> GetByIdOut
	8,  // 34: OptionhubService.AddStudyPlace:output_type -> AddOut
	3,  // 35: OptionhubService.GetHobbyBySearchName:output_type -> GetByNameOut
	5,  // 36: OptionhubService.GetHobbyById:output_type -> GetByIdOut
	8,  // 37: OptionhubService.AddHobby:output_type -> AddOut
	3,  // 38: OptionhubService.GetSkillBySearchName:output_type -> GetByNameOut
	5,  // 39: OptionhubService.GetSkillById:output_type -> GetByIdOut
	8,  // 40: OptionhubService.AddSkill:output_type -> AddOut
	3,  // 41: OptionhubService.GetCityBySearchName:output_type -> GetByNameOut
	5,  // 42: OptionhubService.GetCityById:output_type -> GetByIdOut
	8,  // 43: OptionhubService.AddCity:output_type -> AddOut
	23, // [23:44] is the sub-list for method output_type
	2,  // [2:23] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_optionhub_proto_init() }
func file_optionhub_proto_init() {
	if File_optionhub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_optionhub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_optionhub_proto_goTypes,
		DependencyIndexes: file_optionhub_proto_depIdxs,
		MessageInfos:      file_optionhub_proto_msgTypes,
	}.Build()
	File_optionhub_proto = out.File
	file_optionhub_proto_rawDesc = nil
	file_optionhub_proto_goTypes = nil
	file_optionhub_proto_depIdxs = nil
}
