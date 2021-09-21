// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_common_backends.proto

package common

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

type BackendType int32

const (
	BackendType_BACKEND_TYPE_UNSET             BackendType = 0
	BackendType_BACKEND_TYPE_KAFKA             BackendType = 1
	BackendType_BACKEND_TYPE_RABBIT            BackendType = 2
	BackendType_BACKEND_TYPE_RABBIT_STREAMS    BackendType = 3
	BackendType_BACKEND_TYPE_NSQ               BackendType = 4
	BackendType_BACKEND_TYPE_NATS              BackendType = 5
	BackendType_BACKEND_TYPE_NATS_STREAMING    BackendType = 6
	BackendType_BACKEND_TYPE_GCP_PUBSUB        BackendType = 7
	BackendType_BACKEND_TYPE_AZURE_SERVICE_BUS BackendType = 8
	BackendType_BACKEND_TYPE_AZURE_EVENT_HUB   BackendType = 9
	BackendType_BACKEND_TYPE_AWS_SQS           BackendType = 10
	BackendType_BACKEND_TYPE_AWS_SNS           BackendType = 11
	BackendType_BACKEND_TYPE_REDIS_PUBSUB      BackendType = 12
	BackendType_BACKEND_TYPE_REDIS_STREAMS     BackendType = 13
	BackendType_BACKEND_TYPE_ACTIVEMQ          BackendType = 14
	BackendType_BACKEND_TYPE_PULSAR            BackendType = 15
	BackendType_BACKEND_TYPE_MQTT              BackendType = 16
	BackendType_BACKEND_TYPE_POSTGRES_CDC      BackendType = 17
	BackendType_BACKEND_TYPE_MONGODB_CDC       BackendType = 18
	BackendType_BACKEND_TYPE_KUBE_MQ           BackendType = 19
)

var BackendType_name = map[int32]string{
	0:  "BACKEND_TYPE_UNSET",
	1:  "BACKEND_TYPE_KAFKA",
	2:  "BACKEND_TYPE_RABBIT",
	3:  "BACKEND_TYPE_RABBIT_STREAMS",
	4:  "BACKEND_TYPE_NSQ",
	5:  "BACKEND_TYPE_NATS",
	6:  "BACKEND_TYPE_NATS_STREAMING",
	7:  "BACKEND_TYPE_GCP_PUBSUB",
	8:  "BACKEND_TYPE_AZURE_SERVICE_BUS",
	9:  "BACKEND_TYPE_AZURE_EVENT_HUB",
	10: "BACKEND_TYPE_AWS_SQS",
	11: "BACKEND_TYPE_AWS_SNS",
	12: "BACKEND_TYPE_REDIS_PUBSUB",
	13: "BACKEND_TYPE_REDIS_STREAMS",
	14: "BACKEND_TYPE_ACTIVEMQ",
	15: "BACKEND_TYPE_PULSAR",
	16: "BACKEND_TYPE_MQTT",
	17: "BACKEND_TYPE_POSTGRES_CDC",
	18: "BACKEND_TYPE_MONGODB_CDC",
	19: "BACKEND_TYPE_KUBE_MQ",
}

var BackendType_value = map[string]int32{
	"BACKEND_TYPE_UNSET":             0,
	"BACKEND_TYPE_KAFKA":             1,
	"BACKEND_TYPE_RABBIT":            2,
	"BACKEND_TYPE_RABBIT_STREAMS":    3,
	"BACKEND_TYPE_NSQ":               4,
	"BACKEND_TYPE_NATS":              5,
	"BACKEND_TYPE_NATS_STREAMING":    6,
	"BACKEND_TYPE_GCP_PUBSUB":        7,
	"BACKEND_TYPE_AZURE_SERVICE_BUS": 8,
	"BACKEND_TYPE_AZURE_EVENT_HUB":   9,
	"BACKEND_TYPE_AWS_SQS":           10,
	"BACKEND_TYPE_AWS_SNS":           11,
	"BACKEND_TYPE_REDIS_PUBSUB":      12,
	"BACKEND_TYPE_REDIS_STREAMS":     13,
	"BACKEND_TYPE_ACTIVEMQ":          14,
	"BACKEND_TYPE_PULSAR":            15,
	"BACKEND_TYPE_MQTT":              16,
	"BACKEND_TYPE_POSTGRES_CDC":      17,
	"BACKEND_TYPE_MONGODB_CDC":       18,
	"BACKEND_TYPE_KUBE_MQ":           19,
}

func (x BackendType) String() string {
	return proto.EnumName(BackendType_name, int32(x))
}

func (BackendType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7390f38ee25065b, []int{0}
}

func init() {
	proto.RegisterEnum("protos.common.BackendType", BackendType_name, BackendType_value)
}

func init() { proto.RegisterFile("ps_common_backends.proto", fileDescriptor_b7390f38ee25065b) }

