// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package protobuftest

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

type Person struct {
	// 数字只是标识号
	Id                   *uint32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	IsVip                *bool    `protobuf:"varint,3,req,name=isVip" json:"isVip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_92d71edac9c29a1b, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetIsVip() bool {
	if m != nil && m.IsVip != nil {
		return *m.IsVip
	}
	return false
}

func init() {
	proto.RegisterType((*Person)(nil), "protobuftest.Person")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_92d71edac9c29a1b) }

var fileDescriptor_user_92d71edac9c29a1b = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x53, 0x49, 0xa5, 0x69, 0x25, 0xa9, 0xc5,
	0x25, 0x4a, 0x4e, 0x5c, 0x6c, 0x01, 0xa9, 0x45, 0xc5, 0xf9, 0x79, 0x42, 0x7c, 0x5c, 0x4c, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0xbc, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x4c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x08, 0x17, 0x6b,
	0x66, 0x71, 0x58, 0x66, 0x81, 0x04, 0xb3, 0x02, 0x93, 0x06, 0x47, 0x10, 0x84, 0x03, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x99, 0x67, 0xe2, 0xb3, 0x5e, 0x00, 0x00, 0x00,
}
