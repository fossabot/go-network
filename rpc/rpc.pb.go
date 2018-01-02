// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	Payload
	MultiAddress
	MultiAddresses
	Nothing
*/
package rpc

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

// A Payload is a message containing application-specific data.
type Payload struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Payload) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Payload) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// A MultiAddress is the public multiaddress of a Node in the overlay network.
// It provides the Republic address of the Node, as well as the network
// address.
type MultiAddress struct {
	Multi string `protobuf:"bytes,1,opt,name=multi" json:"multi,omitempty"`
}

func (m *MultiAddress) Reset()                    { *m = MultiAddress{} }
func (m *MultiAddress) String() string            { return proto.CompactTextString(m) }
func (*MultiAddress) ProtoMessage()               {}
func (*MultiAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MultiAddress) GetMulti() string {
	if m != nil {
		return m.Multi
	}
	return ""
}

// MultiAddresses are public multiaddress of multiple Nodes in the overlay
// network.
type MultiAddresses struct {
	Multis []*MultiAddress `protobuf:"bytes,1,rep,name=multis" json:"multis,omitempty"`
}

func (m *MultiAddresses) Reset()                    { *m = MultiAddresses{} }
func (m *MultiAddresses) String() string            { return proto.CompactTextString(m) }
func (*MultiAddresses) ProtoMessage()               {}
func (*MultiAddresses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MultiAddresses) GetMultis() []*MultiAddress {
	if m != nil {
		return m.Multis
	}
	return nil
}

// Nothing is in this message. It is used to send nothing, or signal a
// successful response.
type Nothing struct {
}

func (m *Nothing) Reset()                    { *m = Nothing{} }
func (m *Nothing) String() string            { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()               {}
func (*Nothing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Payload)(nil), "rpc.Payload")
	proto.RegisterType((*MultiAddress)(nil), "rpc.MultiAddress")
	proto.RegisterType((*MultiAddresses)(nil), "rpc.MultiAddresses")
	proto.RegisterType((*Nothing)(nil), "rpc.Nothing")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	// Ping the connection and swap MultiAddresses.
	Ping(ctx context.Context, in *MultiAddress, opts ...grpc.CallOption) (*MultiAddress, error)
	// Get all peers connected to the Node.
	Peers(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*MultiAddresses, error)
	// Send a payload to a target Node.
	Send(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Nothing, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Ping(ctx context.Context, in *MultiAddress, opts ...grpc.CallOption) (*MultiAddress, error) {
	out := new(MultiAddress)
	err := grpc.Invoke(ctx, "/rpc.Node/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Peers(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*MultiAddresses, error) {
	out := new(MultiAddresses)
	err := grpc.Invoke(ctx, "/rpc.Node/Peers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Send(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := grpc.Invoke(ctx, "/rpc.Node/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	// Ping the connection and swap MultiAddresses.
	Ping(context.Context, *MultiAddress) (*MultiAddress, error)
	// Get all peers connected to the Node.
	Peers(context.Context, *Nothing) (*MultiAddresses, error)
	// Send a payload to a target Node.
	Send(context.Context, *Payload) (*Nothing, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Node/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Ping(ctx, req.(*MultiAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Node/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Peers(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Node/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Send(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Node_Ping_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _Node_Peers_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Node_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4f, 0x84, 0x30,
	0x14, 0x84, 0x53, 0xb6, 0xbb, 0x1b, 0x9e, 0x9b, 0x4d, 0x7c, 0x7a, 0x20, 0x9c, 0x48, 0xe3, 0x01,
	0x8d, 0x72, 0xc0, 0xa3, 0x27, 0x7f, 0x80, 0x84, 0xe0, 0x2f, 0xa8, 0xbc, 0x06, 0x9b, 0x20, 0x25,
	0x6d, 0x3d, 0x78, 0xf7, 0x87, 0x1b, 0x6a, 0x0f, 0x6c, 0xb8, 0x75, 0x66, 0xbe, 0x69, 0x5e, 0x06,
	0x52, 0x3b, 0xf7, 0xd5, 0x6c, 0x8d, 0x37, 0xb8, 0xb3, 0x73, 0x2f, 0x9e, 0xe0, 0xd8, 0xca, 0x9f,
	0xd1, 0x48, 0xc2, 0x33, 0x24, 0x9a, 0x32, 0x56, 0xb0, 0x32, 0xed, 0x12, 0x4d, 0x88, 0xc0, 0x49,
	0x7a, 0x99, 0x25, 0xc1, 0x09, 0x6f, 0x71, 0x07, 0xa7, 0xb7, 0xef, 0xd1, 0xeb, 0x57, 0x22, 0xab,
	0x9c, 0xc3, 0x5b, 0xd8, 0x7f, 0x2d, 0x3a, 0xd6, 0xfe, 0x85, 0x78, 0x81, 0xf3, 0x9a, 0x52, 0x0e,
	0xef, 0xe1, 0x10, 0x22, 0x97, 0xb1, 0x62, 0x57, 0x5e, 0xd5, 0xd7, 0xd5, 0x72, 0xc7, 0x1a, 0xea,
	0x22, 0x20, 0x52, 0x38, 0x36, 0xc6, 0x7f, 0xea, 0x69, 0xa8, 0x7f, 0x19, 0xf0, 0xc6, 0x90, 0xc2,
	0x47, 0xe0, 0xad, 0x9e, 0x06, 0xdc, 0xd6, 0xf2, 0xad, 0x85, 0x0f, 0xb0, 0x6f, 0x95, 0xb2, 0x0e,
	0x4f, 0x21, 0x8b, 0xbf, 0xe5, 0x37, 0x1b, 0x52, 0x39, 0x14, 0xc0, 0xdf, 0xd5, 0x44, 0x11, 0x8d,
	0x53, 0xe4, 0x17, 0xc5, 0x8f, 0x43, 0xd8, 0xeb, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xd8,
	0x34, 0xe6, 0x3c, 0x01, 0x00, 0x00,
}
