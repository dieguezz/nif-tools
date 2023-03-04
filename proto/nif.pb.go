// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/nif.proto

package protopb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NIF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nif string `protobuf:"bytes,1,opt,name=nif,proto3" json:"nif,omitempty"`
}

func (x *NIF) Reset() {
	*x = NIF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NIF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NIF) ProtoMessage() {}

func (x *NIF) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NIF.ProtoReflect.Descriptor instead.
func (*NIF) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{0}
}

func (x *NIF) GetNif() string {
	if x != nil {
		return x.Nif
	}
	return ""
}

type NIFs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NIFs []string `protobuf:"bytes,1,rep,name=NIFs,proto3" json:"NIFs,omitempty"`
}

func (x *NIFs) Reset() {
	*x = NIFs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NIFs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NIFs) ProtoMessage() {}

func (x *NIFs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NIFs.ProtoReflect.Descriptor instead.
func (*NIFs) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{1}
}

func (x *NIFs) GetNIFs() []string {
	if x != nil {
		return x.NIFs
	}
	return nil
}

type NIEs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NIEs []string `protobuf:"bytes,1,rep,name=NIEs,proto3" json:"NIEs,omitempty"`
}

func (x *NIEs) Reset() {
	*x = NIEs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NIEs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NIEs) ProtoMessage() {}

func (x *NIEs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NIEs.ProtoReflect.Descriptor instead.
func (*NIEs) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{2}
}

func (x *NIEs) GetNIEs() []string {
	if x != nil {
		return x.NIEs
	}
	return nil
}

type CIF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cif string `protobuf:"bytes,1,opt,name=cif,proto3" json:"cif,omitempty"`
}

func (x *CIF) Reset() {
	*x = CIF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIF) ProtoMessage() {}

func (x *CIF) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIF.ProtoReflect.Descriptor instead.
func (*CIF) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{3}
}

func (x *CIF) GetCif() string {
	if x != nil {
		return x.Cif
	}
	return ""
}

type CIFs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CIFs []string `protobuf:"bytes,1,rep,name=CIFs,proto3" json:"CIFs,omitempty"`
}

func (x *CIFs) Reset() {
	*x = CIFs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIFs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIFs) ProtoMessage() {}

func (x *CIFs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIFs.ProtoReflect.Descriptor instead.
func (*CIFs) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{4}
}

func (x *CIFs) GetCIFs() []string {
	if x != nil {
		return x.CIFs
	}
	return nil
}

type BulkParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount *int32 `protobuf:"varint,1,opt,name=amount,proto3,oneof" json:"amount,omitempty"`
}

func (x *BulkParams) Reset() {
	*x = BulkParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkParams) ProtoMessage() {}

func (x *BulkParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkParams.ProtoReflect.Descriptor instead.
func (*BulkParams) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{5}
}

func (x *BulkParams) GetAmount() int32 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{6}
}

func (x *Document) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

type ControlDigitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControlDigit string `protobuf:"bytes,1,opt,name=control_digit,json=controlDigit,proto3" json:"control_digit,omitempty"`
}

func (x *ControlDigitResponse) Reset() {
	*x = ControlDigitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlDigitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlDigitResponse) ProtoMessage() {}

func (x *ControlDigitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlDigitResponse.ProtoReflect.Descriptor instead.
func (*ControlDigitResponse) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{7}
}

func (x *ControlDigitResponse) GetControlDigit() string {
	if x != nil {
		return x.ControlDigit
	}
	return ""
}

type ParsedDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number  int32  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Kind    string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Control string `protobuf:"bytes,3,opt,name=control,proto3" json:"control,omitempty"`
}

func (x *ParsedDocumentResponse) Reset() {
	*x = ParsedDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParsedDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParsedDocumentResponse) ProtoMessage() {}

func (x *ParsedDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParsedDocumentResponse.ProtoReflect.Descriptor instead.
func (*ParsedDocumentResponse) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{8}
}

func (x *ParsedDocumentResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ParsedDocumentResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ParsedDocumentResponse) GetControl() string {
	if x != nil {
		return x.Control
	}
	return ""
}

type TypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TypeResponse) Reset() {
	*x = TypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeResponse) ProtoMessage() {}

func (x *TypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeResponse.ProtoReflect.Descriptor instead.
func (*TypeResponse) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{9}
}

func (x *TypeResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type NIE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nie string `protobuf:"bytes,1,opt,name=nie,proto3" json:"nie,omitempty"`
}

func (x *NIE) Reset() {
	*x = NIE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NIE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NIE) ProtoMessage() {}

