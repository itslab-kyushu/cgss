// Code generated by protoc-gen-go.
// source: kvs/kvs.proto
// DO NOT EDIT!

/*
Package kvs is a generated protocol buffer package.

It is generated from these files:
	kvs/kvs.proto

It has these top-level messages:
	Key
	Field
	Share
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

// Field represents a finite field Z/pZ.
type Field struct {
	Prime string `protobuf:"bytes,1,opt,name=prime" json:"prime,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Field) GetPrime() string {
	if m != nil {
		return m.Prime
	}
	return ""
}

// Share represents a basic share in Shamir's secret sharing scheme.
type Share struct {
	Key   string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
	Field *Field   `protobuf:"bytes,3,opt,name=field" json:"field,omitempty"`
}

func (m *Share) Reset()                    { *m = Share{} }
func (m *Share) String() string            { return proto.CompactTextString(m) }
func (*Share) ProtoMessage()               {}
func (*Share) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Share) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Share) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Share) GetField() *Field {
	if m != nil {
		return m.Field
	}
	return nil
}

// Value represents a share in CGSS.
type Value struct {
	GroupShare []*Share `protobuf:"bytes,1,rep,name=group_share,json=groupShare" json:"group_share,omitempty"`
	DataShare  *Share   `protobuf:"bytes,2,opt,name=data_share,json=dataShare" json:"data_share,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Value) GetGroupShare() []*Share {
	if m != nil {
		return m.GroupShare
	}
	return nil
}

func (m *Value) GetDataShare() *Share {
	if m != nil {
		return m.DataShare
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
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Key)(nil), "kvs.Key")
	proto.RegisterType((*Field)(nil), "kvs.Field")
	proto.RegisterType((*Share)(nil), "kvs.Share")
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
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xd1, 0x4a, 0xf3, 0x40,
	0x10, 0x85, 0x93, 0x6e, 0x53, 0xfe, 0x4e, 0xe8, 0x4f, 0x19, 0xbd, 0x88, 0x81, 0xc2, 0xb2, 0xa0,
	0xb4, 0x08, 0x55, 0xea, 0x2b, 0x58, 0xbd, 0xa8, 0x17, 0x35, 0x82, 0xb7, 0x25, 0xd2, 0x51, 0x4b,
	0x6a, 0x13, 0xb3, 0x9b, 0x40, 0x5e, 0xc6, 0x67, 0x95, 0x9d, 0x2d, 0x69, 0xf0, 0xee, 0xec, 0x9c,
	0x99, 0x73, 0xbe, 0x40, 0x60, 0x94, 0xd5, 0xfa, 0x26, 0xab, 0xf5, 0xbc, 0x28, 0x73, 0x93, 0xa3,
	0xc8, 0x6a, 0xad, 0x2e, 0x40, 0xac, 0xa8, 0x41, 0x84, 0xfe, 0x21, 0xfd, 0xa2, 0xc8, 0x97, 0xfe,
	0x74, 0x98, 0xb0, 0x56, 0x13, 0x08, 0x1e, 0x76, 0xb4, 0xdf, 0xe2, 0x39, 0x04, 0x45, 0xb9, 0x6b,
	0x5d, 0xf7, 0x50, 0xcf, 0x10, 0xbc, 0x7c, 0xa6, 0x25, 0xe1, 0x18, 0x44, 0x46, 0xcd, 0xd1, 0xb4,
	0xd2, 0x1e, 0xd4, 0xe9, 0xbe, 0xa2, 0xa8, 0x27, 0x85, 0x3d, 0xe0, 0x07, 0x4a, 0x08, 0xde, 0x6d,
	0x5e, 0x24, 0xa4, 0x3f, 0x0d, 0x17, 0x30, 0xb7, 0x28, 0xdc, 0x90, 0x38, 0x43, 0x6d, 0x20, 0x78,
	0xe5, 0xd5, 0x6b, 0x08, 0x3f, 0xca, 0xbc, 0x2a, 0x36, 0xda, 0x36, 0x44, 0xbe, 0x14, 0xed, 0x01,
	0x77, 0x26, 0xc0, 0xb6, 0xeb, 0x9f, 0x01, 0x6c, 0x53, 0x93, 0x1e, 0x77, 0x7b, 0x9d, 0x70, 0xb7,
	0x3b, 0xb4, 0x2e, 0x4b, 0xb5, 0x84, 0x60, 0x79, 0x30, 0x65, 0x83, 0xf1, 0x89, 0x39, 0x5c, 0xfc,
	0xe3, 0xe5, 0x15, 0x35, 0x8e, 0x5e, 0x9e, 0xe8, 0x4f, 0x51, 0xcc, 0x75, 0xfc, 0x12, 0x35, 0x82,
	0x70, 0x5d, 0x99, 0x84, 0x74, 0x91, 0x1f, 0x34, 0xa9, 0x31, 0xfc, 0xbf, 0xa7, 0x3d, 0x19, 0x6a,
	0x27, 0x23, 0x08, 0x9f, 0x76, 0xda, 0x24, 0xf4, 0x5d, 0x91, 0x36, 0x8b, 0x1f, 0x1f, 0xc4, 0xaa,
	0xd6, 0x38, 0x01, 0xf1, 0x48, 0x06, 0xdb, 0xbe, 0xb8, 0x93, 0xad, 0x3c, 0xbc, 0x04, 0xb1, 0xae,
	0x0c, 0xba, 0x21, 0x73, 0xc6, 0x63, 0xd6, 0xdd, 0x32, 0x0f, 0x67, 0x30, 0x70, 0x75, 0x9d, 0xa0,
	0x33, 0x56, 0x7f, 0x28, 0x3c, 0xbc, 0x82, 0xbe, 0xe5, 0x40, 0x17, 0xd3, 0x41, 0x8a, 0xdb, 0x53,
	0xe5, 0xdd, 0xfa, 0x6f, 0x03, 0xfe, 0x23, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x76, 0x83,
	0x19, 0x10, 0x22, 0x02, 0x00, 0x00,
}
