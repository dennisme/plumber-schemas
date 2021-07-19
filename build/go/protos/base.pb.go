// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x4e, 0xc2, 0x30,
	0x1c, 0xc5, 0xe3, 0x85, 0x5f, 0x25, 0x18, 0x6d, 0x34, 0xc2, 0xfc, 0xc2, 0x07, 0x90, 0x19, 0x25,
	0xde, 0x99, 0x28, 0x18, 0xb9, 0x31, 0xd1, 0x80, 0x44, 0xe3, 0x5d, 0xb7, 0xfd, 0x03, 0x4b, 0xc6,
	0x3a, 0xdb, 0x4e, 0xe3, 0xdb, 0xfa, 0x28, 0x66, 0xeb, 0x4a, 0xb7, 0x76, 0xf3, 0x0a, 0x7a, 0xce,
	0xf9, 0xff, 0x72, 0xfa, 0x01, 0x08, 0x79, 0x84, 0x43, 0x3f, 0x61, 0x54, 0x50, 0xbc, 0x91, 0x7f,
	0x70, 0xa7, 0xed, 0xd3, 0x38, 0x06, 0x5f, 0x48, 0xd9, 0x41, 0x0c, 0x48, 0x50, 0x7c, 0x6f, 0x7d,
	0xb3, 0x50, 0x80, 0x5a, 0x30, 0x88, 0xc8, 0x8f, 0x5c, 0x5c, 0xfd, 0x6e, 0xa2, 0xf6, 0x4b, 0x94,
	0x2e, 0x3d, 0x60, 0x53, 0x60, 0x5f, 0xc0, 0xf0, 0x3b, 0xda, 0x1b, 0x83, 0xb8, 0x8f, 0xa2, 0x91,
	0xc4, 0x85, 0x34, 0xe6, 0xb8, 0x27, 0xe3, 0xbc, 0x6f, 0x59, 0x13, 0xf8, 0x4c, 0x81, 0x0b, 0xe7,
	0xfc, 0x9f, 0x04, 0x4f, 0x68, 0xcc, 0x01, 0x3f, 0xa1, 0xf6, 0x18, 0x84, 0x76, 0xf0, 0x71, 0x69,
	0x46, 0xcb, 0x8a, 0x78, 0xd2, 0xe0, 0x16, 0xb4, 0x19, 0xda, 0x1d, 0x31, 0x20, 0x02, 0x4a, 0xc0,
	0x33, 0x35, 0x62, 0x3a, 0x8a, 0xd9, 0x6b, 0x0e, 0x14, 0xd8, 0x67, 0xb4, 0xf3, 0x0a, 0xbc, 0xdc,
	0x72, 0xd5, 0xa3, 0xaa, 0x2b, 0xe4, 0x69, 0x93, 0xad, 0x7b, 0xce, 0x92, 0xa0, 0xa1, 0xa7, 0xe9,
	0x58, 0x3d, 0xed, 0x80, 0xc6, 0x3e, 0x40, 0x04, 0xf5, 0x58, 0xd3, 0xb1, 0xb0, 0x76, 0xa0, 0xc0,
	0x0e, 0xd1, 0xf6, 0x54, 0x10, 0x26, 0x26, 0x40, 0x02, 0xdc, 0x51, 0xf1, 0x95, 0xa4, 0x40, 0xdd,
	0x1a, 0x47, 0x12, 0x2e, 0xd7, 0xf0, 0x2d, 0xda, 0x9a, 0x0a, 0x9a, 0xe4, 0x88, 0x43, 0x1d, 0x94,
	0x8a, 0x22, 0x74, 0x6c, 0xa3, 0xa8, 0x30, 0x40, 0xeb, 0x6f, 0xd9, 0x73, 0xc5, 0xfb, 0x2a, 0x92,
	0x2f, 0xd5, 0xe0, 0x81, 0xa1, 0x16, 0x53, 0x8f, 0xa8, 0x25, 0xef, 0x74, 0x92, 0xbd, 0x6e, 0xec,
	0x54, 0x2f, 0x3a, 0x17, 0x15, 0xe1, 0xa8, 0xd6, 0xd3, 0x1c, 0x79, 0xe6, 0x06, 0xa7, 0x24, 0x5a,
	0x9c, 0x8a, 0x57, 0x70, 0xee, 0xb2, 0x83, 0xcc, 0x76, 0x96, 0x51, 0x8c, 0xcd, 0x96, 0x18, 0xdd,
	0x1a, 0x47, 0x37, 0x91, 0xd7, 0x64, 0x34, 0x29, 0x89, 0x56, 0x93, 0x8a, 0x27, 0x39, 0xc3, 0x9b,
	0x8f, 0xc1, 0x3c, 0x14, 0x8b, 0xd4, 0xeb, 0xfb, 0x74, 0xe9, 0x7a, 0x44, 0xf8, 0x0b, 0x9f, 0xb2,
	0xc4, 0x4d, 0xe4, 0xcf, 0xfe, 0x82, 0xfb, 0x0b, 0x58, 0x12, 0xee, 0x7a, 0x69, 0x18, 0x05, 0xee,
	0x9c, 0xba, 0x92, 0xe5, 0xc9, 0xff, 0x95, 0xeb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xfb,
	0x57, 0xfb, 0x6c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlumberServerClient is the client API for PlumberServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlumberServerClient interface {
	// List configured/known connections
	GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error)
	// Start reading data from a connection
	StartRead(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (PlumberServer_StartReadClient, error)
	// Stop reading data from a connection
	StopRead(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*StopReadResponse, error)
	// Write data to a connection
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	// Create a data relay from plumber server to the Batch platform
	CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error)
	// Update a relay (such as API token) - relay will be interrupted!
	UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error)
	StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error)
	// Delete an existing relay
	DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error)
}