func (x *NIE) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NIE.ProtoReflect.Descriptor instead.
func (*NIE) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{10}
}

func (x *NIE) GetNie() string {
	if x != nil {
		return x.Nie
	}
	return ""
}

type IsValid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (x *IsValid) Reset() {
	*x = IsValid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_nif_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValid) ProtoMessage() {}

func (x *IsValid) ProtoReflect() protoreflect.Message {
	mi := &file_proto_nif_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValid.ProtoReflect.Descriptor instead.
func (*IsValid) Descriptor() ([]byte, []int) {
	return file_proto_nif_proto_rawDescGZIP(), []int{11}
}

func (x *IsValid) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

var File_proto_nif_proto protoreflect.FileDescriptor

var file_proto_nif_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x69, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03, 0x4e, 0x49, 0x46, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x66, 0x22, 0x1a, 0x0a, 0x04,
	0x4e, 0x49, 0x46, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x49, 0x46, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x49, 0x46, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x4e, 0x49, 0x45, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x49, 0x45, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x49, 0x45, 0x73, 0x22, 0x17, 0x0a, 0x03, 0x43, 0x49, 0x46, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x66, 0x22, 0x1a, 0x0a,
	0x04, 0x43, 0x49, 0x46, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x49, 0x46, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x43, 0x49, 0x46, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x42, 0x75, 0x6c,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x26, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x44, 0x69, 0x67, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x44,
	0x69, 0x67, 0x69, 0x74, 0x22, 0x5e, 0x0a, 0x16, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x22, 0x22, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x17, 0x0a, 0x03, 0x4e, 0x49, 0x45, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69,
	0x65, 0x22, 0x23, 0x0a, 0x07, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x32, 0xe7, 0x08, 0x0a, 0x06, 0x4e, 0x69, 0x66, 0x41, 0x70,
	0x69, 0x12, 0x5e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x49, 0x46, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x44, 0x69, 0x67, 0x69, 0x74, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x49, 0x46, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x44, 0x69, 0x67, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x6e, 0x69, 0x66, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x64, 0x69, 0x67, 0x69, 0x74, 0x2f, 0x7b, 0x6e, 0x69, 0x66,
	0x7d, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x49, 0x46, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x6e, 0x69, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x7b, 0x6e, 0x69, 0x66, 0x7d, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4e, 0x49, 0x46, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x49, 0x46, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x6e, 0x69, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x4a, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x49, 0x46, 0x73,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x49, 0x46, 0x73,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x6e, 0x69, 0x66, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x12, 0x46, 0x0a, 0x0b,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x49, 0x46, 0x12, 0x0a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x49, 0x46, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x6e, 0x69, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b,
	0x6e, 0x69, 0x66, 0x7d, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x4e, 0x49, 0x45, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x49, 0x45, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x6e, 0x69, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x4a,
	0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x49, 0x45, 0x73, 0x12, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x49, 0x45, 0x73, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x6e, 0x69, 0x65, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x12, 0x46, 0x0a, 0x0b, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x49, 0x45, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4e, 0x49, 0x45, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x6e, 0x69, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x6e, 0x69,
	0x65, 0x7d, 0x12, 0x67, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12,
	0x1a, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x2f, 0x7b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x12, 0x5a, 0x0a, 0x10, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x12, 0x5e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x49,
	0x46, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x44, 0x69, 0x67, 0x69, 0x74, 0x12, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x49, 0x46, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x44, 0x69, 0x67, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17,
	0x2f, 0x63, 0x69, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x2f, 0x7b, 0x63, 0x69, 0x66, 0x7d, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x43, 0x49, 0x46, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x75, 0x6c, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x49, 0x46, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x63, 0x69, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x0c,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x49, 0x46, 0x73, 0x12, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x49, 0x46, 0x73, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x63, 0x69, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x12, 0x46, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x49, 0x46, 0x12, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x49, 0x46, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x63, 0x69,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x63, 0x69, 0x66, 0x7d,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x65, 0x67, 0x75, 0x65, 0x7a, 0x7a, 0x2f, 0x6e, 0x69, 0x66, 0x2d, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_nif_proto_rawDescOnce sync.Once
	file_proto_nif_proto_rawDescData = file_proto_nif_proto_rawDesc
)

func file_proto_nif_proto_rawDescGZIP() []byte {
	file_proto_nif_proto_rawDescOnce.Do(func() {
		file_proto_nif_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_nif_proto_rawDescData)
	})
	return file_proto_nif_proto_rawDescData
}

