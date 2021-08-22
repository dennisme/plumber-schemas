// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rabbit.proto

package backends

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

type Rabbit struct {
	// Required for reading and writing
	ExchangeName string `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty"`
	// Reading only
	// Queue name to read from
	QueueName string `protobuf:"bytes,2,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	// Routing key for topic based exchanges
	RoutingKey string `protobuf:"bytes,3,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty"`
	// Reading only
	// Whether the queue is exclusive to this connection
	QueueExclusive bool `protobuf:"varint,4,opt,name=queue_exclusive,json=queueExclusive,proto3" json:"queue_exclusive,omitempty"`
	// Reading only
	// Whether to create the queue
	QueueDeclare bool `protobuf:"varint,5,opt,name=queue_declare,json=queueDeclare,proto3" json:"queue_declare,omitempty"`
	// Reading only
	// Whether the queue should survive after disconnect
	QueueDurable bool `protobuf:"varint,6,opt,name=queue_durable,json=queueDurable,proto3" json:"queue_durable,omitempty"`
	// Reading only
	AutoAck bool `protobuf:"varint,7,opt,name=auto_ack,json=autoAck,proto3" json:"auto_ack,omitempty"`
	// Reading only
	AppId                string   `protobuf:"bytes,8,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rabbit) Reset()         { *m = Rabbit{} }
func (m *Rabbit) String() string { return proto.CompactTextString(m) }
func (*Rabbit) ProtoMessage()    {}
func (*Rabbit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e2cf5cfff1350, []int{0}
}

func (m *Rabbit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rabbit.Unmarshal(m, b)
}
func (m *Rabbit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rabbit.Marshal(b, m, deterministic)
}
func (m *Rabbit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rabbit.Merge(m, src)
}
func (m *Rabbit) XXX_Size() int {
	return xxx_messageInfo_Rabbit.Size(m)
}
func (m *Rabbit) XXX_DiscardUnknown() {
	xxx_messageInfo_Rabbit.DiscardUnknown(m)
}

var xxx_messageInfo_Rabbit proto.InternalMessageInfo

func (m *Rabbit) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *Rabbit) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *Rabbit) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func (m *Rabbit) GetQueueExclusive() bool {
	if m != nil {
		return m.QueueExclusive
	}
	return false
}

func (m *Rabbit) GetQueueDeclare() bool {
	if m != nil {
		return m.QueueDeclare
	}
	return false
}

func (m *Rabbit) GetQueueDurable() bool {
	if m != nil {
		return m.QueueDurable
	}
	return false
}

func (m *Rabbit) GetAutoAck() bool {
	if m != nil {
		return m.AutoAck
	}
	return false
}

func (m *Rabbit) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func init() {
	proto.RegisterType((*Rabbit)(nil), "protos.backends.Rabbit")
}

func init() { proto.RegisterFile("rabbit.proto", fileDescriptor_9a3e2cf5cfff1350) }

var fileDescriptor_9a3e2cf5cfff1350 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x69, 0xb5, 0xd3, 0x36, 0x56, 0x0b, 0x01, 0x21, 0x2e, 0xc4, 0xa2, 0x0b, 0xbb, 0xb1,
	0x59, 0xb8, 0x16, 0x51, 0x74, 0x21, 0x82, 0x8b, 0x59, 0xba, 0x19, 0x6e, 0x32, 0x97, 0x99, 0x30,
	0x3f, 0x89, 0x99, 0x44, 0xda, 0xd7, 0xf0, 0x89, 0xa5, 0x77, 0x3a, 0xd0, 0x55, 0xc8, 0x77, 0xbe,
	0x0b, 0x87, 0xc3, 0x16, 0x1e, 0x94, 0x32, 0x61, 0xe3, 0xbc, 0x0d, 0x96, 0x2f, 0xe9, 0xe9, 0x36,
	0x0a, 0x74, 0x85, 0x6d, 0xde, 0xdd, 0xfe, 0x8d, 0x59, 0x92, 0x92, 0xc1, 0xef, 0xd8, 0x39, 0x6e,
	0x75, 0x09, 0x6d, 0x81, 0x59, 0x0b, 0x0d, 0x8a, 0xd1, 0x6a, 0xb4, 0x9e, 0xa7, 0x8b, 0x01, 0x7e,
	0x41, 0x83, 0xfc, 0x9a, 0xb1, 0x9f, 0x88, 0xf1, 0x60, 0x8c, 0xc9, 0x98, 0x13, 0xa1, 0xf8, 0x86,
	0x9d, 0x79, 0x1b, 0x83, 0x69, 0x8b, 0xac, 0xc2, 0x9d, 0x38, 0xa1, 0x9c, 0x1d, 0xd0, 0x27, 0xee,
	0xf8, 0x3d, 0x5b, 0xf6, 0xf7, 0xb8, 0xd5, 0x75, 0xec, 0xcc, 0x2f, 0x8a, 0xd3, 0xd5, 0x68, 0x3d,
	0x4b, 0x2f, 0x08, 0xbf, 0x0f, 0x74, 0xdf, 0xa6, 0x17, 0x73, 0xd4, 0x35, 0x78, 0x14, 0x13, 0xd2,
	0x16, 0x04, 0xdf, 0x7a, 0x76, 0x24, 0x45, 0x0f, 0xaa, 0x46, 0x91, 0x1c, 0x4b, 0x3d, 0xe3, 0x57,
	0x6c, 0x06, 0x31, 0xd8, 0x0c, 0x74, 0x25, 0xa6, 0x94, 0x4f, 0xf7, 0xff, 0x17, 0x5d, 0xf1, 0x4b,
	0x96, 0x80, 0x73, 0x99, 0xc9, 0xc5, 0x8c, 0x9a, 0x4e, 0xc0, 0xb9, 0x8f, 0xfc, 0xf5, 0xf9, 0xfb,
	0xa9, 0x30, 0xa1, 0x8c, 0x6a, 0xa3, 0x6d, 0x23, 0x15, 0x04, 0x5d, 0x6a, 0xeb, 0x9d, 0x74, 0x75,
	0x6c, 0x14, 0xfa, 0x87, 0x4e, 0x97, 0xd8, 0x40, 0x27, 0x55, 0x34, 0x75, 0x2e, 0x0b, 0x2b, 0xfb,
	0x55, 0xe5, 0xb0, 0xaa, 0x4a, 0x08, 0x3c, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xfc, 0xbd,
	0xc6, 0x7d, 0x01, 0x00, 0x00,
}
