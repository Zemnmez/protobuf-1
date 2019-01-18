// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encoding/textpb/testprotos/pb3/test.proto

package pb3

import (
	bytes "bytes"
	gzip "compress/gzip"
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
	Enum_TEN  Enum = 10
)

func (e Enum) Type() protoreflect.EnumType {
	return xxx_Test_protoFile_enumTypes[0]
}
func (e Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var Enum_name = map[int32]string{
	0:  "ZERO",
	1:  "ONE",
	2:  "TWO",
	10: "TEN",
}

var Enum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
	"TWO":  2,
	"TEN":  10,
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422_gzipped, []int{0}
}

type Enums_NestedEnum int32

const (
	Enums_CERO Enums_NestedEnum = 0
	Enums_UNO  Enums_NestedEnum = 1
	Enums_DOS  Enums_NestedEnum = 2
	Enums_DIEZ Enums_NestedEnum = 10
)

func (e Enums_NestedEnum) Type() protoreflect.EnumType {
	return xxx_Test_protoFile_enumTypes[1]
}
func (e Enums_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var Enums_NestedEnum_name = map[int32]string{
	0:  "CERO",
	1:  "UNO",
	2:  "DOS",
	10: "DIEZ",
}

var Enums_NestedEnum_value = map[string]int32{
	"CERO": 0,
	"UNO":  1,
	"DOS":  2,
	"DIEZ": 10,
}

func (x Enums_NestedEnum) String() string {
	return proto.EnumName(Enums_NestedEnum_name, int32(x))
}

func (Enums_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422_gzipped, []int{1, 0}
}

// Scalars contains scalar field types.
type Scalars struct {
	SBool                bool     `protobuf:"varint,1,opt,name=s_bool,json=sBool,proto3" json:"s_bool,omitempty"`
	SInt32               int32    `protobuf:"varint,2,opt,name=s_int32,json=sInt32,proto3" json:"s_int32,omitempty"`
	SInt64               int64    `protobuf:"varint,3,opt,name=s_int64,json=sInt64,proto3" json:"s_int64,omitempty"`
	SUint32              uint32   `protobuf:"varint,4,opt,name=s_uint32,json=sUint32,proto3" json:"s_uint32,omitempty"`
	SUint64              uint64   `protobuf:"varint,5,opt,name=s_uint64,json=sUint64,proto3" json:"s_uint64,omitempty"`
	SSint32              int32    `protobuf:"zigzag32,6,opt,name=s_sint32,json=sSint32,proto3" json:"s_sint32,omitempty"`
	SSint64              int64    `protobuf:"zigzag64,7,opt,name=s_sint64,json=sSint64,proto3" json:"s_sint64,omitempty"`
	SFixed32             uint32   `protobuf:"fixed32,8,opt,name=s_fixed32,json=sFixed32,proto3" json:"s_fixed32,omitempty"`
	SFixed64             uint64   `protobuf:"fixed64,9,opt,name=s_fixed64,json=sFixed64,proto3" json:"s_fixed64,omitempty"`
	SSfixed32            int32    `protobuf:"fixed32,10,opt,name=s_sfixed32,json=sSfixed32,proto3" json:"s_sfixed32,omitempty"`
	SSfixed64            int64    `protobuf:"fixed64,11,opt,name=s_sfixed64,json=sSfixed64,proto3" json:"s_sfixed64,omitempty"`
	SFloat               float32  `protobuf:"fixed32,20,opt,name=s_float,json=sFloat,proto3" json:"s_float,omitempty"`
	SDouble              float64  `protobuf:"fixed64,21,opt,name=s_double,json=sDouble,proto3" json:"s_double,omitempty"`
	SBytes               []byte   `protobuf:"bytes,14,opt,name=s_bytes,json=sBytes,proto3" json:"s_bytes,omitempty"`
	SString              string   `protobuf:"bytes,13,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scalars) ProtoReflect() protoreflect.Message {
	return xxx_Test_protoFile_messageTypes[0].MessageOf(m)
}
func (m *Scalars) Reset()         { *m = Scalars{} }
func (m *Scalars) String() string { return proto.CompactTextString(m) }
func (*Scalars) ProtoMessage()    {}
func (*Scalars) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422_gzipped, []int{0}
}

func (m *Scalars) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scalars.Unmarshal(m, b)
}
func (m *Scalars) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scalars.Marshal(b, m, deterministic)
}
func (m *Scalars) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scalars.Merge(m, src)
}
func (m *Scalars) XXX_Size() int {
	return xxx_messageInfo_Scalars.Size(m)
}
func (m *Scalars) XXX_DiscardUnknown() {
	xxx_messageInfo_Scalars.DiscardUnknown(m)
}

var xxx_messageInfo_Scalars proto.InternalMessageInfo

func (m *Scalars) GetSBool() bool {
	if m != nil {
		return m.SBool
	}
	return false
}

func (m *Scalars) GetSInt32() int32 {
	if m != nil {
		return m.SInt32
	}
	return 0
}

func (m *Scalars) GetSInt64() int64 {
	if m != nil {
		return m.SInt64
	}
	return 0
}

func (m *Scalars) GetSUint32() uint32 {
	if m != nil {
		return m.SUint32
	}
	return 0
}

func (m *Scalars) GetSUint64() uint64 {
	if m != nil {
		return m.SUint64
	}
	return 0
}

func (m *Scalars) GetSSint32() int32 {
	if m != nil {
		return m.SSint32
	}
	return 0
}

func (m *Scalars) GetSSint64() int64 {
	if m != nil {
		return m.SSint64
	}
	return 0
}

func (m *Scalars) GetSFixed32() uint32 {
	if m != nil {
		return m.SFixed32
	}
	return 0
}

func (m *Scalars) GetSFixed64() uint64 {
	if m != nil {
		return m.SFixed64
	}
	return 0
}

func (m *Scalars) GetSSfixed32() int32 {
	if m != nil {
		return m.SSfixed32
	}
	return 0
}

func (m *Scalars) GetSSfixed64() int64 {
	if m != nil {
		return m.SSfixed64
	}
	return 0
}

func (m *Scalars) GetSFloat() float32 {
	if m != nil {
		return m.SFloat
	}
	return 0
}

func (m *Scalars) GetSDouble() float64 {
	if m != nil {
		return m.SDouble
	}
	return 0
}

func (m *Scalars) GetSBytes() []byte {
	if m != nil {
		return m.SBytes
	}
	return nil
}

func (m *Scalars) GetSString() string {
	if m != nil {
		return m.SString
	}
	return ""
}

// Message contains enum fields.
type Enums struct {
	SEnum                Enum             `protobuf:"varint,1,opt,name=s_enum,json=sEnum,proto3,enum=pb3.Enum" json:"s_enum,omitempty"`
	SNestedEnum          Enums_NestedEnum `protobuf:"varint,3,opt,name=s_nested_enum,json=sNestedEnum,proto3,enum=pb3.Enums_NestedEnum" json:"s_nested_enum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Enums) ProtoReflect() protoreflect.Message {
	return xxx_Test_protoFile_messageTypes[1].MessageOf(m)
}
func (m *Enums) Reset()         { *m = Enums{} }
func (m *Enums) String() string { return proto.CompactTextString(m) }
func (*Enums) ProtoMessage()    {}
func (*Enums) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422_gzipped, []int{1}
}

