// Code generated by protoc-gen-go.
// source: kvs/kvs.proto
// DO NOT EDIT!

/*
Package kvs is a generated protocol buffer package.

It is generated from these files:
	kvs/kvs.proto

It has these top-level messages:
	Key
	Value
	Entry
	PutResponse
	DeleteResponse
	ListRequest
*/
package kvs

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

// Key defines a simple key in the KVS.
type Key struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Key) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Value represents a share in CGSS.
type Value struct {
	Field       string   `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	GroupKey    string   `protobuf:"bytes,2,opt,name=group_key,json=groupKey" json:"group_key,omitempty"`
	GroupShares []string `protobuf:"bytes,3,rep,name=group_shares,json=groupShares" json:"group_shares,omitempty"`
	DataKey     string   `protobuf:"bytes,4,opt,name=data_key,json=dataKey" json:"data_key,omitempty"`
	DataShares  []string `protobuf:"bytes,5,rep,name=data_shares,json=dataShares" json:"data_shares,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Value) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Value) GetGroupKey() string {
	if m != nil {
		return m.GroupKey
	}
	return ""
}

func (m *Value) GetGroupShares() []string {
	if m != nil {
		return m.GroupShares
	}
	return nil
}

func (m *Value) GetDataKey() string {
	if m != nil {
		return m.DataKey
	}
	return ""
}

func (m *Value) GetDataShares() []string {
	if m != nil {
		return m.DataShares
	}
	return nil
}

// Entry defines a pair of key and value as an entry in the KVS.
type Entry struct {
	Key   *Key   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Entry) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// Define empty requests/responses.
