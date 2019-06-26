// Code generated by protoc-gen-go. DO NOT EDIT.
// source: live_entity.proto

package live_entity

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LiveEntity struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiveEntity) Reset()         { *m = LiveEntity{} }
func (m *LiveEntity) String() string { return proto.CompactTextString(m) }
func (*LiveEntity) ProtoMessage()    {}
func (*LiveEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd09aacd1224108, []int{0}
}

func (m *LiveEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiveEntity.Unmarshal(m, b)
}
func (m *LiveEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiveEntity.Marshal(b, m, deterministic)
}
func (m *LiveEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiveEntity.Merge(m, src)
}
func (m *LiveEntity) XXX_Size() int {
	return xxx_messageInfo_LiveEntity.Size(m)
}
func (m *LiveEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_LiveEntity.DiscardUnknown(m)
}

var xxx_messageInfo_LiveEntity proto.InternalMessageInfo

func (m *LiveEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LiveEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LiveEntity) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RequestGetLiveEntity struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestGetLiveEntity) Reset()         { *m = RequestGetLiveEntity{} }
func (m *RequestGetLiveEntity) String() string { return proto.CompactTextString(m) }
func (*RequestGetLiveEntity) ProtoMessage()    {}
func (*RequestGetLiveEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd09aacd1224108, []int{1}
}

func (m *RequestGetLiveEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestGetLiveEntity.Unmarshal(m, b)
}
func (m *RequestGetLiveEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestGetLiveEntity.Marshal(b, m, deterministic)
}
func (m *RequestGetLiveEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestGetLiveEntity.Merge(m, src)
}
func (m *RequestGetLiveEntity) XXX_Size() int {
	return xxx_messageInfo_RequestGetLiveEntity.Size(m)
}
func (m *RequestGetLiveEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestGetLiveEntity.DiscardUnknown(m)
}

var xxx_messageInfo_RequestGetLiveEntity proto.InternalMessageInfo

func (m *RequestGetLiveEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RequestDeleteLiveEntity struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestDeleteLiveEntity) Reset()         { *m = RequestDeleteLiveEntity{} }
func (m *RequestDeleteLiveEntity) String() string { return proto.CompactTextString(m) }
func (*RequestDeleteLiveEntity) ProtoMessage()    {}
func (*RequestDeleteLiveEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd09aacd1224108, []int{2}
}

func (m *RequestDeleteLiveEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestDeleteLiveEntity.Unmarshal(m, b)
}
func (m *RequestDeleteLiveEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestDeleteLiveEntity.Marshal(b, m, deterministic)
}
func (m *RequestDeleteLiveEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestDeleteLiveEntity.Merge(m, src)
}
func (m *RequestDeleteLiveEntity) XXX_Size() int {
	return xxx_messageInfo_RequestDeleteLiveEntity.Size(m)
}
func (m *RequestDeleteLiveEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestDeleteLiveEntity.DiscardUnknown(m)
}

var xxx_messageInfo_RequestDeleteLiveEntity proto.InternalMessageInfo

func (m *RequestDeleteLiveEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RequestCreateLiveEntity struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestCreateLiveEntity) Reset()         { *m = RequestCreateLiveEntity{} }
func (m *RequestCreateLiveEntity) String() string { return proto.CompactTextString(m) }
func (*RequestCreateLiveEntity) ProtoMessage()    {}
func (*RequestCreateLiveEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd09aacd1224108, []int{3}
}

func (m *RequestCreateLiveEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestCreateLiveEntity.Unmarshal(m, b)
}
func (m *RequestCreateLiveEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestCreateLiveEntity.Marshal(b, m, deterministic)
}
func (m *RequestCreateLiveEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCreateLiveEntity.Merge(m, src)
}
func (m *RequestCreateLiveEntity) XXX_Size() int {
	return xxx_messageInfo_RequestCreateLiveEntity.Size(m)
}
func (m *RequestCreateLiveEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCreateLiveEntity.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCreateLiveEntity proto.InternalMessageInfo

func (m *RequestCreateLiveEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestCreateLiveEntity) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RequestUpdateLiveEntity struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUpdateLiveEntity) Reset()         { *m = RequestUpdateLiveEntity{} }
func (m *RequestUpdateLiveEntity) String() string { return proto.CompactTextString(m) }
func (*RequestUpdateLiveEntity) ProtoMessage()    {}
func (*RequestUpdateLiveEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd09aacd1224108, []int{4}
}