func (m *Enums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enums.Unmarshal(m, b)
}
func (m *Enums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enums.Marshal(b, m, deterministic)
}
func (m *Enums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enums.Merge(m, src)
}
func (m *Enums) XXX_Size() int {
	return xxx_messageInfo_Enums.Size(m)
}
func (m *Enums) XXX_DiscardUnknown() {
	xxx_messageInfo_Enums.DiscardUnknown(m)
}

var xxx_messageInfo_Enums proto.InternalMessageInfo

func (m *Enums) GetSEnum() Enum {
	if m != nil {
		return m.SEnum
	}
	return Enum_ZERO
}

func (m *Enums) GetSNestedEnum() Enums_NestedEnum {
	if m != nil {
		return m.SNestedEnum
	}
	return Enums_CERO
}

// Message contains nested message field.
type Nests struct {
	SNested              *Nested  `protobuf:"bytes,1,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nests) ProtoReflect() protoreflect.Message {
	return xxx_Test_protoFile_messageTypes[2].MessageOf(m)
}
func (m *Nests) Reset()         { *m = Nests{} }
func (m *Nests) String() string { return proto.CompactTextString(m) }
func (*Nests) ProtoMessage()    {}
func (*Nests) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422_gzipped, []int{2}
}

func (m *Nests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nests.Unmarshal(m, b)
}
func (m *Nests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nests.Marshal(b, m, deterministic)
}
func (m *Nests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nests.Merge(m, src)
}
func (m *Nests) XXX_Size() int {
	return xxx_messageInfo_Nests.Size(m)
}
func (m *Nests) XXX_DiscardUnknown() {
	xxx_messageInfo_Nests.DiscardUnknown(m)
}

var xxx_messageInfo_Nests proto.InternalMessageInfo

func (m *Nests) GetSNested() *Nested {
	if m != nil {
		return m.SNested
	}
	return nil
}

// Message type used as submessage.
type Nested struct {
	SString              string   `protobuf:"bytes,1,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	SNested              *Nested  `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) ProtoReflect() protoreflect.Message {
	return xxx_Test_protoFile_messageTypes[3].MessageOf(m)
}
func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_0854715c5b41c422_gzipped, []int{3}
}