type PutResponse struct {
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Key)(nil), "kvs.Key")
	proto.RegisterType((*Value)(nil), "kvs.Value")
	proto.RegisterType((*Entry)(nil), "kvs.Entry")
	proto.RegisterType((*PutResponse)(nil), "kvs.PutResponse")
	proto.RegisterType((*DeleteResponse)(nil), "kvs.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "kvs.ListRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Kvs service

type KvsClient interface {
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Put(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*PutResponse, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Kvs_ListClient, error)
}

type kvsClient struct {
	cc *grpc.ClientConn
}

func NewKvsClient(cc *grpc.ClientConn) KvsClient {
	return &kvsClient{cc}
}

func (c *kvsClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := grpc.Invoke(ctx, "/kvs.Kvs/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvsClient) Put(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/kvs.Kvs/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvsClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/kvs.Kvs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Kvs_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Kvs_serviceDesc.Streams[0], c.cc, "/kvs.Kvs/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &kvsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kvs_ListClient interface {
	Recv() (*Key, error)
	grpc.ClientStream
}

type kvsListClient struct {
	grpc.ClientStream
}

func (x *kvsListClient) Recv() (*Key, error) {
	m := new(Key)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Kvs service

type KvsServer interface {
	Get(context.Context, *Key) (*Value, error)
	Put(context.Context, *Entry) (*PutResponse, error)
	Delete(context.Context, *Key) (*DeleteResponse, error)
	List(*ListRequest, Kvs_ListServer) error
}

func RegisterKvsServer(s *grpc.Server, srv KvsServer) {
	s.RegisterService(&_Kvs_serviceDesc, srv)
}

func _Kvs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.Kvs/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvsServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvs_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvsServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.Kvs/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvsServer).Put(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.Kvs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvsServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvs_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KvsServer).List(m, &kvsListServer{stream})
}

type Kvs_ListServer interface {
	Send(*Key) error
	grpc.ServerStream
}

type kvsListServer struct {
	grpc.ServerStream
}

func (x *kvsListServer) Send(m *Key) error {
	return x.ServerStream.SendMsg(m)
}

var _Kvs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvs.Kvs",
	HandlerType: (*KvsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Kvs_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Kvs_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Kvs_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Kvs_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kvs/kvs.proto",
}

func init() { proto.RegisterFile("kvs/kvs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x13, 0xb7, 0xa9, 0xed, 0xac, 0x95, 0x32, 0x7a, 0x68, 0x23, 0x62, 0x5d, 0x50, 0xea,
	0xa5, 0x4a, 0xfd, 0x0b, 0x16, 0x0f, 0xf5, 0x50, 0x56, 0xf0, 0x2a, 0x91, 0x8e, 0x5a, 0x52, 0x9b,
	0x98, 0xdd, 0x04, 0xf2, 0x53, 0xbc, 0xf8, 0x5b, 0x65, 0x67, 0x4b, 0x08, 0xde, 0x66, 0xdf, 0xe3,
	0x7b, 0xf3, 0x86, 0x85, 0x41, 0x5a, 0x99, 0xdb, 0xb4, 0x32, 0xb3, 0xbc, 0xc8, 0x6c, 0x86, 0x22,
	0xad, 0x8c, 0x1a, 0x83, 0x58, 0x52, 0x8d, 0x08, 0x9d, 0x5d, 0xf2, 0x45, 0xa3, 0x70, 0x12, 0x4e,
	0xfb, 0x9a, 0x67, 0xf5, 0x13, 0x42, 0xf4, 0x92, 0x6c, 0x4b, 0xc2, 0x53, 0x88, 0xde, 0x37, 0xb4,
	0x5d, 0xef, 0x6d, 0xff, 0xc0, 0x33, 0xe8, 0x7f, 0x14, 0x59, 0x99, 0xbf, 0xa6, 0x54, 0x8f, 0x0e,
	0xd8, 0xe9, 0xb1, 0xe0, 0x02, 0x2f, 0xe1, 0xc8, 0x9b, 0xe6, 0x33, 0x29, 0xc8, 0x8c, 0xc4, 0x44,
	0x4c, 0xfb, 0x5a, 0xb2, 0xf6, 0xcc, 0x12, 0x8e, 0xa1, 0xb7, 0x4e, 0x6c, 0xc2, 0x78, 0x87, 0xf1,
	0x43, 0xf7, 0x76, 0xf4, 0x05, 0x48, 0xb6, 0xf6, 0x70, 0xc4, 0x30, 0x38, 0xc9, 0xb3, 0x6a, 0x01,
	0xd1, 0x62, 0x67, 0x8b, 0x1a, 0x63, 0x10, 0x8e, 0x77, 0xc5, 0xe4, 0xbc, 0x37, 0x73, 0xd7, 0x2d,
	0xa9, 0xd6, 0x4e, 0xc4, 0x09, 0x44, 0x95, 0xeb, 0xcf, 0xe5, 0xe4, 0x1c, 0xd8, 0xe5, 0x8b, 0xb4,
	0x37, 0xd4, 0x00, 0xe4, 0xaa, 0xb4, 0x9a, 0x4c, 0x9e, 0xed, 0x0c, 0xa9, 0x21, 0x1c, 0x3f, 0xd0,
	0x96, 0x2c, 0x35, 0xca, 0x00, 0xe4, 0xd3, 0xc6, 0x58, 0x4d, 0xdf, 0x25, 0x19, 0x3b, 0xff, 0x0d,
	0x41, 0x2c, 0x2b, 0x83, 0xe7, 0x20, 0x1e, 0xc9, 0x62, 0xb3, 0x2f, 0x6e, 0x65, 0xab, 0x00, 0xaf,
	0x40, 0xac, 0x4a, 0x8b, 0x5e, 0xe4, 0x9e, 0xf1, 0x90, 0xe7, 0xf6, 0xb2, 0x00, 0x6f, 0xa0, 0xeb,
	0xd7, 0xb5, 0x82, 0x4e, 0x78, 0xfa, 0xd7, 0x22, 0xc0, 0x6b, 0xe8, 0xb8, 0x1e, 0xe8, 0x63, 0x5a,
	0x95, 0xe2, 0x06, 0x55, 0xc1, 0x5d, 0xf8, 0xd6, 0xe5, 0xaf, 0xbd, 0xff, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x85, 0x0b, 0xed, 0x8c, 0xeb, 0x01, 0x00, 0x00,
}