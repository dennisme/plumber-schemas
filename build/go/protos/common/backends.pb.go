// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backends.proto

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
	return fileDescriptor_058942009bb99fa7, []int{0}
}

func init() {
	proto.RegisterEnum("protos.common.BackendType", BackendType_name, BackendType_value)
}

func init() { proto.RegisterFile("backends.proto", fileDescriptor_058942009bb99fa7) }

var fileDescriptor_058942009bb99fa7 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xf9, 0x53, 0x0a, 0x4c, 0x69, 0x99, 0x4e, 0x5b, 0xda, 0xd2, 0x52, 0x10, 0x47, 0x24,
	0xea, 0x03, 0x47, 0xc4, 0x61, 0xd7, 0x59, 0x8c, 0x65, 0xe2, 0xd8, 0x9e, 0xdd, 0x20, 0x7a, 0x59,
	0xc5, 0xae, 0xd5, 0x54, 0xd4, 0xb5, 0x15, 0x27, 0x07, 0x3e, 0x1c, 0xdf, 0x0d, 0x25, 0x69, 0x0e,
	0x2b, 0xfb, 0xb4, 0xda, 0xf7, 0x7e, 0xfb, 0x34, 0xfb, 0x34, 0xb0, 0x97, 0x4f, 0x8a, 0x3f, 0xe5,
	0xfd, 0x75, 0x7b, 0xd9, 0xcc, 0xea, 0x79, 0x4d, 0xbb, 0xab, 0xa3, 0xbd, 0x2c, 0xea, 0xaa, 0xaa,
	0xef, 0x3f, 0xfd, 0xdb, 0x82, 0x1d, 0xb9, 0x26, 0xf4, 0xdf, 0xa6, 0xa4, 0x37, 0x40, 0x52, 0xf8,
	0x91, 0x8a, 0x07, 0x56, 0xff, 0x4e, 0x94, 0x35, 0x31, 0x2b, 0x8d, 0x8f, 0x3a, 0x7a, 0x24, 0xbe,
	0x47, 0x02, 0x1f, 0xd3, 0x31, 0x1c, 0x38, 0x7a, 0x26, 0xa4, 0x0c, 0x35, 0x3e, 0xa1, 0xf7, 0x70,
	0xd6, 0x63, 0x58, 0xd6, 0x99, 0x12, 0x43, 0xc6, 0xa7, 0x74, 0x08, 0xe8, 0x00, 0x31, 0xa7, 0xb8,
	0x45, 0x47, 0xb0, 0xef, 0xaa, 0x42, 0x33, 0x3e, 0xeb, 0xa4, 0x2d, 0xe5, 0x87, 0xac, 0x30, 0x0e,
	0x70, 0x9b, 0xce, 0xe0, 0xd8, 0x01, 0x02, 0x3f, 0xb1, 0x89, 0x91, 0x6c, 0x24, 0x3e, 0xa7, 0x8f,
	0x70, 0xe1, 0x98, 0xe2, 0xca, 0x64, 0xca, 0xb2, 0xca, 0xc6, 0xa1, 0xaf, 0xac, 0x34, 0x8c, 0x2f,
	0xe8, 0x03, 0x9c, 0xf7, 0x30, 0x6a, 0xac, 0x62, 0x6d, 0x7f, 0x18, 0x89, 0x2f, 0xe9, 0x04, 0x0e,
	0x5d, 0xe2, 0x17, 0x5b, 0x4e, 0x19, 0xa1, 0xdf, 0x89, 0x19, 0x77, 0xe8, 0x1d, 0x9c, 0xba, 0x2d,
	0xa8, 0x41, 0xc8, 0x9b, 0xc1, 0x5e, 0xd1, 0x05, 0xbc, 0xed, 0xb1, 0x37, 0x1d, 0xed, 0xd2, 0x29,
	0x1c, 0xb9, 0xc1, 0xbe, 0x0e, 0xc7, 0x6a, 0x98, 0xe2, 0x5e, 0xa7, 0xf8, 0xc4, 0xfc, 0x64, 0x91,
	0xe1, 0xeb, 0x4e, 0x83, 0xc3, 0x54, 0x6b, 0xc4, 0xce, 0x24, 0xc9, 0x88, 0x75, 0x90, 0x29, 0xb6,
	0xfe, 0xc0, 0xc7, 0x7d, 0x3a, 0x87, 0x13, 0xf7, 0xd5, 0x28, 0x0e, 0x46, 0x03, 0xb9, 0x72, 0xa9,
	0xf3, 0xc1, 0xc8, 0xc8, 0x65, 0x30, 0x1e, 0xc8, 0x6f, 0x57, 0x5f, 0x6f, 0x6e, 0xe7, 0xd3, 0x45,
	0xbe, 0x5c, 0x28, 0x2f, 0x9f, 0xcc, 0x8b, 0x69, 0x51, 0xcf, 0x1a, 0xaf, 0xb9, 0x5b, 0x54, 0x79,
	0x39, 0xfb, 0xdc, 0x16, 0xd3, 0xb2, 0x9a, 0xb4, 0x5e, 0xbe, 0xb8, 0xbd, 0xbb, 0xf6, 0x6e, 0x6a,
	0x6f, 0xbd, 0x7e, 0xde, 0x7a, 0xfd, 0xf2, 0xed, 0xd5, 0xf5, 0xcb, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x86, 0x9d, 0xd6, 0x21, 0xa6, 0x02, 0x00, 0x00,
}