func (m *RequestUpdateLiveEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdateLiveEntity.Unmarshal(m, b)
}
func (m *RequestUpdateLiveEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdateLiveEntity.Marshal(b, m, deterministic)
}
func (m *RequestUpdateLiveEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdateLiveEntity.Merge(m, src)
}
func (m *RequestUpdateLiveEntity) XXX_Size() int {
	return xxx_messageInfo_RequestUpdateLiveEntity.Size(m)
}
func (m *RequestUpdateLiveEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdateLiveEntity.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdateLiveEntity proto.InternalMessageInfo

func (m *RequestUpdateLiveEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RequestUpdateLiveEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestUpdateLiveEntity) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ResponseLiveEntity struct {
	Entity               *LiveEntity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResponseLiveEntity) Reset()         { *m = ResponseLiveEntity{} }
func (m *ResponseLiveEntity) String() string { return proto.CompactTextString(m) }
func (*ResponseLiveEntity) ProtoMessage()    {}
func (*ResponseLiveEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd09aacd1224108, []int{5}
}

func (m *ResponseLiveEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseLiveEntity.Unmarshal(m, b)
}
func (m *ResponseLiveEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseLiveEntity.Marshal(b, m, deterministic)
}
func (m *ResponseLiveEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseLiveEntity.Merge(m, src)
}
func (m *ResponseLiveEntity) XXX_Size() int {
	return xxx_messageInfo_ResponseLiveEntity.Size(m)
}
func (m *ResponseLiveEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseLiveEntity.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseLiveEntity proto.InternalMessageInfo

func (m *ResponseLiveEntity) GetEntity() *LiveEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*LiveEntity)(nil), "live_entity.LiveEntity")
	proto.RegisterType((*RequestGetLiveEntity)(nil), "live_entity.requestGetLiveEntity")
	proto.RegisterType((*RequestDeleteLiveEntity)(nil), "live_entity.requestDeleteLiveEntity")
	proto.RegisterType((*RequestCreateLiveEntity)(nil), "live_entity.requestCreateLiveEntity")
	proto.RegisterType((*RequestUpdateLiveEntity)(nil), "live_entity.requestUpdateLiveEntity")
	proto.RegisterType((*ResponseLiveEntity)(nil), "live_entity.responseLiveEntity")
}

func init() { proto.RegisterFile("live_entity.proto", fileDescriptor_8bd09aacd1224108) }

