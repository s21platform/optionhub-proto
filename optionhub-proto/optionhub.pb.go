// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
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

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_optionhub_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Record) GetValue() string {
	if x != nil {
		return x.Value
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByNameIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByNameIn) ProtoMessage() {}

func (x *GetByNameIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_optionhub_proto_rawDescGZIP(), []int{1}
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

	Values []*Record `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *GetByNameOut) Reset() {
	*x = GetByNameOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByNameOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByNameOut) ProtoMessage() {}

func (x *GetByNameOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_optionhub_proto_rawDescGZIP(), []int{2}
}

func (x *GetByNameOut) GetValues() []*Record {
	if x != nil {
		return x.Values
	}
	return nil
}

// message request
type GetByIdIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user's id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdIn) Reset() {
	*x = GetByIdIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdIn) ProtoMessage() {}

func (x *GetByIdIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_optionhub_proto_rawDescGZIP(), []int{3}
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdOut) ProtoMessage() {}

func (x *GetByIdOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_optionhub_proto_rawDescGZIP(), []int{4}
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

type GetAllIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllIn) Reset() {
	*x = GetAllIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllIn) ProtoMessage() {}

func (x *GetAllIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllIn.ProtoReflect.Descriptor instead.
func (*GetAllIn) Descriptor() ([]byte, []int) {
	return file_optionhub_proto_rawDescGZIP(), []int{5}
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOut) ProtoMessage() {}

func (x *GetAllOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddIn) ProtoMessage() {}

func (x *AddIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOut) ProtoMessage() {}

func (x *AddOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetByIdIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetByIdIn) ProtoMessage() {}

func (x *SetByIdIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetByIdOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetByIdOut) ProtoMessage() {}

func (x *SetByIdOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByIdIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIdIn) ProtoMessage() {}

func (x *DeleteByIdIn) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_optionhub_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByIdOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByIdOut) ProtoMessage() {}

func (x *DeleteByIdOut) ProtoReflect() protoreflect.Message {
	mi := &file_optionhub_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
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
	0x6f, 0x22, 0x2e, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x4f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x49, 0x6e, 0x22, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x75, 0x74, 0x12,
	0x1f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x1d, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x2e, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x1b, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0a,
	0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x32, 0xcb, 0x0a, 0x0a, 0x10, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x68, 0x75, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x73, 0x42,
	0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x4f, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x12, 0x23, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x73, 0x12, 0x09,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x4f, 0x73,
	0x12, 0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x4f, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0a, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x53,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x09,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x57, 0x6f,
	0x72, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a,
	0x07, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a,
	0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x53, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a,
	0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22,
	0x00, 0x12, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x79,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e,
	0x1a, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x22,
	0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12,
	0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x12, 0x2e, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64,
	0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x62,
	0x62, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x1d, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x06, 0x2e, 0x41,
	0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12,
	0x29, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0a, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x53, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x29,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12,
	0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0a, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b,
	0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a,
	0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22,
	0x00, 0x12, 0x28, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f,
	0x75, 0x74, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x06, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x1a, 0x07, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x12, 0x28, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0a, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0b, 0x2e,
	0x53, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x49, 0x6e, 0x1a, 0x0e, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42,
	0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x68, 0x75, 0x62, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*Record)(nil),        // 0: Record
	(*GetByNameIn)(nil),   // 1: GetByNameIn
	(*GetByNameOut)(nil),  // 2: GetByNameOut
	(*GetByIdIn)(nil),     // 3: GetByIdIn
	(*GetByIdOut)(nil),    // 4: GetByIdOut
	(*GetAllIn)(nil),      // 5: GetAllIn
	(*GetAllOut)(nil),     // 6: GetAllOut
	(*AddIn)(nil),         // 7: AddIn
	(*AddOut)(nil),        // 8: AddOut
	(*SetByIdIn)(nil),     // 9: SetByIdIn
	(*SetByIdOut)(nil),    // 10: SetByIdOut
	(*DeleteByIdIn)(nil),  // 11: DeleteByIdIn
	(*DeleteByIdOut)(nil), // 12: DeleteByIdOut
}
var file_optionhub_proto_depIdxs = []int32{
	0,  // 0: GetByNameOut.values:type_name -> Record
	0,  // 1: GetAllOut.values:type_name -> Record
	1,  // 2: OptionhubService.GetOsBySearchName:input_type -> GetByNameIn
	3,  // 3: OptionhubService.GetOsById:input_type -> GetByIdIn
	5,  // 4: OptionhubService.GetAllOs:input_type -> GetAllIn
	7,  // 5: OptionhubService.AddOs:input_type -> AddIn
	9,  // 6: OptionhubService.SetOsById:input_type -> SetByIdIn
	11, // 7: OptionhubService.DeleteOsById:input_type -> DeleteByIdIn
	3,  // 8: OptionhubService.GetWorkPlaceById:input_type -> GetByIdIn
	5,  // 9: OptionhubService.GetAllWorkPlace:input_type -> GetAllIn
	7,  // 10: OptionhubService.AddWorkPlace:input_type -> AddIn
	9,  // 11: OptionhubService.SetWorkPlaceById:input_type -> SetByIdIn
	11, // 12: OptionhubService.DeleteWorkPlaceById:input_type -> DeleteByIdIn
	3,  // 13: OptionhubService.GetStudyPlaceById:input_type -> GetByIdIn
	5,  // 14: OptionhubService.GetAllStudyPlace:input_type -> GetAllIn
	7,  // 15: OptionhubService.AddStudyPlace:input_type -> AddIn
	9,  // 16: OptionhubService.SetStudyPlaceById:input_type -> SetByIdIn
	11, // 17: OptionhubService.DeleteStudyPlaceById:input_type -> DeleteByIdIn
	3,  // 18: OptionhubService.GetHobbyById:input_type -> GetByIdIn
	5,  // 19: OptionhubService.GetHobbyPlace:input_type -> GetAllIn
	7,  // 20: OptionhubService.AddHobby:input_type -> AddIn
	9,  // 21: OptionhubService.SetHobbyById:input_type -> SetByIdIn
	11, // 22: OptionhubService.DeleteHobbyById:input_type -> DeleteByIdIn
	3,  // 23: OptionhubService.GetSkillById:input_type -> GetByIdIn
	5,  // 24: OptionhubService.GetSkillPlace:input_type -> GetAllIn
	7,  // 25: OptionhubService.AddSkill:input_type -> AddIn
	9,  // 26: OptionhubService.SetSkillById:input_type -> SetByIdIn
	11, // 27: OptionhubService.DeleteSkillById:input_type -> DeleteByIdIn
	3,  // 28: OptionhubService.GetCityById:input_type -> GetByIdIn
	5,  // 29: OptionhubService.GetCityPlace:input_type -> GetAllIn
	7,  // 30: OptionhubService.AddCity:input_type -> AddIn
	9,  // 31: OptionhubService.SetCityById:input_type -> SetByIdIn
	11, // 32: OptionhubService.DeleteCityById:input_type -> DeleteByIdIn
	2,  // 33: OptionhubService.GetOsBySearchName:output_type -> GetByNameOut
	4,  // 34: OptionhubService.GetOsById:output_type -> GetByIdOut
	6,  // 35: OptionhubService.GetAllOs:output_type -> GetAllOut
	8,  // 36: OptionhubService.AddOs:output_type -> AddOut
	10, // 37: OptionhubService.SetOsById:output_type -> SetByIdOut
	12, // 38: OptionhubService.DeleteOsById:output_type -> DeleteByIdOut
	4,  // 39: OptionhubService.GetWorkPlaceById:output_type -> GetByIdOut
	6,  // 40: OptionhubService.GetAllWorkPlace:output_type -> GetAllOut
	8,  // 41: OptionhubService.AddWorkPlace:output_type -> AddOut
	10, // 42: OptionhubService.SetWorkPlaceById:output_type -> SetByIdOut
	12, // 43: OptionhubService.DeleteWorkPlaceById:output_type -> DeleteByIdOut
	4,  // 44: OptionhubService.GetStudyPlaceById:output_type -> GetByIdOut
	6,  // 45: OptionhubService.GetAllStudyPlace:output_type -> GetAllOut
	8,  // 46: OptionhubService.AddStudyPlace:output_type -> AddOut
	10, // 47: OptionhubService.SetStudyPlaceById:output_type -> SetByIdOut
	12, // 48: OptionhubService.DeleteStudyPlaceById:output_type -> DeleteByIdOut
	4,  // 49: OptionhubService.GetHobbyById:output_type -> GetByIdOut
	6,  // 50: OptionhubService.GetHobbyPlace:output_type -> GetAllOut
	8,  // 51: OptionhubService.AddHobby:output_type -> AddOut
	10, // 52: OptionhubService.SetHobbyById:output_type -> SetByIdOut
	12, // 53: OptionhubService.DeleteHobbyById:output_type -> DeleteByIdOut
	4,  // 54: OptionhubService.GetSkillById:output_type -> GetByIdOut
	6,  // 55: OptionhubService.GetSkillPlace:output_type -> GetAllOut
	8,  // 56: OptionhubService.AddSkill:output_type -> AddOut
	10, // 57: OptionhubService.SetSkillById:output_type -> SetByIdOut
	12, // 58: OptionhubService.DeleteSkillById:output_type -> DeleteByIdOut
	4,  // 59: OptionhubService.GetCityById:output_type -> GetByIdOut
	6,  // 60: OptionhubService.GetCityPlace:output_type -> GetAllOut
	8,  // 61: OptionhubService.AddCity:output_type -> AddOut
	10, // 62: OptionhubService.SetCityById:output_type -> SetByIdOut
	12, // 63: OptionhubService.DeleteCityById:output_type -> DeleteByIdOut
	33, // [33:64] is the sub-list for method output_type
	2,  // [2:33] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_optionhub_proto_init() }
func file_optionhub_proto_init() {
	if File_optionhub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_optionhub_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Record); i {
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
		file_optionhub_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetByNameIn); i {
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
		file_optionhub_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetByNameOut); i {
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
		file_optionhub_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetByIdIn); i {
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
		file_optionhub_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetByIdOut); i {
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
		file_optionhub_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllIn); i {
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
		file_optionhub_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllOut); i {
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
		file_optionhub_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AddIn); i {
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
		file_optionhub_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*AddOut); i {
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
		file_optionhub_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SetByIdIn); i {
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
		file_optionhub_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SetByIdOut); i {
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
		file_optionhub_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteByIdIn); i {
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
		file_optionhub_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteByIdOut); i {
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
