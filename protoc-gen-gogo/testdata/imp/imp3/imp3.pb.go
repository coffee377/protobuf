// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imp/imp3/imp3.proto

package imp3 // import "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imp/imp3"

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ForeignImportedMessage struct {
	Tuber                *string  `protobuf:"bytes,1,opt,name=tuber" json:"tuber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForeignImportedMessage) Reset()         { *m = ForeignImportedMessage{} }
func (m *ForeignImportedMessage) String() string { return proto.CompactTextString(m) }
func (*ForeignImportedMessage) ProtoMessage()    {}
func (*ForeignImportedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_imp3_79b1ca30d90e3720, []int{0}
}
func (m *ForeignImportedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForeignImportedMessage.Unmarshal(m, b)
}
func (m *ForeignImportedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForeignImportedMessage.Marshal(b, m, deterministic)
}
func (dst *ForeignImportedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForeignImportedMessage.Merge(dst, src)
}
func (m *ForeignImportedMessage) XXX_Size() int {
	return xxx_messageInfo_ForeignImportedMessage.Size(m)
}
func (m *ForeignImportedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ForeignImportedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ForeignImportedMessage proto.InternalMessageInfo

func (m *ForeignImportedMessage) GetTuber() string {
	if m != nil && m.Tuber != nil {
		return *m.Tuber
	}
	return ""
}

func init() {
	proto.RegisterType((*ForeignImportedMessage)(nil), "imp.ForeignImportedMessage")
}

func init() { proto.RegisterFile("imp/imp3/imp3.proto", fileDescriptor_imp3_79b1ca30d90e3720) }

var fileDescriptor_imp3_79b1ca30d90e3720 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcc, 0x2d, 0xd0,
	0xcf, 0xcc, 0x2d, 0x30, 0x06, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xcc, 0x99, 0xb9,
	0x05, 0x4a, 0x7a, 0x5c, 0x62, 0x6e, 0xf9, 0x45, 0xa9, 0x99, 0xe9, 0x79, 0x9e, 0xb9, 0x05, 0xf9,
	0x45, 0x25, 0xa9, 0x29, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x22, 0x5c, 0xac, 0x25,
	0xa5, 0x49, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x93, 0x4d, 0x94,
	0x55, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe,
	0x3e, 0xd8, 0xb4, 0xa4, 0xd2, 0x34, 0x08, 0x23, 0x59, 0x37, 0x3d, 0x35, 0x4f, 0x17, 0x2c, 0x51,
	0x92, 0x5a, 0x5c, 0x92, 0x92, 0x58, 0x92, 0xa8, 0x0f, 0xb3, 0x1d, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xad, 0x08, 0x34, 0x94, 0x88, 0x00, 0x00, 0x00,
}