func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetSString() string {
	if m != nil {
		return m.SString
	}
	return ""
}

func (m *Nested) GetSNested() *Nested {
	if m != nil {
		return m.SNested
	}
	return nil
}

func init() {
	proto.RegisterFile("encoding/textpb/testprotos/pb3/test.proto", fileDescriptor_0854715c5b41c422_gzipped)
	proto.RegisterEnum("pb3.Enum", Enum_name, Enum_value)
	proto.RegisterEnum("pb3.Enums_NestedEnum", Enums_NestedEnum_name, Enums_NestedEnum_value)
	proto.RegisterType((*Scalars)(nil), "pb3.Scalars")
	proto.RegisterType((*Enums)(nil), "pb3.Enums")
	proto.RegisterType((*Nests)(nil), "pb3.Nests")
	proto.RegisterType((*Nested)(nil), "pb3.Nested")
}

var fileDescriptor_0854715c5b41c422 = []byte{
	// 863 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x29, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x70,
	0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x33,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x62, 0x33,
	0x22, 0x9e, 0x03, 0x0a, 0x07, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x42,
	0x6f, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x73, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x73,
	0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x07, 0x73, 0x53, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x07, 0x52, 0x08, 0x73, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x06, 0x52, 0x08, 0x73, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0f, 0x52,
	0x09, 0x73, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x5f,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x10, 0x52, 0x09,
	0x73, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x22, 0x98, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x06, 0x73,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x70, 0x62,
	0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x39, 0x0a,
	0x0d, 0x73, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0b, 0x73, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x32, 0x0a, 0x0a, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x45, 0x52, 0x4f, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x55, 0x4e, 0x4f, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x4f, 0x53,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x45, 0x5a, 0x10, 0x0a, 0x22, 0x2f, 0x0a, 0x05,
	0x4e, 0x65, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x73, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x52, 0x07, 0x73, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0x4b, 0x0a,
	0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x26, 0x0a, 0x08, 0x73, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x52, 0x07, 0x73, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2a, 0x2b, 0x0a, 0x04, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x45, 0x4e, 0x10, 0x0a, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var fileDescriptor_0854715c5b41c422_gzipped = func() []byte {
	bb := new(bytes.Buffer)
	zw, _ := gzip.NewWriterLevel(bb, gzip.NoCompression)
	zw.Write(fileDescriptor_0854715c5b41c422)
	zw.Close()
	return bb.Bytes()
}()

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Test_protoFile protoreflect.FileDescriptor

var xxx_Test_protoFile_enumTypes [2]protoreflect.EnumType
var xxx_Test_protoFile_messageTypes [4]protoimpl.MessageType
var xxx_Test_protoFile_goTypes = []interface{}{
	(Enum)(0),             // 0: pb3.Enum
	(Enums_NestedEnum)(0), // 1: pb3.Enums.NestedEnum
	(*Scalars)(nil),       // 2: pb3.Scalars
	(*Enums)(nil),         // 3: pb3.Enums
	(*Nests)(nil),         // 4: pb3.Nests
	(*Nested)(nil),        // 5: pb3.Nested
}
var xxx_Test_protoFile_depIdxs = []int32{
	0, // pb3.Enums.s_enum:type_name -> pb3.Enum
	1, // pb3.Enums.s_nested_enum:type_name -> pb3.Enums.NestedEnum
	5, // pb3.Nests.s_nested:type_name -> pb3.Nested
	5, // pb3.Nested.s_nested:type_name -> pb3.Nested
}

func init() {
	var messageTypes [4]protoreflect.MessageType
	Test_protoFile = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_0854715c5b41c422,
		GoTypes:            xxx_Test_protoFile_goTypes,
		DependencyIndexes:  xxx_Test_protoFile_depIdxs,
		EnumOutputTypes:    xxx_Test_protoFile_enumTypes[:],
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_Test_protoFile_goTypes[2:][:4]
	for i, mt := range messageTypes[:] {
		xxx_Test_protoFile_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_Test_protoFile_messageTypes[i].PBType = mt
	}
	xxx_Test_protoFile_goTypes = nil
	xxx_Test_protoFile_depIdxs = nil
}