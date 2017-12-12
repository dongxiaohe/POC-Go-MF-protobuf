// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zendesk/protobuf.proto

/*
Package com_zendesk_protobuf is a generated protocol buffer package.

It is generated from these files:
	zendesk/protobuf.proto

It has these top-level messages:
	ProtobufHeader
	TypedEvent
*/
package com_zendesk_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtobufHeader_SchemaFormat int32

const (
	ProtobufHeader_UNKNOWN ProtobufHeader_SchemaFormat = 0
	ProtobufHeader_PATH    ProtobufHeader_SchemaFormat = 1
)

var ProtobufHeader_SchemaFormat_name = map[int32]string{
	0: "UNKNOWN",
	1: "PATH",
}
var ProtobufHeader_SchemaFormat_value = map[string]int32{
	"UNKNOWN": 0,
	"PATH":    1,
}

func (x ProtobufHeader_SchemaFormat) String() string {
	return proto.EnumName(ProtobufHeader_SchemaFormat_name, int32(x))
}
func (ProtobufHeader_SchemaFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type ProtobufHeader struct {
	SchemaFormat ProtobufHeader_SchemaFormat `protobuf:"varint,1,opt,name=schema_format,json=schemaFormat,enum=com.zendesk.protobuf.ProtobufHeader_SchemaFormat" json:"schema_format,omitempty"`
	SchemaPath   string                      `protobuf:"bytes,2,opt,name=schema_path,json=schemaPath" json:"schema_path,omitempty"`
}

func (m *ProtobufHeader) Reset()                    { *m = ProtobufHeader{} }
func (m *ProtobufHeader) String() string            { return proto.CompactTextString(m) }
func (*ProtobufHeader) ProtoMessage()               {}
func (*ProtobufHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtobufHeader) GetSchemaFormat() ProtobufHeader_SchemaFormat {
	if m != nil {
		return m.SchemaFormat
	}
	return ProtobufHeader_UNKNOWN
}

func (m *ProtobufHeader) GetSchemaPath() string {
	if m != nil {
		return m.SchemaPath
	}
	return ""
}

// Supertype of all zendesk toplevel protobuf messages
type TypedEvent struct {
	Header *ProtobufHeader `protobuf:"bytes,2040,opt,name=header" json:"header,omitempty"`
}

func (m *TypedEvent) Reset()                    { *m = TypedEvent{} }
func (m *TypedEvent) String() string            { return proto.CompactTextString(m) }
func (*TypedEvent) ProtoMessage()               {}
func (*TypedEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TypedEvent) GetHeader() *ProtobufHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtobufHeader)(nil), "com.zendesk.protobuf.ProtobufHeader")
	proto.RegisterType((*TypedEvent)(nil), "com.zendesk.protobuf.TypedEvent")
	proto.RegisterEnum("com.zendesk.protobuf.ProtobufHeader_SchemaFormat", ProtobufHeader_SchemaFormat_name, ProtobufHeader_SchemaFormat_value)
}

func init() { proto.RegisterFile("zendesk/protobuf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xab, 0x4a, 0xcd, 0x4b,
	0x49, 0x2d, 0xce, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x03, 0x33, 0x84,
	0x44, 0x92, 0xf3, 0x73, 0xf5, 0xa0, 0x72, 0x7a, 0x30, 0x39, 0xa5, 0x0d, 0x8c, 0x5c, 0x7c, 0x01,
	0x50, 0x8e, 0x47, 0x6a, 0x62, 0x4a, 0x6a, 0x91, 0x50, 0x18, 0x17, 0x6f, 0x71, 0x72, 0x46, 0x6a,
	0x6e, 0x62, 0x7c, 0x5a, 0x7e, 0x51, 0x6e, 0x62, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f, 0x91,
	0xa1, 0x1e, 0x36, 0x03, 0xf4, 0x50, 0x35, 0xeb, 0x05, 0x83, 0x75, 0xba, 0x81, 0x35, 0x06, 0xf1,
	0x14, 0x23, 0xf1, 0x84, 0xe4, 0xb9, 0xb8, 0xa1, 0xe6, 0x16, 0x24, 0x96, 0x64, 0x48, 0x30, 0x29,
	0x30, 0x6a, 0x70, 0x06, 0x71, 0x41, 0x84, 0x02, 0x12, 0x4b, 0x32, 0x94, 0x54, 0xb9, 0x78, 0x90,
	0xb5, 0x0b, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x08, 0x71,
	0x70, 0xb1, 0x04, 0x38, 0x86, 0x78, 0x08, 0x30, 0x2a, 0x79, 0x73, 0x71, 0x85, 0x54, 0x16, 0xa4,
	0xa6, 0xb8, 0x96, 0xa5, 0xe6, 0x95, 0x08, 0xd9, 0x72, 0xb1, 0x65, 0x80, 0xad, 0x96, 0xf8, 0xc1,
	0xaf, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0x42, 0x8c, 0x3b, 0x83, 0xa0, 0x9a, 0x92, 0xd8, 0xc0, 0x0a,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x94, 0x07, 0xb6, 0x36, 0x01, 0x00, 0x00,
}
