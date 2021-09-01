// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafka.proto

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

type KafkaConn_SASLType int32

const (
	KafkaConn_NONE  KafkaConn_SASLType = 0
	KafkaConn_PLAIN KafkaConn_SASLType = 1
	KafkaConn_SCRAM KafkaConn_SASLType = 2
)

var KafkaConn_SASLType_name = map[int32]string{
	0: "NONE",
	1: "PLAIN",
	2: "SCRAM",
}

var KafkaConn_SASLType_value = map[string]int32{
	"NONE":  0,
	"PLAIN": 1,
	"SCRAM": 2,
}

func (x KafkaConn_SASLType) String() string {
	return proto.EnumName(KafkaConn_SASLType_name, int32(x))
}

func (KafkaConn_SASLType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{0, 0}
}

type KafkaConn struct {
	// @gotags: kong:"help='Kafka broker address (you may specify this flag multiple times',env=PLUMBER_RELAY_KAFKA_ADDRESS,default='localhost:9092',required"
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty" kong:"help='Kafka broker address (you may specify this flag multiple times',env=PLUMBER_RELAY_KAFKA_ADDRESS,default='localhost:9092',required"`
	// @gotags: kong:"help='Connect timeout',env=PLUMBER_RELAY_TIMEOUT_SECONDS,default=10"
	TimeoutSeconds int32 `protobuf:"varint,2,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty" kong:"help='Connect timeout',env=PLUMBER_RELAY_TIMEOUT_SECONDS,default=10"`
	// @gotags: kong:"help='Enable TLS usage',env=PLUMBER_RELAY_USE_TLS"
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Enable TLS usage',env=PLUMBER_RELAY_USE_TLS"`
	// @gotags: kong:"help='Allow insecure TLS usage',env=PLUMBER_RELAY_KAFKA_INSECURE_TLS"
	InsecureTls bool `protobuf:"varint,4,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty" kong:"help='Allow insecure TLS usage',env=PLUMBER_RELAY_KAFKA_INSECURE_TLS"`
	// @gotags: kong:"help='SASL authentication type (0 - none, 1 - plain, 2 - scram)',env=PLUMBER_RELAY_KAFKA_SASL_TYPE,default=2"
	SaslType KafkaConn_SASLType `protobuf:"varint,5,opt,name=sasl_type,json=saslType,proto3,enum=protos.backends.KafkaConn_SASLType" json:"sasl_type,omitempty" kong:"help='SASL authentication type (0 - none, 1 - plain, 2 - scram)',env=PLUMBER_RELAY_KAFKA_SASL_TYPE,default=2"`
	// @gotags: kong:"help='SASL Username',env=PLUMBER_RELAY_KAFKA_USERNAME"
	SaslUsername string `protobuf:"bytes,6,opt,name=sasl_username,json=saslUsername,proto3" json:"sasl_username,omitempty" kong:"help='SASL Username',env=PLUMBER_RELAY_KAFKA_USERNAME"`
	// Required if sasl_type is not NONE
	// @gotags: kong:"help='SASL Password. If omitted, you will be prompted for the password',env=PLUMBER_RELAY_KAFKA_PASSWORD"
	SaslPassword         string   `protobuf:"bytes,7,opt,name=sasl_password,json=saslPassword,proto3" json:"sasl_password,omitempty" kong:"help='SASL Password. If omitted, you will be prompted for the password',env=PLUMBER_RELAY_KAFKA_PASSWORD"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaConn) Reset()         { *m = KafkaConn{} }
func (m *KafkaConn) String() string { return proto.CompactTextString(m) }
func (*KafkaConn) ProtoMessage()    {}
func (*KafkaConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{0}
}

func (m *KafkaConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaConn.Unmarshal(m, b)
}
func (m *KafkaConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaConn.Marshal(b, m, deterministic)
}
func (m *KafkaConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaConn.Merge(m, src)
}
func (m *KafkaConn) XXX_Size() int {
	return xxx_messageInfo_KafkaConn.Size(m)
}
func (m *KafkaConn) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaConn.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaConn proto.InternalMessageInfo

func (m *KafkaConn) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *KafkaConn) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *KafkaConn) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *KafkaConn) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

func (m *KafkaConn) GetSaslType() KafkaConn_SASLType {
	if m != nil {
		return m.SaslType
	}
	return KafkaConn_NONE
}

func (m *KafkaConn) GetSaslUsername() string {
	if m != nil {
		return m.SaslUsername
	}
	return ""
}

func (m *KafkaConn) GetSaslPassword() string {
	if m != nil {
		return m.SaslPassword
	}
	return ""
}

type KafkaReadArgs struct {
	// @gotags: kong:"help='Topic(s) to read, write or get lag stats for',required"
	Topics []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty" kong:"help='Topic(s) to read, write or get lag stats for',required"`
	// @gotags: kong:"help='Specify what offset the consumer should read from (only works if '--use-consumer-group' is false)',default=0"
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty" kong:"help='Specify what offset the consumer should read from (only works if '--use-consumer-group' is false)',default=0"`
	// @gotags: kong:"help='Whether plumber should use a consumer group',default=true"
	UseConsumerGroup bool `protobuf:"varint,3,opt,name=use_consumer_group,json=useConsumerGroup,proto3" json:"use_consumer_group,omitempty" kong:"help='Whether plumber should use a consumer group',default=true"`
	// @gotags: kong:"help='Specify a specific group-id to use when reading from kafka',default=plumber"
	ConsumerGroupName string `protobuf:"bytes,4,opt,name=consumer_group_name,json=consumerGroupName,proto3" json:"consumer_group_name,omitempty" kong:"help='Specify a specific group-id to use when reading from kafka',default=plumber"`
	// @gotags: kong:"help='How long to wait for new data when reading batches of messages',default=1"
	MaxWaitSeconds int32 `protobuf:"varint,5,opt,name=max_wait_seconds,json=maxWaitSeconds,proto3" json:"max_wait_seconds,omitempty" kong:"help='How long to wait for new data when reading batches of messages',default=1"`
	// @gotags: kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"
	MinBytes int32 `protobuf:"varint,6,opt,name=min_bytes,json=minBytes,proto3" json:"min_bytes,omitempty" kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"`
	// @gotags: kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"
	MaxBytes int32 `protobuf:"varint,7,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty" kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"`
	// @gotags: kong:"help='How often to commit offsets to broker (0 = synchronous)',default=5"
	CommitIntervalSeconds int32 `protobuf:"varint,8,opt,name=commit_interval_seconds,json=commitIntervalSeconds,proto3" json:"commit_interval_seconds,omitempty" kong:"help='How often to commit offsets to broker (0 = synchronous)',default=5"`
	// @gotags: kong:"help='How long a coordinator will wait for member joins as part of a rebalance',default=0"
	RebalanceTimeoutSeconds int32 `protobuf:"varint,9,opt,name=rebalance_timeout_seconds,json=rebalanceTimeoutSeconds,proto3" json:"rebalance_timeout_seconds,omitempty" kong:"help='How long a coordinator will wait for member joins as part of a rebalance',default=0"`
	// @gotags: kong:"help='Internal library queue capacity (throughput optimization)',env=PLUMBER_RELAY_KAFKA_QUEUE_CAPACITY,default=1"
	QueueCapacity        int32    `protobuf:"varint,10,opt,name=queue_capacity,json=queueCapacity,proto3" json:"queue_capacity,omitempty" kong:"help='Internal library queue capacity (throughput optimization)',env=PLUMBER_RELAY_KAFKA_QUEUE_CAPACITY,default=1"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaReadArgs) Reset()         { *m = KafkaReadArgs{} }
func (m *KafkaReadArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaReadArgs) ProtoMessage()    {}
func (*KafkaReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{1}
}

func (m *KafkaReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaReadArgs.Unmarshal(m, b)
}
func (m *KafkaReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaReadArgs.Marshal(b, m, deterministic)
}
func (m *KafkaReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaReadArgs.Merge(m, src)
}
func (m *KafkaReadArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaReadArgs.Size(m)
}
func (m *KafkaReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaReadArgs proto.InternalMessageInfo

func (m *KafkaReadArgs) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *KafkaReadArgs) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *KafkaReadArgs) GetUseConsumerGroup() bool {
	if m != nil {
		return m.UseConsumerGroup
	}
	return false
}

func (m *KafkaReadArgs) GetConsumerGroupName() string {
	if m != nil {
		return m.ConsumerGroupName
	}
	return ""
}

func (m *KafkaReadArgs) GetMaxWaitSeconds() int32 {
	if m != nil {
		return m.MaxWaitSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetMinBytes() int32 {
	if m != nil {
		return m.MinBytes
	}
	return 0
}

func (m *KafkaReadArgs) GetMaxBytes() int32 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *KafkaReadArgs) GetCommitIntervalSeconds() int32 {
	if m != nil {
		return m.CommitIntervalSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetRebalanceTimeoutSeconds() int32 {
	if m != nil {
		return m.RebalanceTimeoutSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetQueueCapacity() int32 {
	if m != nil {
		return m.QueueCapacity
	}
	return 0
}

type KafkaWriteArgs struct {
	// @gotags: kong:"help='Key to write to kafka (optional)'"
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" kong:"help='Key to write to kafka (optional)'"`
	// @gotags: kong:"help='Add one or more headers (optional; repeat flags to specify multiple)'"
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" kong:"help='Add one or more headers (optional; repeat flags to specify multiple)'"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KafkaWriteArgs) Reset()         { *m = KafkaWriteArgs{} }
func (m *KafkaWriteArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaWriteArgs) ProtoMessage()    {}
func (*KafkaWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{2}
}

func (m *KafkaWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaWriteArgs.Unmarshal(m, b)
}
func (m *KafkaWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaWriteArgs.Marshal(b, m, deterministic)
}
func (m *KafkaWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaWriteArgs.Merge(m, src)
}
func (m *KafkaWriteArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaWriteArgs.Size(m)
}
func (m *KafkaWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaWriteArgs proto.InternalMessageInfo

func (m *KafkaWriteArgs) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KafkaWriteArgs) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type KafkaRelayArgs struct {
	// @gotags: kong:"help='Topic(s) to read, write or get lag stats for',env=PLUMBER_RELAY_KAFKA_TOPIC,required"
	Topics []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty" kong:"help='Topic(s) to read, write or get lag stats for',env=PLUMBER_RELAY_KAFKA_TOPIC,required"`
	// @gotags: kong:"help='Specify what offset the consumer should read from (only works if '--use-consumer-group' is false)',env=PLUMBER_RELAY_KAFKA_READ_OFFSET,default=0"
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty" kong:"help='Specify what offset the consumer should read from (only works if '--use-consumer-group' is false)',env=PLUMBER_RELAY_KAFKA_READ_OFFSET,default=0"`
	// @gotags: kong:"help='Whether plumber should use a consumer group',env=PLUMBER_RELAY_KAFKA_USE_CONSUMER_GROUP,default=true"
	UseConsumerGroup bool `protobuf:"varint,3,opt,name=use_consumer_group,json=useConsumerGroup,proto3" json:"use_consumer_group,omitempty" kong:"help='Whether plumber should use a consumer group',env=PLUMBER_RELAY_KAFKA_USE_CONSUMER_GROUP,default=true"`
	// @gotags: kong:"help='Specify a specific group-id to use when reading from kafka',env=PLUMBER_RELAY_KAFKA_GROUP_ID,default=plumber"
	ConsumerGroupName string `protobuf:"bytes,4,opt,name=consumer_group_name,json=consumerGroupName,proto3" json:"consumer_group_name,omitempty" kong:"help='Specify a specific group-id to use when reading from kafka',env=PLUMBER_RELAY_KAFKA_GROUP_ID,default=plumber"`
	// @gotags: kong:"help='How long to wait for new data when reading batches of messages',env=PLUMBER_RELAY_KAFKA_MAX_WAIT,default=5"
	MaxWaitSeconds int32 `protobuf:"varint,5,opt,name=max_wait_seconds,json=maxWaitSeconds,proto3" json:"max_wait_seconds,omitempty" kong:"help='How long to wait for new data when reading batches of messages',env=PLUMBER_RELAY_KAFKA_MAX_WAIT,default=5"`
	// @gotags: kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MIN_BYTES,default=1048576"
	MinBytes int32 `protobuf:"varint,6,opt,name=min_bytes,json=minBytes,proto3" json:"min_bytes,omitempty" kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MIN_BYTES,default=1048576"`
	// @gotags: kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MAX_BYTES,default=1048576"
	MaxBytes int32 `protobuf:"varint,7,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty" kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MAX_BYTES,default=1048576"`
	// @gotags: kong:"help='How often to commit offsets to broker (0 = synchronous)',env=PLUMBER_RELAY_KAFKA_COMMIT_INTERVAL,default=5"
	CommitIntervalSeconds int32 `protobuf:"varint,8,opt,name=commit_interval_seconds,json=commitIntervalSeconds,proto3" json:"commit_interval_seconds,omitempty" kong:"help='How often to commit offsets to broker (0 = synchronous)',env=PLUMBER_RELAY_KAFKA_COMMIT_INTERVAL,default=5"`
	// @gotags: kong:"help='How long a coordinator will wait for member joins as part of a rebalance',env=PLUMBER_RELAY_KAFKA_REBALANCE_TIMEOUT,default=5"
	RebalanceTimeoutSeconds int32 `protobuf:"varint,9,opt,name=rebalance_timeout_seconds,json=rebalanceTimeoutSeconds,proto3" json:"rebalance_timeout_seconds,omitempty" kong:"help='How long a coordinator will wait for member joins as part of a rebalance',env=PLUMBER_RELAY_KAFKA_REBALANCE_TIMEOUT,default=5"`
	// @gotags: kong:"help='Internal library queue capacity (throughput optimization)',env=PLUMBER_RELAY_KAFKA_QUEUE_CAPACITY,default=1000"
	QueueCapacity        int32    `protobuf:"varint,10,opt,name=queue_capacity,json=queueCapacity,proto3" json:"queue_capacity,omitempty" kong:"help='Internal library queue capacity (throughput optimization)',env=PLUMBER_RELAY_KAFKA_QUEUE_CAPACITY,default=1000"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaRelayArgs) Reset()         { *m = KafkaRelayArgs{} }
func (m *KafkaRelayArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaRelayArgs) ProtoMessage()    {}
func (*KafkaRelayArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{3}
}

func (m *KafkaRelayArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaRelayArgs.Unmarshal(m, b)
}
func (m *KafkaRelayArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaRelayArgs.Marshal(b, m, deterministic)
}
func (m *KafkaRelayArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaRelayArgs.Merge(m, src)
}
func (m *KafkaRelayArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaRelayArgs.Size(m)
}
func (m *KafkaRelayArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaRelayArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaRelayArgs proto.InternalMessageInfo

func (m *KafkaRelayArgs) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *KafkaRelayArgs) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *KafkaRelayArgs) GetUseConsumerGroup() bool {
	if m != nil {
		return m.UseConsumerGroup
	}
	return false
}

func (m *KafkaRelayArgs) GetConsumerGroupName() string {
	if m != nil {
		return m.ConsumerGroupName
	}
	return ""
}

func (m *KafkaRelayArgs) GetMaxWaitSeconds() int32 {
	if m != nil {
		return m.MaxWaitSeconds
	}
	return 0
}

func (m *KafkaRelayArgs) GetMinBytes() int32 {
	if m != nil {
		return m.MinBytes
	}
	return 0
}

func (m *KafkaRelayArgs) GetMaxBytes() int32 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *KafkaRelayArgs) GetCommitIntervalSeconds() int32 {
	if m != nil {
		return m.CommitIntervalSeconds
	}
	return 0
}

func (m *KafkaRelayArgs) GetRebalanceTimeoutSeconds() int32 {
	if m != nil {
		return m.RebalanceTimeoutSeconds
	}
	return 0
}

func (m *KafkaRelayArgs) GetQueueCapacity() int32 {
	if m != nil {
		return m.QueueCapacity
	}
	return 0
}

type KafkaLagArgs struct {
	// @gotags: kong:"help='What consumer group to look up lag for (leave blank if all)'"
	GroupId              string   `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" kong:"help='What consumer group to look up lag for (leave blank if all)'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaLagArgs) Reset()         { *m = KafkaLagArgs{} }
func (m *KafkaLagArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaLagArgs) ProtoMessage()    {}
func (*KafkaLagArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_68928ed13de9fb92, []int{4}
}

func (m *KafkaLagArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaLagArgs.Unmarshal(m, b)
}
func (m *KafkaLagArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaLagArgs.Marshal(b, m, deterministic)
}
func (m *KafkaLagArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaLagArgs.Merge(m, src)
}
func (m *KafkaLagArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaLagArgs.Size(m)
}
func (m *KafkaLagArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaLagArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaLagArgs proto.InternalMessageInfo

func (m *KafkaLagArgs) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.backends.KafkaConn_SASLType", KafkaConn_SASLType_name, KafkaConn_SASLType_value)
	proto.RegisterType((*KafkaConn)(nil), "protos.backends.KafkaConn")
	proto.RegisterType((*KafkaReadArgs)(nil), "protos.backends.KafkaReadArgs")
	proto.RegisterType((*KafkaWriteArgs)(nil), "protos.backends.KafkaWriteArgs")
	proto.RegisterMapType((map[string]string)(nil), "protos.backends.KafkaWriteArgs.HeadersEntry")
	proto.RegisterType((*KafkaRelayArgs)(nil), "protos.backends.KafkaRelayArgs")
	proto.RegisterType((*KafkaLagArgs)(nil), "protos.backends.KafkaLagArgs")
}

func init() { proto.RegisterFile("kafka.proto", fileDescriptor_68928ed13de9fb92) }

var fileDescriptor_68928ed13de9fb92 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x52, 0xdb, 0x6e, 0xd3, 0x4a,
	0x14, 0x3d, 0x4e, 0x9a, 0x8b, 0x77, 0xd2, 0x34, 0xc7, 0xe7, 0x40, 0x5d, 0x78, 0x20, 0xa4, 0x42,
	0x04, 0x54, 0x1c, 0xa9, 0x48, 0x08, 0x55, 0x42, 0x90, 0x46, 0x05, 0x2a, 0x4a, 0x5a, 0x39, 0x45,
	0x95, 0x78, 0xb1, 0xc6, 0xf6, 0x6e, 0x32, 0x8a, 0xed, 0x31, 0x73, 0x69, 0x93, 0x0f, 0xa2, 0x1f,
	0xc2, 0x97, 0x21, 0x8f, 0x9d, 0xf4, 0xa2, 0x7e, 0x00, 0x0f, 0x3c, 0xd9, 0x7b, 0xad, 0xb5, 0x67,
	0xf6, 0xec, 0xb5, 0xa0, 0x31, 0x23, 0xe7, 0x33, 0xe2, 0xa4, 0x9c, 0x49, 0x66, 0x6d, 0xe8, 0x8f,
	0x70, 0x7c, 0x12, 0xcc, 0x30, 0x09, 0x45, 0xf7, 0x57, 0x09, 0xcc, 0x2f, 0x99, 0x60, 0xc8, 0x92,
	0xc4, 0xb2, 0xa1, 0x46, 0xc2, 0x90, 0xa3, 0x10, 0xb6, 0xd1, 0x29, 0xf7, 0x4c, 0x77, 0x59, 0x5a,
	0xcf, 0x61, 0x43, 0xd2, 0x18, 0x99, 0x92, 0x9e, 0xc0, 0x80, 0x25, 0xa1, 0xb0, 0x4b, 0x1d, 0xa3,
	0x57, 0x71, 0x5b, 0x05, 0x3c, 0xce, 0x51, 0x6b, 0x13, 0x6a, 0x4a, 0xa0, 0x27, 0x23, 0x61, 0x97,
	0x3b, 0x46, 0xaf, 0xee, 0x56, 0x95, 0xc0, 0xd3, 0x48, 0x58, 0x4f, 0xa1, 0x49, 0x13, 0x81, 0x81,
	0xe2, 0x39, 0xbb, 0xa6, 0xd9, 0xc6, 0x12, 0xcb, 0x24, 0x1f, 0xc0, 0x14, 0x44, 0x44, 0x9e, 0x5c,
	0xa4, 0x68, 0x57, 0x3a, 0x46, 0xaf, 0xb5, 0xbb, 0xed, 0xdc, 0x99, 0xd8, 0x59, 0x4d, 0xeb, 0x8c,
	0x07, 0xe3, 0xa3, 0xd3, 0x45, 0x8a, 0x6e, 0x3d, 0xeb, 0xca, 0xfe, 0xac, 0x6d, 0x58, 0xd7, 0x27,
	0x28, 0x81, 0x3c, 0x21, 0x31, 0xda, 0xd5, 0x8e, 0xd1, 0x33, 0xdd, 0x66, 0x06, 0x7e, 0x2b, 0xb0,
	0x95, 0x28, 0x25, 0x42, 0x5c, 0x32, 0x1e, 0xda, 0xb5, 0x6b, 0xd1, 0x49, 0x81, 0x75, 0x5f, 0x42,
	0x7d, 0x79, 0xbe, 0x55, 0x87, 0xb5, 0xd1, 0xf1, 0xe8, 0xa0, 0xfd, 0x8f, 0x65, 0x42, 0xe5, 0xe4,
	0x68, 0x70, 0x38, 0x6a, 0x1b, 0xd9, 0xef, 0x78, 0xe8, 0x0e, 0xbe, 0xb6, 0x4b, 0xdd, 0x9f, 0x65,
	0x58, 0xd7, 0x63, 0xb9, 0x48, 0xc2, 0x01, 0x9f, 0x08, 0xeb, 0x21, 0x54, 0x25, 0x4b, 0x69, 0xb0,
	0xdc, 0x63, 0x51, 0x59, 0x4f, 0xa0, 0xc1, 0x91, 0x84, 0x1e, 0x3b, 0x3f, 0x17, 0x28, 0xf5, 0x0a,
	0xcb, 0x2e, 0x64, 0xd0, 0xb1, 0x46, 0xac, 0x1d, 0xb0, 0xb2, 0xf5, 0x05, 0x2c, 0x11, 0x2a, 0x46,
	0xee, 0x4d, 0x38, 0x53, 0x69, 0xb1, 0xc9, 0xb6, 0x12, 0x38, 0x2c, 0x88, 0x4f, 0x19, 0x6e, 0x39,
	0xf0, 0xdf, 0x6d, 0xa5, 0xa7, 0x1f, 0xbd, 0xa6, 0xdf, 0xf3, 0x6f, 0x70, 0x53, 0x3b, 0xca, 0x5e,
	0xde, 0x83, 0x76, 0x4c, 0xe6, 0xde, 0x25, 0xa1, 0xd7, 0x36, 0x56, 0x72, 0x1b, 0x63, 0x32, 0x3f,
	0x23, 0x74, 0x65, 0xe3, 0x63, 0x30, 0x63, 0x9a, 0x78, 0xfe, 0x42, 0xa2, 0xd0, 0x4b, 0xac, 0xb8,
	0xf5, 0x98, 0x26, 0xfb, 0x59, 0xad, 0x49, 0x32, 0x2f, 0xc8, 0x5a, 0x41, 0x92, 0x79, 0x4e, 0xbe,
	0x81, 0xcd, 0x80, 0xc5, 0x31, 0x95, 0x1e, 0x4d, 0x24, 0xf2, 0x0b, 0x12, 0xad, 0xae, 0xaa, 0x6b,
	0xe9, 0x83, 0x9c, 0x3e, 0x2c, 0xd8, 0xe5, 0x8d, 0x7b, 0xb0, 0xc5, 0xd1, 0x27, 0x11, 0x49, 0x02,
	0xf4, 0xee, 0x66, 0xcd, 0xd4, 0x9d, 0x9b, 0x2b, 0xc1, 0xe9, 0xed, 0xd0, 0x3d, 0x83, 0xd6, 0x0f,
	0x85, 0x0a, 0xbd, 0x80, 0xa4, 0x24, 0xa0, 0x72, 0x61, 0x83, 0x6e, 0x58, 0xd7, 0xe8, 0xb0, 0x00,
	0xbb, 0x57, 0x06, 0xb4, 0xb4, 0x4f, 0x67, 0x9c, 0x4a, 0xd4, 0x46, 0xb5, 0xa1, 0x3c, 0xc3, 0x85,
	0x6d, 0xe8, 0x8d, 0x65, 0xbf, 0xd6, 0x47, 0xa8, 0x4d, 0x91, 0x84, 0xc8, 0xb3, 0x84, 0x97, 0x7b,
	0x8d, 0xdd, 0x9d, 0xfb, 0x23, 0xb8, 0x3a, 0xc3, 0xf9, 0x9c, 0xcb, 0x0f, 0x12, 0xc9, 0x17, 0xee,
	0xb2, 0xf9, 0xd1, 0x1e, 0x34, 0x6f, 0x12, 0xf7, 0xdc, 0xf4, 0x3f, 0x54, 0x2e, 0x48, 0xa4, 0x50,
	0xc7, 0xc0, 0x74, 0xf3, 0x62, 0xaf, 0xf4, 0xd6, 0xe8, 0x5e, 0x95, 0x8b, 0x41, 0x5d, 0x8c, 0xc8,
	0xe2, 0x6f, 0xa2, 0xfe, 0xd8, 0x44, 0xbd, 0x80, 0xa6, 0xf6, 0xe9, 0x88, 0x4c, 0xb4, 0x4b, 0x5b,
	0x50, 0xcf, 0xb7, 0x46, 0xc3, 0xc2, 0xe9, 0x9a, 0xae, 0x0f, 0xc3, 0xfd, 0xf7, 0xdf, 0xdf, 0x4d,
	0xa8, 0x9c, 0x2a, 0xdf, 0x09, 0x58, 0xdc, 0xf7, 0x89, 0x0c, 0xa6, 0x01, 0xe3, 0x69, 0x3f, 0x8d,
	0x54, 0xec, 0x23, 0x7f, 0x25, 0x82, 0x29, 0xc6, 0x44, 0xf4, 0x7d, 0x45, 0xa3, 0xb0, 0x3f, 0x61,
	0xfd, 0x3c, 0x75, 0xfd, 0x65, 0xea, 0xfc, 0xaa, 0x06, 0x5e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0xae, 0x64, 0xb5, 0xd1, 0x05, 0x00, 0x00,
}