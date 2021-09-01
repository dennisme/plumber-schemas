// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pulsar.proto

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

type SubscriptionType int32

const (
	SubscriptionType_SHARED    SubscriptionType = 0
	SubscriptionType_EXCLUSIVE SubscriptionType = 1
	SubscriptionType_FAILOVER  SubscriptionType = 2
	SubscriptionType_KEYSHARED SubscriptionType = 3
)

var SubscriptionType_name = map[int32]string{
	0: "SHARED",
	1: "EXCLUSIVE",
	2: "FAILOVER",
	3: "KEYSHARED",
}

var SubscriptionType_value = map[string]int32{
	"SHARED":    0,
	"EXCLUSIVE": 1,
	"FAILOVER":  2,
	"KEYSHARED": 3,
}

func (x SubscriptionType) String() string {
	return proto.EnumName(SubscriptionType_name, int32(x))
}

func (SubscriptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ffc7a491d9398c1a, []int{0}
}

type PulsarConn struct {
	// @gotags: kong:"help='Full DSN to connect to Pulsar',default='pulsar://localhost:6650',required"
	Dsn string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty" kong:"help='Full DSN to connect to Pulsar',default='pulsar://localhost:6650',required"`
	// @gotags: kong:"help='Connection timeout',default=10"
	ConnectTimeoutSeconds uint32 `protobuf:"varint,2,opt,name=connect_timeout_seconds,json=connectTimeoutSeconds,proto3" json:"connect_timeout_seconds,omitempty" kong:"help='Connection timeout',default=10"`
	// @gotags: kong:"help='Whether to verify server certificate'"
	InsecureTls bool `protobuf:"varint,3,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty" kong:"help='Whether to verify server certificate'"`
	// @gotags: kong:"help='TLS client certificate file',type=existingfile"
	TlsClientCert []byte `protobuf:"bytes,4,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty" kong:"help='TLS client certificate file',type=existingfile"`
	// @gotags: kong:"help='TLS client key file',type=existingfile"
	TlsClientKey         []byte   `protobuf:"bytes,5,opt,name=tls_client_key,json=tlsClientKey,proto3" json:"tls_client_key,omitempty" kong:"help='TLS client key file',type=existingfile"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PulsarConn) Reset()         { *m = PulsarConn{} }
func (m *PulsarConn) String() string { return proto.CompactTextString(m) }
func (*PulsarConn) ProtoMessage()    {}
func (*PulsarConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc7a491d9398c1a, []int{0}
}

func (m *PulsarConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PulsarConn.Unmarshal(m, b)
}
func (m *PulsarConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PulsarConn.Marshal(b, m, deterministic)
}
func (m *PulsarConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PulsarConn.Merge(m, src)
}
func (m *PulsarConn) XXX_Size() int {
	return xxx_messageInfo_PulsarConn.Size(m)
}
func (m *PulsarConn) XXX_DiscardUnknown() {
	xxx_messageInfo_PulsarConn.DiscardUnknown(m)
}

var xxx_messageInfo_PulsarConn proto.InternalMessageInfo

func (m *PulsarConn) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

func (m *PulsarConn) GetConnectTimeoutSeconds() uint32 {
	if m != nil {
		return m.ConnectTimeoutSeconds
	}
	return 0
}

func (m *PulsarConn) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

func (m *PulsarConn) GetTlsClientCert() []byte {
	if m != nil {
		return m.TlsClientCert
	}
	return nil
}

func (m *PulsarConn) GetTlsClientKey() []byte {
	if m != nil {
		return m.TlsClientKey
	}
	return nil
}

type PulsarReadArgs struct {
	// @gotags: kong:"help='Topic to read messages from',required"
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic to read messages from',required"`
	// @gotags: kong:"help='Subscription name',required"
	SubscriptionName string `protobuf:"bytes,2,opt,name=subscription_name,json=subscriptionName,proto3" json:"subscription_name,omitempty" kong:"help='Subscription name',required"`
	// @gotags: kong:"help='Subscription type (0: shared, 1: exclusive, 2: failover, 3: keyshared)',required"
	SubscriptionType     SubscriptionType `protobuf:"varint,3,opt,name=subscription_type,json=subscriptionType,proto3,enum=protos.backends.SubscriptionType" json:"subscription_type,omitempty" kong:"help='Subscription type (0: shared, 1: exclusive, 2: failover, 3: keyshared)',required"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PulsarReadArgs) Reset()         { *m = PulsarReadArgs{} }
func (m *PulsarReadArgs) String() string { return proto.CompactTextString(m) }
func (*PulsarReadArgs) ProtoMessage()    {}
func (*PulsarReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc7a491d9398c1a, []int{1}
}

func (m *PulsarReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PulsarReadArgs.Unmarshal(m, b)
}
func (m *PulsarReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PulsarReadArgs.Marshal(b, m, deterministic)
}
func (m *PulsarReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PulsarReadArgs.Merge(m, src)
}
func (m *PulsarReadArgs) XXX_Size() int {
	return xxx_messageInfo_PulsarReadArgs.Size(m)
}
func (m *PulsarReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_PulsarReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_PulsarReadArgs proto.InternalMessageInfo

func (m *PulsarReadArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PulsarReadArgs) GetSubscriptionName() string {
	if m != nil {
		return m.SubscriptionName
	}
	return ""
}

func (m *PulsarReadArgs) GetSubscriptionType() SubscriptionType {
	if m != nil {
		return m.SubscriptionType
	}
	return SubscriptionType_SHARED
}

type PulsarWriteArgs struct {
	// @gotags: kong:"help='topic to write messages to'"
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='topic to write messages to'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PulsarWriteArgs) Reset()         { *m = PulsarWriteArgs{} }
func (m *PulsarWriteArgs) String() string { return proto.CompactTextString(m) }
func (*PulsarWriteArgs) ProtoMessage()    {}
func (*PulsarWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffc7a491d9398c1a, []int{2}
}

func (m *PulsarWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PulsarWriteArgs.Unmarshal(m, b)
}
func (m *PulsarWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PulsarWriteArgs.Marshal(b, m, deterministic)
}
func (m *PulsarWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PulsarWriteArgs.Merge(m, src)
}
func (m *PulsarWriteArgs) XXX_Size() int {
	return xxx_messageInfo_PulsarWriteArgs.Size(m)
}
func (m *PulsarWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_PulsarWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_PulsarWriteArgs proto.InternalMessageInfo

func (m *PulsarWriteArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.backends.SubscriptionType", SubscriptionType_name, SubscriptionType_value)
	proto.RegisterType((*PulsarConn)(nil), "protos.backends.PulsarConn")
	proto.RegisterType((*PulsarReadArgs)(nil), "protos.backends.PulsarReadArgs")
	proto.RegisterType((*PulsarWriteArgs)(nil), "protos.backends.PulsarWriteArgs")
}

func init() { proto.RegisterFile("pulsar.proto", fileDescriptor_ffc7a491d9398c1a) }

var fileDescriptor_ffc7a491d9398c1a = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xcd, 0xd6, 0x5d, 0xb6, 0xc7, 0xb4, 0x8d, 0x83, 0x62, 0x2e, 0xbb, 0x45, 0xb4, 0x28,
	0x26, 0xa0, 0xe0, 0x9d, 0x48, 0xad, 0x11, 0xd7, 0x5d, 0x56, 0x49, 0xeb, 0xfa, 0xe7, 0x26, 0x24,
	0x93, 0x43, 0x3b, 0x6c, 0x32, 0x33, 0xcc, 0x39, 0xb9, 0xe8, 0xf3, 0xf8, 0x30, 0xbe, 0x96, 0x34,
	0xe9, 0x4a, 0x2d, 0x78, 0x35, 0x33, 0xdf, 0xf7, 0xe3, 0x70, 0x7e, 0x03, 0xbe, 0x6d, 0x2a, 0xca,
	0x5d, 0x64, 0x9d, 0x61, 0x23, 0x46, 0xed, 0x41, 0x51, 0x91, 0xcb, 0x1b, 0xd4, 0x25, 0x4d, 0x7e,
	0x7b, 0x00, 0x5f, 0x5a, 0x62, 0x6e, 0xb4, 0x16, 0x01, 0xf4, 0x4a, 0xd2, 0xa1, 0x37, 0xf6, 0xa6,
	0xfd, 0x74, 0x7b, 0x15, 0xaf, 0xe1, 0x91, 0x34, 0x5a, 0xa3, 0xe4, 0x8c, 0x55, 0x8d, 0xa6, 0xe1,
	0x8c, 0x50, 0x1a, 0x5d, 0x52, 0x78, 0x34, 0xf6, 0xa6, 0x83, 0xf4, 0xe1, 0xae, 0x5e, 0x76, 0xed,
	0xa2, 0x2b, 0xc5, 0x19, 0xf8, 0x4a, 0x13, 0xca, 0xc6, 0x61, 0xc6, 0x15, 0x85, 0xbd, 0xb1, 0x37,
	0x3d, 0x4d, 0xef, 0xdd, 0x66, 0xcb, 0x8a, 0xc4, 0x13, 0x18, 0x71, 0x45, 0x99, 0xac, 0x14, 0x6a,
	0xce, 0x24, 0x3a, 0x0e, 0xef, 0x8e, 0xbd, 0xa9, 0x9f, 0x0e, 0xb8, 0xa2, 0x79, 0x9b, 0xce, 0xd1,
	0xb1, 0x78, 0x0c, 0xc3, 0x3d, 0xee, 0x06, 0x37, 0xe1, 0x71, 0x8b, 0xf9, 0x7f, 0xb1, 0x0b, 0xdc,
	0x4c, 0x7e, 0x79, 0x30, 0xec, 0x4c, 0x52, 0xcc, 0xcb, 0x99, 0x5b, 0x91, 0x78, 0x00, 0xc7, 0x6c,
	0xac, 0x92, 0x3b, 0x9f, 0xee, 0x21, 0x9e, 0xc3, 0x7d, 0x6a, 0x0a, 0x92, 0x4e, 0x59, 0x56, 0x46,
	0x67, 0x3a, 0xaf, 0xb1, 0x75, 0xe9, 0xa7, 0xc1, 0x7e, 0x71, 0x95, 0xd7, 0x28, 0xae, 0x0e, 0x60,
	0xde, 0x58, 0x6c, 0x5d, 0x86, 0x2f, 0xcf, 0xa2, 0x83, 0xcf, 0x8c, 0x16, 0x7b, 0xe4, 0x72, 0x63,
	0xf1, 0xdf, 0x79, 0xdb, 0x64, 0xf2, 0x14, 0x46, 0xdd, 0x92, 0xdf, 0x9c, 0x62, 0xfc, 0xff, 0x96,
	0xcf, 0x3e, 0x41, 0x70, 0x38, 0x4e, 0x00, 0x9c, 0x2c, 0x3e, 0xce, 0xd2, 0xe4, 0x7d, 0x70, 0x47,
	0x0c, 0xa0, 0x9f, 0x7c, 0x9f, 0x5f, 0x7e, 0x5d, 0x9c, 0x5f, 0x27, 0x81, 0x27, 0x7c, 0x38, 0xfd,
	0x30, 0x3b, 0xbf, 0xfc, 0x7c, 0x9d, 0xa4, 0xc1, 0xd1, 0xb6, 0xbc, 0x48, 0x7e, 0xec, 0xd8, 0xde,
	0xbb, 0xb7, 0x3f, 0xdf, 0xac, 0x14, 0xaf, 0x9b, 0x22, 0x92, 0xa6, 0x8e, 0x8b, 0x9c, 0xe5, 0x5a,
	0x1a, 0x67, 0x63, 0x5b, 0x35, 0x75, 0x81, 0xee, 0x05, 0xc9, 0x35, 0xd6, 0x39, 0xc5, 0x45, 0xa3,
	0xaa, 0x32, 0x5e, 0x99, 0xb8, 0x13, 0x8b, 0x6f, 0xc5, 0x8a, 0x93, 0x36, 0x78, 0xf5, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x92, 0x3c, 0xab, 0x6b, 0x4d, 0x02, 0x00, 0x00,
}