var fileDescriptor_b7390f38ee25065b = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xf9, 0x53, 0x0a, 0x4c, 0x29, 0x4c, 0xa7, 0x2d, 0x4d, 0x69, 0x29, 0x88, 0x23, 0x12,
	0xf5, 0x81, 0x23, 0xe2, 0xb0, 0xeb, 0x2c, 0xc1, 0x32, 0x71, 0x6c, 0xcf, 0x6e, 0x10, 0xbd, 0xac,
	0x62, 0xd7, 0x6a, 0x2a, 0xea, 0xda, 0x8a, 0x93, 0x03, 0x1f, 0x8e, 0xef, 0x86, 0x12, 0x93, 0xc3,
	0xca, 0x3e, 0xad, 0xf6, 0xfd, 0xde, 0x3e, 0xcd, 0x3e, 0x0d, 0x0c, 0xea, 0xc6, 0xe6, 0x55, 0x59,
	0x56, 0xf7, 0x36, 0x9b, 0xe5, 0xbf, 0x8b, 0xfb, 0xeb, 0xe6, 0xb2, 0x5e, 0x54, 0xcb, 0x8a, 0xf6,
	0x37, 0x47, 0x73, 0xd9, 0xd2, 0x8f, 0x7f, 0x77, 0x60, 0x4f, 0xb6, 0x0e, 0xfd, 0xa7, 0x2e, 0xe8,
	0x35, 0x90, 0x14, 0x7e, 0xa8, 0xa2, 0xa1, 0xd5, 0xbf, 0x62, 0x65, 0x4d, 0xc4, 0x4a, 0xe3, 0x83,
	0x8e, 0x1e, 0x8a, 0x6f, 0xa1, 0xc0, 0x87, 0x74, 0x02, 0x87, 0x8e, 0x9e, 0x0a, 0x29, 0x03, 0x8d,
	0x8f, 0xe8, 0x1d, 0x9c, 0xf5, 0x00, 0xcb, 0x3a, 0x55, 0x62, 0xcc, 0xf8, 0x98, 0x8e, 0x00, 0x1d,
	0x43, 0xc4, 0x09, 0xee, 0xd0, 0x31, 0x1c, 0xb8, 0xaa, 0xd0, 0x8c, 0x4f, 0x3a, 0x69, 0x6b, 0xf9,
	0x7f, 0x56, 0x10, 0x8d, 0x70, 0x97, 0xce, 0xe0, 0xc4, 0x31, 0x8c, 0xfc, 0xd8, 0xc6, 0x46, 0xb2,
	0x91, 0xf8, 0x94, 0x3e, 0xc0, 0x85, 0x03, 0xc5, 0x95, 0x49, 0x95, 0x65, 0x95, 0x4e, 0x03, 0x5f,
	0x59, 0x69, 0x18, 0x9f, 0xd1, 0x7b, 0x38, 0xef, 0xf1, 0xa8, 0xa9, 0x8a, 0xb4, 0xfd, 0x6e, 0x24,
	0x3e, 0xa7, 0x01, 0x1c, 0xb9, 0x8e, 0x9f, 0x6c, 0x39, 0x61, 0x84, 0x7e, 0x12, 0x31, 0xee, 0xd1,
	0x5b, 0x38, 0x75, 0x5b, 0x50, 0xc3, 0x80, 0xb7, 0x83, 0xbd, 0xa0, 0x0b, 0x78, 0xd3, 0x83, 0xb7,
	0x1d, 0xed, 0xd3, 0x29, 0x1c, 0xbb, 0xc1, 0xbe, 0x0e, 0xa6, 0x6a, 0x9c, 0xe0, 0xcb, 0x4e, 0xf1,
	0xb1, 0xf9, 0xc1, 0x22, 0xc5, 0x57, 0x9d, 0x06, 0xc7, 0x89, 0xd6, 0x88, 0x9d, 0x49, 0xe2, 0x09,
	0xeb, 0x51, 0xaa, 0xd8, 0xfa, 0x43, 0x1f, 0x0f, 0xe8, 0x1c, 0x06, 0xee, 0xab, 0x49, 0x34, 0x9a,
	0x0c, 0xe5, 0x86, 0x52, 0xe7, 0x83, 0xa1, 0x91, 0xeb, 0x60, 0x3c, 0x94, 0x5f, 0xaf, 0xbe, 0xdc,
	0xdc, 0x2e, 0xe7, 0xab, 0x6c, 0xbd, 0x50, 0x5e, 0x36, 0x5b, 0xe6, 0xf3, 0xbc, 0x5a, 0xd4, 0x5e,
	0x7d, 0xb7, 0x2a, 0xb3, 0x62, 0xf1, 0xa9, 0xc9, 0xe7, 0x45, 0x39, 0x6b, 0xbc, 0x6c, 0x75, 0x7b,
	0x77, 0xed, 0xdd, 0x54, 0x5e, 0xbb, 0x7e, 0x5e, 0xbb, 0x7e, 0xd9, 0xee, 0xe6, 0xfa, 0xf9, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x3e, 0x34, 0xfa, 0xb0, 0x02, 0x00, 0x00,
}