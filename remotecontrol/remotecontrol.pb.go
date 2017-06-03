// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remotecontrol.proto

/*
Package remotecontrol is a generated protocol buffer package.

It is generated from these files:
	remotecontrol.proto

It has these top-level messages:
	ControlRequest
	ControlReply
*/
package remotecontrol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ControlRequest struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *ControlRequest) Reset()                    { *m = ControlRequest{} }
func (m *ControlRequest) String() string            { return proto.CompactTextString(m) }
func (*ControlRequest) ProtoMessage()               {}
func (*ControlRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ControlRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ControlReply struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *ControlReply) Reset()                    { *m = ControlReply{} }
func (m *ControlReply) String() string            { return proto.CompactTextString(m) }
func (*ControlReply) ProtoMessage()               {}
func (*ControlReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ControlReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ControlRequest)(nil), "remotecontrol.ControlRequest")
	proto.RegisterType((*ControlReply)(nil), "remotecontrol.ControlReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RemoteController service

type RemoteControllerClient interface {
	Left(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error)
	Right(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error)
	Forward(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error)
	Backward(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error)
}

type remoteControllerClient struct {
	cc *grpc.ClientConn
}

func NewRemoteControllerClient(cc *grpc.ClientConn) RemoteControllerClient {
	return &remoteControllerClient{cc}
}

func (c *remoteControllerClient) Left(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error) {
	out := new(ControlReply)
	err := grpc.Invoke(ctx, "/remotecontrol.RemoteController/Left", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteControllerClient) Right(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error) {
	out := new(ControlReply)
	err := grpc.Invoke(ctx, "/remotecontrol.RemoteController/Right", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteControllerClient) Forward(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error) {
	out := new(ControlReply)
	err := grpc.Invoke(ctx, "/remotecontrol.RemoteController/Forward", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteControllerClient) Backward(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlReply, error) {
	out := new(ControlReply)
	err := grpc.Invoke(ctx, "/remotecontrol.RemoteController/Backward", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RemoteController service

type RemoteControllerServer interface {
	Left(context.Context, *ControlRequest) (*ControlReply, error)
	Right(context.Context, *ControlRequest) (*ControlReply, error)
	Forward(context.Context, *ControlRequest) (*ControlReply, error)
	Backward(context.Context, *ControlRequest) (*ControlReply, error)
}

func RegisterRemoteControllerServer(s *grpc.Server, srv RemoteControllerServer) {
	s.RegisterService(&_RemoteController_serviceDesc, srv)
}

func _RemoteController_Left_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteControllerServer).Left(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remotecontrol.RemoteController/Left",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteControllerServer).Left(ctx, req.(*ControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteController_Right_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteControllerServer).Right(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remotecontrol.RemoteController/Right",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteControllerServer).Right(ctx, req.(*ControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteController_Forward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteControllerServer).Forward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remotecontrol.RemoteController/Forward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteControllerServer).Forward(ctx, req.(*ControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteController_Backward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteControllerServer).Backward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remotecontrol.RemoteController/Backward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteControllerServer).Backward(ctx, req.(*ControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remotecontrol.RemoteController",
	HandlerType: (*RemoteControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Left",
			Handler:    _RemoteController_Left_Handler,
		},
		{
			MethodName: "Right",
			Handler:    _RemoteController_Right_Handler,
		},
		{
			MethodName: "Forward",
			Handler:    _RemoteController_Forward_Handler,
		},
		{
			MethodName: "Backward",
			Handler:    _RemoteController_Backward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remotecontrol.proto",
}

func init() { proto.RegisterFile("remotecontrol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0xcd, 0xcd,
	0x2f, 0x49, 0x4d, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0x52, 0xe3, 0xe2, 0x73, 0x86, 0x30, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x20, 0x1c, 0x25, 0x0d, 0x2e, 0x1e, 0xb8, 0xba, 0x82, 0x9c, 0x4a, 0x21, 0x09, 0x2e, 0xf6,
	0xe2, 0xd2, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0xb0, 0x3a, 0x8e, 0x20, 0x18, 0xd7, 0x68, 0x13, 0x13,
	0x97, 0x40, 0x10, 0xd8, 0x0e, 0xa8, 0x86, 0x9c, 0xd4, 0x22, 0x21, 0x17, 0x2e, 0x16, 0x9f, 0xd4,
	0xb4, 0x12, 0x21, 0x59, 0x3d, 0x54, 0x37, 0xa1, 0xda, 0x2d, 0x25, 0x8d, 0x4b, 0xba, 0x20, 0xa7,
	0x52, 0x89, 0x41, 0xc8, 0x95, 0x8b, 0x35, 0x28, 0x33, 0x3d, 0x83, 0x52, 0x63, 0xdc, 0xb9, 0xd8,
	0xdd, 0xf2, 0x8b, 0xca, 0x13, 0x8b, 0x52, 0x28, 0x34, 0xc8, 0x83, 0x8b, 0xc3, 0x29, 0x31, 0x39,
	0x9b, 0x72, 0x93, 0x92, 0xd8, 0xc0, 0x91, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xbe,
	0xdf, 0x74, 0xb3, 0x01, 0x00, 0x00,
}