type plumberServerClient struct {
	cc *grpc.ClientConn
}

func NewPlumberServerClient(cc *grpc.ClientConn) PlumberServerClient {
	return &plumberServerClient{cc}
}

func (c *plumberServerClient) GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error) {
	out := new(GetAllConnectionsResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error) {
	out := new(GetConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error) {
	out := new(CreateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error) {
	out := new(TestConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/TestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error) {
	out := new(UpdateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error) {
	out := new(DeleteConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StartRead(ctx context.Context, in *StartReadRequest, opts ...grpc.CallOption) (PlumberServer_StartReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlumberServer_serviceDesc.Streams[0], "/protos.PlumberServer/StartRead", opts...)
	if err != nil {
		return nil, err
	}
	x := &plumberServerStartReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlumberServer_StartReadClient interface {
	Recv() (*StartReadResponse, error)
	grpc.ClientStream
}

type plumberServerStartReadClient struct {
	grpc.ClientStream
}

func (x *plumberServerStartReadClient) Recv() (*StartReadResponse, error) {
	m := new(StartReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *plumberServerClient) StopRead(ctx context.Context, in *StopReadRequest, opts ...grpc.CallOption) (*StopReadResponse, error) {
	out := new(StopReadResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error) {
	out := new(CreateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error) {
	out := new(UpdateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error) {
	out := new(StopRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error) {
	out := new(DeleteRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlumberServerServer is the server API for PlumberServer service.
type PlumberServerServer interface {
	// List configured/known connections
	GetAllConnections(context.Context, *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(context.Context, *GetConnectionRequest) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error)
	// Start reading data from a connection
	StartRead(*StartReadRequest, PlumberServer_StartReadServer) error
	// Stop reading data from a connection
	StopRead(context.Context, *StopReadRequest) (*StopReadResponse, error)
	// Write data to a connection
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	// Create a data relay from plumber server to the Batch platform
	CreateRelay(context.Context, *CreateRelayRequest) (*CreateRelayResponse, error)
	// Update a relay (such as API token) - relay will be interrupted!
	UpdateRelay(context.Context, *UpdateRelayRequest) (*UpdateRelayResponse, error)
	StopRelay(context.Context, *StopRelayRequest) (*StopRelayResponse, error)
	// Delete an existing relay
	DeleteRelay(context.Context, *DeleteRelayRequest) (*DeleteRelayResponse, error)
}

// UnimplementedPlumberServerServer can be embedded to have forward compatible implementations.
type UnimplementedPlumberServerServer struct {
}

func (*UnimplementedPlumberServerServer) GetAllConnections(ctx context.Context, req *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConnections not implemented")
}
func (*UnimplementedPlumberServerServer) GetConnection(ctx context.Context, req *GetConnectionRequest) (*GetConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (*UnimplementedPlumberServerServer) CreateConnection(ctx context.Context, req *CreateConnectionRequest) (*CreateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) TestConnection(ctx context.Context, req *TestConnectionRequest) (*TestConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateConnection(ctx context.Context, req *UpdateConnectionRequest) (*UpdateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteConnection(ctx context.Context, req *DeleteConnectionRequest) (*DeleteConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (*UnimplementedPlumberServerServer) StartRead(req *StartReadRequest, srv PlumberServer_StartReadServer) error {
	return status.Errorf(codes.Unimplemented, "method StartRead not implemented")
}
func (*UnimplementedPlumberServerServer) StopRead(ctx context.Context, req *StopReadRequest) (*StopReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRead not implemented")
}
func (*UnimplementedPlumberServerServer) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedPlumberServerServer) CreateRelay(ctx context.Context, req *CreateRelayRequest) (*CreateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateRelay(ctx context.Context, req *UpdateRelayRequest) (*UpdateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) StopRelay(ctx context.Context, req *StopRelayRequest) (*StopRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRelay not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteRelay(ctx context.Context, req *DeleteRelayRequest) (*DeleteRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelay not implemented")
}

func RegisterPlumberServerServer(s *grpc.Server, srv PlumberServerServer) {
	s.RegisterService(&_PlumberServer_serviceDesc, srv)
}

func _PlumberServer_GetAllConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllConnections(ctx, req.(*GetAllConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetConnection(ctx, req.(*GetConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateConnection(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/TestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).TestConnection(ctx, req.(*TestConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateConnection(ctx, req.(*UpdateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StartRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlumberServerServer).StartRead(m, &plumberServerStartReadServer{stream})
}

type PlumberServer_StartReadServer interface {
	Send(*StartReadResponse) error
	grpc.ServerStream
}

type plumberServerStartReadServer struct {
	grpc.ServerStream
}

func (x *plumberServerStartReadServer) Send(m *StartReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PlumberServer_StopRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopRead(ctx, req.(*StopReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateRelay(ctx, req.(*CreateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateRelay(ctx, req.(*UpdateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StopRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopRelay(ctx, req.(*StopRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteRelay(ctx, req.(*DeleteRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlumberServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PlumberServer",
	HandlerType: (*PlumberServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllConnections",
			Handler:    _PlumberServer_GetAllConnections_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _PlumberServer_GetConnection_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _PlumberServer_CreateConnection_Handler,
		},
		{
			MethodName: "TestConnection",
			Handler:    _PlumberServer_TestConnection_Handler,
		},
		{
			MethodName: "UpdateConnection",
			Handler:    _PlumberServer_UpdateConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _PlumberServer_DeleteConnection_Handler,
		},
		{
			MethodName: "StopRead",
			Handler:    _PlumberServer_StopRead_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _PlumberServer_Write_Handler,
		},
		{
			MethodName: "CreateRelay",
			Handler:    _PlumberServer_CreateRelay_Handler,
		},
		{
			MethodName: "UpdateRelay",
			Handler:    _PlumberServer_UpdateRelay_Handler,
		},
		{
			MethodName: "StopRelay",
			Handler:    _PlumberServer_StopRelay_Handler,
		},
		{
			MethodName: "DeleteRelay",
			Handler:    _PlumberServer_DeleteRelay_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartRead",
			Handler:       _PlumberServer_StartRead_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "base.proto",
}
