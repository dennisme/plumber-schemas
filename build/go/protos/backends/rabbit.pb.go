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

type RabbitConn struct {
	// @gotags: kong:"help='Destination host address (full DSN)',env='PLUMBER_RELAY_RABBIT_ADDRESS',default='amqp://localhost',required"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='Destination host address (full DSN)',env='PLUMBER_RELAY_RABBIT_ADDRESS',default='amqp://localhost',required"`
	// @gotags: kong:"help='Force TLS usage (regardless of DSN)',env='PLUMBER_RELAY_RABBIT_USE_TLS'"
	UseTls bool `protobuf:"varint,2,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Force TLS usage (regardless of DSN)',env='PLUMBER_RELAY_RABBIT_USE_TLS'"`
	// @gotags: kong:"help='Whether to verify server TLS certificate',env='PLUMBER_RELAY_RABBIT_SKIP_VERIFY_TLS'"
	InsecureTls          bool     `protobuf:"varint,3,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty" kong:"help='Whether to verify server TLS certificate',env='PLUMBER_RELAY_RABBIT_SKIP_VERIFY_TLS'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitConn) Reset()         { *m = RabbitConn{} }
func (m *RabbitConn) String() string { return proto.CompactTextString(m) }
func (*RabbitConn) ProtoMessage()    {}
func (*RabbitConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e2cf5cfff1350, []int{0}
}

func (m *RabbitConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitConn.Unmarshal(m, b)
}
func (m *RabbitConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitConn.Marshal(b, m, deterministic)
}
func (m *RabbitConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitConn.Merge(m, src)
}
func (m *RabbitConn) XXX_Size() int {
	return xxx_messageInfo_RabbitConn.Size(m)
}
func (m *RabbitConn) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitConn.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitConn proto.InternalMessageInfo

func (m *RabbitConn) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RabbitConn) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *RabbitConn) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

type RabbitReadArgs struct {
	// @gotags: kong:"help='Name of the exchange',env='PLUMBER_RELAY_RABBIT_EXCHANGE',required"
	ExchangeName string `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty" kong:"help='Name of the exchange',env='PLUMBER_RELAY_RABBIT_EXCHANGE',required"`
	// @gotags: kong:"help='Name of the queue where messages will be routed to',env='PLUMBER_RELAY_RABBIT_QUEUE',required"
	QueueName string `protobuf:"bytes,2,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty" kong:"help='Name of the queue where messages will be routed to',env='PLUMBER_RELAY_RABBIT_QUEUE',required"`
	// @gotags: kong:"help='Binding key for topic based exchanges',env='PLUMBER_RELAY_RABBIT_ROUTING_KEY',required"
	BindingKey string `protobuf:"bytes,3,opt,name=binding_key,json=bindingKey,proto3" json:"binding_key,omitempty" kong:"help='Binding key for topic based exchanges',env='PLUMBER_RELAY_RABBIT_ROUTING_KEY',required"`
	// @gotags: kong:"help='Whether plumber should be the only one using the queue',env='PLUMBER_RELAY_RABBIT_QUEUE_EXCLUSIVE'"
	QueueExclusive bool `protobuf:"varint,4,opt,name=queue_exclusive,json=queueExclusive,proto3" json:"queue_exclusive,omitempty" kong:"help='Whether plumber should be the only one using the queue',env='PLUMBER_RELAY_RABBIT_QUEUE_EXCLUSIVE'"`
	// @gotags: kong:"help='Whether to create/declare the queue (if it does not exist)',env='PLUMBER_RELAY_RABBIT_QUEUE_DECLARE',default=true"
	QueueDeclare bool `protobuf:"varint,5,opt,name=queue_declare,json=queueDeclare,proto3" json:"queue_declare,omitempty" kong:"help='Whether to create/declare the queue (if it does not exist)',env='PLUMBER_RELAY_RABBIT_QUEUE_DECLARE',default=true"`
	// @gotags: kong:"help='Whether the queue should survive after disconnect',env='PLUMBER_RELAY_RABBIT_QUEUE_DURABLE'"
	QueueDurable bool `protobuf:"varint,6,opt,name=queue_durable,json=queueDurable,proto3" json:"queue_durable,omitempty" kong:"help='Whether the queue should survive after disconnect',env='PLUMBER_RELAY_RABBIT_QUEUE_DURABLE'"`
	// @gotags: kong:"help='Automatically acknowledge receipt of read/received messages',env='PLUMBER_RELAY_RABBIT_AUTOACK',default=true"
	AutoAck bool `protobuf:"varint,7,opt,name=auto_ack,json=autoAck,proto3" json:"auto_ack,omitempty" kong:"help='Automatically acknowledge receipt of read/received messages',env='PLUMBER_RELAY_RABBIT_AUTOACK',default=true"`
	// @gotags: kong:"help='How to identify the consumer to RabbitMQ',env='PLUMBER_RELAY_CONSUMER_TAG',default=plumber"
	ConsumerTag string `protobuf:"bytes,8,opt,name=consumer_tag,json=consumerTag,proto3" json:"consumer_tag,omitempty" kong:"help='How to identify the consumer to RabbitMQ',env='PLUMBER_RELAY_CONSUMER_TAG',default=plumber"`
	// @gotags: kong:"help='Whether to auto-delete the queue after plumber has disconnected',env='PLUMBER_RELAY_RABBIT_QUEUE_AUTO_DELETE',default=true"
	QueueDelete          bool     `protobuf:"varint,9,opt,name=queue_delete,json=queueDelete,proto3" json:"queue_delete,omitempty" kong:"help='Whether to auto-delete the queue after plumber has disconnected',env='PLUMBER_RELAY_RABBIT_QUEUE_AUTO_DELETE',default=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitReadArgs) Reset()         { *m = RabbitReadArgs{} }
func (m *RabbitReadArgs) String() string { return proto.CompactTextString(m) }
func (*RabbitReadArgs) ProtoMessage()    {}
func (*RabbitReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e2cf5cfff1350, []int{1}
}

func (m *RabbitReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitReadArgs.Unmarshal(m, b)
}
func (m *RabbitReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitReadArgs.Marshal(b, m, deterministic)
}
func (m *RabbitReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitReadArgs.Merge(m, src)
}
func (m *RabbitReadArgs) XXX_Size() int {
	return xxx_messageInfo_RabbitReadArgs.Size(m)
}
func (m *RabbitReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitReadArgs proto.InternalMessageInfo

func (m *RabbitReadArgs) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *RabbitReadArgs) GetBindingKey() string {
	if m != nil {
		return m.BindingKey
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueExclusive() bool {
	if m != nil {
		return m.QueueExclusive
	}
	return false
}

func (m *RabbitReadArgs) GetQueueDeclare() bool {
	if m != nil {
		return m.QueueDeclare
	}
	return false
}

func (m *RabbitReadArgs) GetQueueDurable() bool {
	if m != nil {
		return m.QueueDurable
	}
	return false
}

func (m *RabbitReadArgs) GetAutoAck() bool {
	if m != nil {
		return m.AutoAck
	}
	return false
}

func (m *RabbitReadArgs) GetConsumerTag() string {
	if m != nil {
		return m.ConsumerTag
	}
	return ""
}

func (m *RabbitReadArgs) GetQueueDelete() bool {
	if m != nil {
		return m.QueueDelete
	}
	return false
}

type RabbitWriteArgs struct {
	// @gotags: kong:"help='Exchange to write message(s) to',required"
	ExchangeName string `protobuf:"bytes,1,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty" kong:"help='Exchange to write message(s) to',required"`
	// @gotags: kong:"help='Routing key to write message(s) to',required"
	RoutingKey string `protobuf:"bytes,2,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty" kong:"help='Routing key to write message(s) to',required"`
	// @gotags: kong:"help='Fills message properties $app_id with this value',default=plumber"
	AppId                string   `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" kong:"help='Fills message properties  with this value',default=plumber"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RabbitWriteArgs) Reset()         { *m = RabbitWriteArgs{} }
func (m *RabbitWriteArgs) String() string { return proto.CompactTextString(m) }
func (*RabbitWriteArgs) ProtoMessage()    {}
func (*RabbitWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e2cf5cfff1350, []int{2}
}

func (m *RabbitWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RabbitWriteArgs.Unmarshal(m, b)
}
func (m *RabbitWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RabbitWriteArgs.Marshal(b, m, deterministic)
}
func (m *RabbitWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RabbitWriteArgs.Merge(m, src)
}
func (m *RabbitWriteArgs) XXX_Size() int {
	return xxx_messageInfo_RabbitWriteArgs.Size(m)
}
func (m *RabbitWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RabbitWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RabbitWriteArgs proto.InternalMessageInfo

func (m *RabbitWriteArgs) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *RabbitWriteArgs) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func (m *RabbitWriteArgs) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func init() {
	proto.RegisterType((*RabbitConn)(nil), "protos.backends.RabbitConn")
	proto.RegisterType((*RabbitReadArgs)(nil), "protos.backends.RabbitReadArgs")
	proto.RegisterType((*RabbitWriteArgs)(nil), "protos.backends.RabbitWriteArgs")
}

func init() { proto.RegisterFile("rabbit.proto", fileDescriptor_9a3e2cf5cfff1350) }

var fileDescriptor_9a3e2cf5cfff1350 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x8b, 0x14, 0x31,
	0x10, 0x86, 0x99, 0xd1, 0x9d, 0x8f, 0x9a, 0x71, 0x07, 0x1a, 0xc4, 0x78, 0x90, 0x5d, 0xc7, 0x83,
	0x7b, 0x71, 0xfa, 0xe0, 0x59, 0x64, 0xfd, 0x38, 0x88, 0xe0, 0xa1, 0x19, 0x10, 0xbc, 0x34, 0x95,
	0xa4, 0xe8, 0x09, 0x93, 0x4e, 0xda, 0x7c, 0xc8, 0xee, 0x4f, 0xf4, 0x5f, 0x49, 0x27, 0x1d, 0x98,
	0xa3, 0xa7, 0x26, 0x4f, 0xbd, 0x4d, 0xbd, 0x3c, 0x14, 0x6c, 0x1d, 0x72, 0xae, 0xc2, 0x61, 0x70,
	0x36, 0xd8, 0x6a, 0x97, 0x3e, 0xfe, 0xc0, 0x51, 0x9c, 0xc9, 0x48, 0xbf, 0xe7, 0x00, 0x4d, 0x0a,
	0x7c, 0xb6, 0xc6, 0x54, 0x0c, 0x96, 0x28, 0xa5, 0x23, 0xef, 0xd9, 0xec, 0x76, 0x76, 0xb7, 0x6e,
	0xca, 0xb3, 0x7a, 0x01, 0xcb, 0xe8, 0xa9, 0x0d, 0xda, 0xb3, 0xf9, 0xed, 0xec, 0x6e, 0xd5, 0x2c,
	0xa2, 0xa7, 0xa3, 0xf6, 0xd5, 0x6b, 0xd8, 0x2a, 0xe3, 0x49, 0x44, 0x97, 0xa7, 0x4f, 0xd2, 0x74,
	0x53, 0xd8, 0x51, 0xfb, 0xfd, 0xdf, 0x39, 0x5c, 0xe7, 0x25, 0x0d, 0xa1, 0xbc, 0x77, 0x9d, 0xaf,
	0xde, 0xc0, 0x33, 0x7a, 0x10, 0x27, 0x34, 0x1d, 0xb5, 0x06, 0x7b, 0x9a, 0xd6, 0x6d, 0x0b, 0xfc,
	0x81, 0x3d, 0x55, 0xaf, 0x00, 0x7e, 0x47, 0x8a, 0x53, 0x62, 0x9e, 0x12, 0xeb, 0x44, 0xd2, 0xf8,
	0x06, 0x36, 0x5c, 0x19, 0xa9, 0x4c, 0xd7, 0x9e, 0xe9, 0x31, 0x2d, 0x5e, 0x37, 0x30, 0xa1, 0xef,
	0xf4, 0x58, 0xbd, 0x85, 0x5d, 0xfe, 0x9f, 0x1e, 0x84, 0x8e, 0x5e, 0xfd, 0x21, 0xf6, 0x34, 0xb5,
	0xbb, 0x4e, 0xf8, 0x6b, 0xa1, 0x63, 0x9b, 0x1c, 0x94, 0x24, 0x34, 0x3a, 0x62, 0x57, 0x29, 0xb6,
	0x4d, 0xf0, 0x4b, 0x66, 0x17, 0xa1, 0xe8, 0x90, 0x6b, 0x62, 0x8b, 0xcb, 0x50, 0x66, 0xd5, 0x4b,
	0x58, 0x61, 0x0c, 0xb6, 0x45, 0x71, 0x66, 0xcb, 0x34, 0x5f, 0x8e, 0xef, 0x7b, 0x71, 0x1e, 0x45,
	0x09, 0x6b, 0x7c, 0xec, 0xc9, 0xb5, 0x01, 0x3b, 0xb6, 0x4a, 0x7d, 0x37, 0x85, 0x1d, 0xb1, 0x1b,
	0x23, 0xa5, 0x87, 0xa6, 0x40, 0x6c, 0x9d, 0x5d, 0x4e, 0x35, 0x46, 0xb4, 0x37, 0xb0, 0xcb, 0x2a,
	0x7f, 0x3a, 0x15, 0xe8, 0xff, 0x5d, 0xde, 0xc0, 0xc6, 0xd9, 0x18, 0x8a, 0xac, 0x2c, 0x13, 0x26,
	0x34, 0xca, 0x7a, 0x0e, 0x0b, 0x1c, 0x86, 0x56, 0xc9, 0x49, 0xe4, 0x15, 0x0e, 0xc3, 0x37, 0xf9,
	0xe9, 0xe3, 0xaf, 0x0f, 0x9d, 0x0a, 0xa7, 0xc8, 0x0f, 0xc2, 0xf6, 0x35, 0xc7, 0x20, 0x4e, 0xc2,
	0xba, 0xa1, 0x1e, 0x74, 0xec, 0x39, 0xb9, 0x77, 0x5e, 0x9c, 0xa8, 0x47, 0x5f, 0xf3, 0xa8, 0xb4,
	0xac, 0x3b, 0x5b, 0xe7, 0x03, 0xab, 0xcb, 0x81, 0xf1, 0x45, 0x02, 0xef, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x49, 0xf0, 0x33, 0xc6, 0x88, 0x02, 0x00, 0x00,
}
