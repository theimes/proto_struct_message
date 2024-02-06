// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	structpb "google.golang.org/protobuf/types/known/structpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Field                *structpb.Struct `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebceca9e8703e37f, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetField() *structpb.Struct {
	if m != nil {
		return m.Field
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
}

func init() {
	proto.RegisterFile("message/message.proto", fileDescriptor_ebceca9e8703e37f)
}

var fileDescriptor_ebceca9e8703e37f = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x52, 0x32, 0xe9, 0xf9,
	0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x71, 0x49, 0x51, 0x69, 0x72,
	0x09, 0x44, 0x56, 0xc9, 0x82, 0x8b, 0xdd, 0x17, 0xa2, 0x5c, 0x48, 0x97, 0x8b, 0x35, 0x2d, 0x33,
	0x35, 0x27, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5c, 0x0f, 0xa2, 0x51, 0x0f, 0xa6,
	0x51, 0x2f, 0x18, 0xac, 0x31, 0x08, 0xa2, 0xca, 0x48, 0x93, 0x8b, 0x3d, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0x55, 0x48, 0x8e, 0x8b, 0xcd, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x55, 0x88, 0x43, 0x0f,
	0x6a, 0x9a, 0x14, 0x9c, 0xa5, 0xc4, 0xe0, 0x64, 0x18, 0xa5, 0x9f, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x92, 0x91, 0x9a, 0x99, 0x9b, 0x5a, 0x0c, 0x71, 0x50, 0x3c,
	0xc4, 0x35, 0xf1, 0x68, 0x6e, 0x4f, 0x62, 0x03, 0xcb, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xee, 0x6e, 0xb6, 0xd7, 0xd5, 0x00, 0x00, 0x00,
}
