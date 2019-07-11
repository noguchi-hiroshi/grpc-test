// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bistream/bistream.proto

package bistream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Request struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Finished             bool     `protobuf:"varint,2,opt,name=finished,proto3" json:"finished,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f7e07ada552956b, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Request) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f7e07ada552956b, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("proto/bistream/bistream.proto", fileDescriptor_7f7e07ada552956b) }

var fileDescriptor_7f7e07ada552956b = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xca, 0x2c, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x85, 0x33, 0xf4, 0xc0, 0xe2, 0x4a,
	0xf6, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9,
	0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x14,
	0x17, 0x47, 0x5a, 0x66, 0x5e, 0x66, 0x71, 0x46, 0x6a, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47,
	0x10, 0x9c, 0xaf, 0xa4, 0xc2, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0xdb,
	0x04, 0x23, 0x3d, 0x2e, 0x0e, 0xa7, 0xcc, 0x60, 0xb0, 0xc5, 0x42, 0x4a, 0x5c, 0xac, 0xee, 0x45,
	0xa9, 0xa9, 0x25, 0x42, 0x1c, 0x7a, 0x50, 0xab, 0xa5, 0x38, 0xf5, 0x60, 0x66, 0x28, 0x31, 0x68,
	0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0x5d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x41, 0x1e,
	0x24, 0x56, 0xbe, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BiStreamClient is the client API for BiStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BiStreamClient interface {
	Greet(ctx context.Context, opts ...grpc.CallOption) (BiStream_GreetClient, error)
}

type biStreamClient struct {
	cc *grpc.ClientConn
}

func NewBiStreamClient(cc *grpc.ClientConn) BiStreamClient {
	return &biStreamClient{cc}
}

func (c *biStreamClient) Greet(ctx context.Context, opts ...grpc.CallOption) (BiStream_GreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BiStream_serviceDesc.Streams[0], "/BiStream/Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &biStreamGreetClient{stream}
	return x, nil
}

type BiStream_GreetClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type biStreamGreetClient struct {
	grpc.ClientStream
}

func (x *biStreamGreetClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *biStreamGreetClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BiStreamServer is the server API for BiStream service.
type BiStreamServer interface {
	Greet(BiStream_GreetServer) error
}

func RegisterBiStreamServer(s *grpc.Server, srv BiStreamServer) {
	s.RegisterService(&_BiStream_serviceDesc, srv)
}

func _BiStream_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BiStreamServer).Greet(&biStreamGreetServer{stream})
}

type BiStream_GreetServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type biStreamGreetServer struct {
	grpc.ServerStream
}

func (x *biStreamGreetServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *biStreamGreetServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BiStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BiStream",
	HandlerType: (*BiStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet",
			Handler:       _BiStream_Greet_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/bistream/bistream.proto",
}
