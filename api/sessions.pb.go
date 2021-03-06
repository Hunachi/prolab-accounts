// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sessions.proto

package api_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Session struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ddf231f23319645b, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type GetSessionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionRequest) Reset()         { *m = GetSessionRequest{} }
func (m *GetSessionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionRequest) ProtoMessage()    {}
func (*GetSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ddf231f23319645b, []int{1}
}
func (m *GetSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionRequest.Unmarshal(m, b)
}
func (m *GetSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionRequest.Marshal(b, m, deterministic)
}
func (dst *GetSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionRequest.Merge(dst, src)
}
func (m *GetSessionRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionRequest.Size(m)
}
func (m *GetSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionRequest proto.InternalMessageInfo

type CreateSessionRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionRequest) Reset()         { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()    {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ddf231f23319645b, []int{2}
}
func (m *CreateSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionRequest.Unmarshal(m, b)
}
func (m *CreateSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionRequest.Merge(dst, src)
}
func (m *CreateSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionRequest.Size(m)
}
func (m *CreateSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionRequest proto.InternalMessageInfo

func (m *CreateSessionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSessionRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeleteSessionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionRequest) Reset()         { *m = DeleteSessionRequest{} }
func (m *DeleteSessionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionRequest) ProtoMessage()    {}
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sessions_ddf231f23319645b, []int{3}
}
func (m *DeleteSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionRequest.Unmarshal(m, b)
}
func (m *DeleteSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionRequest.Merge(dst, src)
}
func (m *DeleteSessionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionRequest.Size(m)
}
func (m *DeleteSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Session)(nil), "com.github.ProgrammingLab.prolab_accounts.api.Session")
	proto.RegisterType((*GetSessionRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.GetSessionRequest")
	proto.RegisterType((*CreateSessionRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.CreateSessionRequest")
	proto.RegisterType((*DeleteSessionRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.DeleteSessionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	GetSession(context.Context, *GetSessionRequest) (*Session, error)
	CreateSession(context.Context, *CreateSessionRequest) (*Session, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*empty.Empty, error)
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.SessionService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.ProgrammingLab.prolab_accounts.api.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _SessionService_GetSession_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _SessionService_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sessions.proto",
}

func init() { proto.RegisterFile("sessions.proto", fileDescriptor_sessions_ddf231f23319645b) }

var fileDescriptor_sessions_ddf231f23319645b = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x95, 0xda, 0x8e, 0xb4, 0xd0, 0xb5, 0x94, 0x12, 0x15, 0x24, 0xa7, 0x22, 0xb8,
	0x01, 0x05, 0xcf, 0x62, 0xfd, 0x83, 0xe0, 0x41, 0xda, 0x9b, 0x97, 0x32, 0x49, 0xc7, 0xb8, 0xd0,
	0x64, 0xd7, 0xec, 0x46, 0xf1, 0x24, 0xf8, 0x0a, 0x9e, 0x3c, 0xf9, 0x48, 0x1e, 0x7c, 0x05, 0x1f,
	0x44, 0x9a, 0x6e, 0xad, 0xd1, 0x5e, 0x82, 0xb7, 0xcc, 0x37, 0x93, 0x2f, 0xbf, 0x7c, 0x33, 0xd0,
	0xd4, 0xa4, 0xb5, 0x90, 0x89, 0xe6, 0x2a, 0x95, 0x46, 0xb2, 0xbd, 0x50, 0xc6, 0x3c, 0x12, 0xe6,
	0x36, 0x0b, 0xf8, 0x55, 0x2a, 0xa3, 0x14, 0xe3, 0x58, 0x24, 0xd1, 0x25, 0x06, 0xd3, 0x81, 0x09,
	0x06, 0x23, 0x0c, 0x43, 0x99, 0x25, 0x46, 0x73, 0x54, 0xc2, 0xdd, 0x8a, 0xa4, 0x8c, 0x26, 0xe4,
	0xa3, 0x12, 0x3e, 0x26, 0x89, 0x34, 0x68, 0x16, 0x66, 0xee, 0xa6, 0xed, 0xe6, 0x55, 0x90, 0xdd,
	0xf8, 0x14, 0x2b, 0xf3, 0x38, 0x6b, 0x7a, 0x3d, 0x58, 0x1b, 0xce, 0xbe, 0xcd, 0xb6, 0x01, 0x2c,
	0xc6, 0x48, 0x8c, 0xbb, 0xce, 0x8e, 0xd3, 0xab, 0x0f, 0xea, 0x56, 0xb9, 0x18, 0x7b, 0x1b, 0xd0,
	0x3a, 0x27, 0x63, 0x87, 0x07, 0x74, 0x97, 0x91, 0x36, 0xde, 0x19, 0xb4, 0xfb, 0x29, 0xa1, 0xa1,
	0xa2, 0xce, 0x18, 0xac, 0x26, 0x18, 0x93, 0x75, 0xc9, 0x9f, 0x99, 0x0b, 0x35, 0x85, 0x5a, 0x3f,
	0xc8, 0x74, 0xdc, 0xad, 0xe4, 0xfa, 0x77, 0xed, 0x75, 0xa0, 0x7d, 0x42, 0x13, 0xfa, 0xed, 0xb3,
	0xff, 0xbe, 0x02, 0x4d, 0x2b, 0x0d, 0x29, 0xbd, 0x17, 0x21, 0xb1, 0x57, 0x07, 0x60, 0x01, 0xc2,
	0x8e, 0x78, 0xa9, 0xac, 0xf8, 0x9f, 0x7f, 0x70, 0x0f, 0x4b, 0x3a, 0xd8, 0xd7, 0xbd, 0xd6, 0xf3,
	0xc7, 0xe7, 0x4b, 0x65, 0x9d, 0xd5, 0xfd, 0xf9, 0xf6, 0xd8, 0x9b, 0x03, 0x8d, 0x42, 0x1e, 0xac,
	0x5f, 0xd2, 0x7c, 0x59, 0x9a, 0xff, 0x25, 0xf4, 0x7e, 0x10, 0x3e, 0x41, 0xa3, 0x10, 0x74, 0x69,
	0xc0, 0x65, 0x6b, 0x72, 0x3b, 0x7c, 0x76, 0x63, 0x7c, 0x7e, 0x63, 0xfc, 0x74, 0x7a, 0x63, 0x73,
	0x80, 0xdd, 0x05, 0xc0, 0x71, 0xed, 0xba, 0x8a, 0x4a, 0x8c, 0x54, 0x10, 0x54, 0xf3, 0xe1, 0x83,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x2e, 0x9b, 0xad, 0xfd, 0x02, 0x00, 0x00,
}