var fileDescriptor_8bd09aacd1224108 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x49, 0xfa, 0xa3, 0xf0, 0x9b, 0x62, 0xa9, 0x5b, 0xb5, 0xb1, 0x8a, 0x7f, 0x82, 0x88,
	0xf6, 0xd0, 0x85, 0x7a, 0xf3, 0xaa, 0x45, 0x0f, 0x9e, 0x2a, 0x3d, 0xcb, 0xda, 0x0c, 0x65, 0xa1,
	0xee, 0xc6, 0xec, 0x36, 0x50, 0xc4, 0x8b, 0xaf, 0xe0, 0xa3, 0xf9, 0x0a, 0x1e, 0x7c, 0x0c, 0xc9,
	0x36, 0xd6, 0x64, 0x59, 0x5b, 0xbd, 0x25, 0xc3, 0x27, 0xf3, 0x99, 0xfd, 0xce, 0x06, 0xd6, 0x27,
	0x3c, 0xc5, 0x3b, 0x14, 0x9a, 0xeb, 0x59, 0x37, 0x4e, 0xa4, 0x96, 0xa4, 0x56, 0x28, 0xb5, 0x77,
	0xc7, 0x52, 0x8e, 0x27, 0x48, 0x59, 0xcc, 0x29, 0x13, 0x42, 0x6a, 0xa6, 0xb9, 0x14, 0x6a, 0x8e,
	0x86, 0xd7, 0x00, 0x37, 0x3c, 0xc5, 0xbe, 0x61, 0x49, 0x1d, 0x7c, 0x1e, 0x05, 0xde, 0x81, 0x77,
	0xf2, 0x7f, 0xe0, 0xf3, 0x88, 0x10, 0xf8, 0x27, 0xd8, 0x03, 0x06, 0xbe, 0xa9, 0x98, 0x67, 0xb2,
	0x05, 0x55, 0xa5, 0x99, 0x9e, 0xaa, 0xa0, 0x62, 0xaa, 0xf9, 0x5b, 0x78, 0x0c, 0x1b, 0x09, 0x3e,
	0x4e, 0x51, 0xe9, 0x2b, 0xd4, 0x3f, 0xf7, 0x0c, 0x4f, 0xa1, 0x95, 0x73, 0x97, 0x38, 0x41, 0x8d,
	0x4b, 0xd0, 0xfe, 0x02, 0xbd, 0x48, 0x90, 0x95, 0xd0, 0xaf, 0xc9, 0x3c, 0xe7, 0x64, 0x7e, 0x69,
	0xb2, 0xe1, 0xa2, 0xcd, 0x30, 0x8e, 0xd8, 0x32, 0xe3, 0x9f, 0x0e, 0xdc, 0x07, 0x92, 0xa0, 0x8a,
	0xa5, 0x50, 0xc5, 0x8e, 0x14, 0xaa, 0xf3, 0xe0, 0x4d, 0xd7, 0x5a, 0xaf, 0xd5, 0x2d, 0xee, 0xe7,
	0x1b, 0x1c, 0xe4, 0x58, 0xef, 0xa3, 0x02, 0xb5, 0xac, 0x7c, 0x8b, 0x49, 0xca, 0x47, 0x48, 0x62,
	0x58, 0x1b, 0x97, 0x02, 0x3c, 0x2c, 0x75, 0x70, 0x65, 0xdc, 0xde, 0xb7, 0x10, 0x7b, 0xaa, 0x70,
	0xe7, 0xe5, 0xed, 0xfd, 0xd5, 0xdf, 0x24, 0x4d, 0x6a, 0x18, 0x8e, 0x8a, 0x66, 0x5f, 0xd0, 0x27,
	0x1e, 0x3d, 0x13, 0x0d, 0x8d, 0x91, 0x9d, 0xef, 0x91, 0x4b, 0x6a, 0x6f, 0x61, 0xb5, 0x77, 0xdb,
	0x78, 0x9b, 0x61, 0xbd, 0xec, 0x3d, 0xf7, 0x3a, 0x64, 0x06, 0x8d, 0xa9, 0xbd, 0x0e, 0xa7, 0xd5,
	0x5e, 0xda, 0x6a, 0xeb, 0x9e, 0xb1, 0x06, 0x6d, 0xd7, 0x69, 0x33, 0x75, 0x0a, 0x8d, 0xc8, 0xbe,
	0x7b, 0x4e, 0xb5, 0x7d, 0x43, 0x7f, 0x1d, 0x74, 0xc7, 0xa5, 0xbe, 0xaf, 0x9a, 0x7f, 0xee, 0xec,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xd0, 0x13, 0xda, 0xb3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LiveServiceClient is the client API for LiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LiveServiceClient interface {
	GetLiveEntity(ctx context.Context, in *RequestGetLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error)
	CreateLiveEntity(ctx context.Context, in *RequestCreateLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error)
	UpdateLiveEntity(ctx context.Context, in *RequestUpdateLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error)
	DeleteLiveEntity(ctx context.Context, in *RequestDeleteLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error)
}

type liveServiceClient struct {
	cc *grpc.ClientConn
}

