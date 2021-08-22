// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aws-sqs.proto

package conns

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AWSSQS struct {
	// Required
	AwsRegion string `protobuf:"bytes,1,opt,name=aws_region,json=awsRegion,proto3" json:"aws_region,omitempty"`
	// Required
	AwsAccessKeyId string `protobuf:"bytes,2,opt,name=aws_access_key_id,json=awsAccessKeyId,proto3" json:"aws_access_key_id,omitempty"`
	// Required
	AwsSecretAccessKey   string   `protobuf:"bytes,3,opt,name=aws_secret_access_key,json=awsSecretAccessKey,proto3" json:"aws_secret_access_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSSQS) Reset()         { *m = AWSSQS{} }
func (m *AWSSQS) String() string { return proto.CompactTextString(m) }
func (*AWSSQS) ProtoMessage()    {}
func (*AWSSQS) Descriptor() ([]byte, []int) {
	return fileDescriptor_07517116f22c3664, []int{0}
}

func (m *AWSSQS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSSQS.Unmarshal(m, b)
}
func (m *AWSSQS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSSQS.Marshal(b, m, deterministic)
}
func (m *AWSSQS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSSQS.Merge(m, src)
}
func (m *AWSSQS) XXX_Size() int {
	return xxx_messageInfo_AWSSQS.Size(m)
}
func (m *AWSSQS) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSSQS.DiscardUnknown(m)
}

var xxx_messageInfo_AWSSQS proto.InternalMessageInfo

func (m *AWSSQS) GetAwsRegion() string {
	if m != nil {
		return m.AwsRegion
	}
	return ""
}

func (m *AWSSQS) GetAwsAccessKeyId() string {
	if m != nil {
		return m.AwsAccessKeyId
	}
	return ""
}

func (m *AWSSQS) GetAwsSecretAccessKey() string {
	if m != nil {
		return m.AwsSecretAccessKey
	}
	return ""
}

func init() {
	proto.RegisterType((*AWSSQS)(nil), "protos.conns.AWSSQS")
}

func init() { proto.RegisterFile("aws-sqs.proto", fileDescriptor_07517116f22c3664) }

var fileDescriptor_07517116f22c3664 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0xc2, 0xc1, 0x2d, 0x2a, 0xb8, 0x20, 0xa4, 0x11, 0xc4, 0x4a, 0x8b, 0xcb, 0x22,
	0x76, 0x62, 0x73, 0x76, 0x62, 0xe5, 0x6d, 0x21, 0xd8, 0x84, 0xd9, 0xc9, 0x90, 0x2c, 0x26, 0xd9,
	0xb8, 0xb3, 0x61, 0xc9, 0x0f, 0xf0, 0x7f, 0x4b, 0x46, 0x90, 0xab, 0x86, 0x79, 0xef, 0x7b, 0xc5,
	0xa7, 0xce, 0x21, 0xf3, 0x8e, 0xbf, 0xb9, 0x9a, 0x62, 0x48, 0x41, 0x9f, 0xc9, 0xe1, 0x0a, 0xc3,
	0x38, 0xf2, 0xed, 0x4f, 0xa1, 0x36, 0xfb, 0x0f, 0x6b, 0xdf, 0xad, 0xbe, 0x56, 0x0a, 0x32, 0xd7,
	0x91, 0x5a, 0x1f, 0xc6, 0xb2, 0xb8, 0x29, 0xee, 0xb6, 0x87, 0x2d, 0x64, 0x3e, 0x48, 0xa0, 0xef,
	0xd5, 0xe5, 0x5a, 0x03, 0x22, 0x31, 0xd7, 0x5f, 0xb4, 0xd4, 0xbe, 0x29, 0x4f, 0x84, 0xba, 0x80,
	0xcc, 0x7b, 0xc9, 0xdf, 0x68, 0x79, 0x6d, 0xf4, 0x83, 0xba, 0x5a, 0x51, 0x26, 0x8c, 0x94, 0x8e,
	0x16, 0xe5, 0xa9, 0xe0, 0x1a, 0x32, 0x5b, 0xe9, 0xfe, 0x47, 0x2f, 0xcf, 0x9f, 0x4f, 0xad, 0x4f,
	0xdd, 0xec, 0x2a, 0x0c, 0x83, 0x71, 0x90, 0xb0, 0xc3, 0x10, 0x27, 0x33, 0xf5, 0xf3, 0xe0, 0x28,
	0xee, 0x18, 0x3b, 0x1a, 0x80, 0x8d, 0x9b, 0x7d, 0xdf, 0x98, 0x36, 0x98, 0x3f, 0x0b, 0x23, 0x16,
	0x6e, 0x23, 0xdf, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x36, 0xf1, 0x10, 0xeb, 0x00,
	0x00, 0x00,
}