var file_proto_nif_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_nif_proto_goTypes = []interface{}{
	(*NIF)(nil),                    // 0: proto.NIF
	(*NIFs)(nil),                   // 1: proto.NIFs
	(*NIEs)(nil),                   // 2: proto.NIEs
	(*CIF)(nil),                    // 3: proto.CIF
	(*CIFs)(nil),                   // 4: proto.CIFs
	(*BulkParams)(nil),             // 5: proto.BulkParams
	(*Document)(nil),               // 6: proto.Document
	(*ControlDigitResponse)(nil),   // 7: proto.ControlDigitResponse
	(*ParsedDocumentResponse)(nil), // 8: proto.ParsedDocumentResponse
	(*TypeResponse)(nil),           // 9: proto.TypeResponse
	(*NIE)(nil),                    // 10: proto.NIE
	(*IsValid)(nil),                // 11: proto.IsValid
	(*emptypb.Empty)(nil),          // 12: google.protobuf.Empty
}
var file_proto_nif_proto_depIdxs = []int32{
	0,  // 0: proto.NifApi.GetNIFControlDigit:input_type -> proto.NIF
	0,  // 1: proto.NifApi.GetType:input_type -> proto.NIF
	12, // 2: proto.NifApi.GenerateNIF:input_type -> google.protobuf.Empty
	5,  // 3: proto.NifApi.GenerateNIFs:input_type -> proto.BulkParams
	0,  // 4: proto.NifApi.ValidateNIF:input_type -> proto.NIF
	12, // 5: proto.NifApi.GenerateNIE:input_type -> google.protobuf.Empty
	5,  // 6: proto.NifApi.GenerateNIEs:input_type -> proto.BulkParams
	10, // 7: proto.NifApi.ValidateNIE:input_type -> proto.NIE
	6,  // 8: proto.NifApi.GetParsedDocument:input_type -> proto.Document
	6,  // 9: proto.NifApi.ValidateDocument:input_type -> proto.Document
	3,  // 10: proto.NifApi.GetCIFControlDigit:input_type -> proto.CIF
	5,  // 11: proto.NifApi.GenerateCIF:input_type -> proto.BulkParams
	5,  // 12: proto.NifApi.GenerateCIFs:input_type -> proto.BulkParams
	3,  // 13: proto.NifApi.ValidateCIF:input_type -> proto.CIF
	7,  // 14: proto.NifApi.GetNIFControlDigit:output_type -> proto.ControlDigitResponse
	9,  // 15: proto.NifApi.GetType:output_type -> proto.TypeResponse
	0,  // 16: proto.NifApi.GenerateNIF:output_type -> proto.NIF
	1,  // 17: proto.NifApi.GenerateNIFs:output_type -> proto.NIFs
	11, // 18: proto.NifApi.ValidateNIF:output_type -> proto.IsValid
	10, // 19: proto.NifApi.GenerateNIE:output_type -> proto.NIE
	2,  // 20: proto.NifApi.GenerateNIEs:output_type -> proto.NIEs
	11, // 21: proto.NifApi.ValidateNIE:output_type -> proto.IsValid
	8,  // 22: proto.NifApi.GetParsedDocument:output_type -> proto.ParsedDocumentResponse
	11, // 23: proto.NifApi.ValidateDocument:output_type -> proto.IsValid
	7,  // 24: proto.NifApi.GetCIFControlDigit:output_type -> proto.ControlDigitResponse
	3,  // 25: proto.NifApi.GenerateCIF:output_type -> proto.CIF
	4,  // 26: proto.NifApi.GenerateCIFs:output_type -> proto.CIFs
	11, // 27: proto.NifApi.ValidateCIF:output_type -> proto.IsValid
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_nif_proto_init() }
func file_proto_nif_proto_init() {
	if File_proto_nif_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_nif_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NIF); i {
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
		file_proto_nif_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NIFs); i {
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
		file_proto_nif_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NIEs); i {
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
		file_proto_nif_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIF); i {
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
		file_proto_nif_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIFs); i {
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
		file_proto_nif_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkParams); i {
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
		file_proto_nif_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_proto_nif_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlDigitResponse); i {
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
		file_proto_nif_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParsedDocumentResponse); i {
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
		file_proto_nif_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeResponse); i {
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
		file_proto_nif_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NIE); i {
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
		file_proto_nif_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValid); i {
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
	file_proto_nif_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_nif_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_nif_proto_goTypes,
		DependencyIndexes: file_proto_nif_proto_depIdxs,
		MessageInfos:      file_proto_nif_proto_msgTypes,
	}.Build()
	File_proto_nif_proto = out.File
	file_proto_nif_proto_rawDesc = nil
	file_proto_nif_proto_goTypes = nil
	file_proto_nif_proto_depIdxs = nil
}