func NewLiveServiceClient(cc *grpc.ClientConn) LiveServiceClient {
	return &liveServiceClient{cc}
}

func (c *liveServiceClient) GetLiveEntity(ctx context.Context, in *RequestGetLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error) {
	out := new(ResponseLiveEntity)
	err := c.cc.Invoke(ctx, "/live_entity.LiveService/getLiveEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) CreateLiveEntity(ctx context.Context, in *RequestCreateLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error) {
	out := new(ResponseLiveEntity)
	err := c.cc.Invoke(ctx, "/live_entity.LiveService/createLiveEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) UpdateLiveEntity(ctx context.Context, in *RequestUpdateLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error) {
	out := new(ResponseLiveEntity)
	err := c.cc.Invoke(ctx, "/live_entity.LiveService/updateLiveEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) DeleteLiveEntity(ctx context.Context, in *RequestDeleteLiveEntity, opts ...grpc.CallOption) (*ResponseLiveEntity, error) {
	out := new(ResponseLiveEntity)
	err := c.cc.Invoke(ctx, "/live_entity.LiveService/deleteLiveEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveServiceServer is the server API for LiveService service.
type LiveServiceServer interface {
	GetLiveEntity(context.Context, *RequestGetLiveEntity) (*ResponseLiveEntity, error)
	CreateLiveEntity(context.Context, *RequestCreateLiveEntity) (*ResponseLiveEntity, error)
	UpdateLiveEntity(context.Context, *RequestUpdateLiveEntity) (*ResponseLiveEntity, error)
	DeleteLiveEntity(context.Context, *RequestDeleteLiveEntity) (*ResponseLiveEntity, error)
}

// UnimplementedLiveServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLiveServiceServer struct {
}

func (*UnimplementedLiveServiceServer) GetLiveEntity(ctx context.Context, req *RequestGetLiveEntity) (*ResponseLiveEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveEntity not implemented")
}
func (*UnimplementedLiveServiceServer) CreateLiveEntity(ctx context.Context, req *RequestCreateLiveEntity) (*ResponseLiveEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLiveEntity not implemented")
}
func (*UnimplementedLiveServiceServer) UpdateLiveEntity(ctx context.Context, req *RequestUpdateLiveEntity) (*ResponseLiveEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLiveEntity not implemented")
}
func (*UnimplementedLiveServiceServer) DeleteLiveEntity(ctx context.Context, req *RequestDeleteLiveEntity) (*ResponseLiveEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLiveEntity not implemented")
}

func RegisterLiveServiceServer(s *grpc.Server, srv LiveServiceServer) {
	s.RegisterService(&_LiveService_serviceDesc, srv)
}

func _LiveService_GetLiveEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetLiveEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).GetLiveEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/live_entity.LiveService/GetLiveEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).GetLiveEntity(ctx, req.(*RequestGetLiveEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_CreateLiveEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateLiveEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).CreateLiveEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/live_entity.LiveService/CreateLiveEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).CreateLiveEntity(ctx, req.(*RequestCreateLiveEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_UpdateLiveEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateLiveEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).UpdateLiveEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/live_entity.LiveService/UpdateLiveEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).UpdateLiveEntity(ctx, req.(*RequestUpdateLiveEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_DeleteLiveEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteLiveEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).DeleteLiveEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/live_entity.LiveService/DeleteLiveEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).DeleteLiveEntity(ctx, req.(*RequestDeleteLiveEntity))
	}
	return interceptor(ctx, in, info, handler)
}

var _LiveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "live_entity.LiveService",
	HandlerType: (*LiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getLiveEntity",
			Handler:    _LiveService_GetLiveEntity_Handler,
		},
		{
			MethodName: "createLiveEntity",
			Handler:    _LiveService_CreateLiveEntity_Handler,
		},
		{
			MethodName: "updateLiveEntity",
			Handler:    _LiveService_UpdateLiveEntity_Handler,
		},
		{
			MethodName: "deleteLiveEntity",
			Handler:    _LiveService_DeleteLiveEntity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "live_entity.proto",
}
