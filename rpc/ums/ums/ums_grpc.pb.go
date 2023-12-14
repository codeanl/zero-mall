// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ums.proto

package ums

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Ums_SaveOrUpdateMember_FullMethodName = "/ums.Ums/SaveOrUpdateMember"
	Ums_MemberDelete_FullMethodName       = "/ums.Ums/MemberDelete"
	Ums_MemberList_FullMethodName         = "/ums.Ums/MemberList"
)

// UmsClient is the client API for Ums service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UmsClient interface {
	// 添加｜｜更新会员
	SaveOrUpdateMember(ctx context.Context, in *SaveOrUpdateMemberReq, opts ...grpc.CallOption) (*SaveOrUpdateMemberResp, error)
	// 删除会员
	MemberDelete(ctx context.Context, in *MemberDeleteReq, opts ...grpc.CallOption) (*MemberDeleteResp, error)
	// 会员列表
	MemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error)
}

type umsClient struct {
	cc grpc.ClientConnInterface
}

func NewUmsClient(cc grpc.ClientConnInterface) UmsClient {
	return &umsClient{cc}
}

func (c *umsClient) SaveOrUpdateMember(ctx context.Context, in *SaveOrUpdateMemberReq, opts ...grpc.CallOption) (*SaveOrUpdateMemberResp, error) {
	out := new(SaveOrUpdateMemberResp)
	err := c.cc.Invoke(ctx, Ums_SaveOrUpdateMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) MemberDelete(ctx context.Context, in *MemberDeleteReq, opts ...grpc.CallOption) (*MemberDeleteResp, error) {
	out := new(MemberDeleteResp)
	err := c.cc.Invoke(ctx, Ums_MemberDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) MemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error) {
	out := new(MemberListResp)
	err := c.cc.Invoke(ctx, Ums_MemberList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UmsServer is the server API for Ums service.
// All implementations must embed UnimplementedUmsServer
// for forward compatibility
type UmsServer interface {
	// 添加｜｜更新会员
	SaveOrUpdateMember(context.Context, *SaveOrUpdateMemberReq) (*SaveOrUpdateMemberResp, error)
	// 删除会员
	MemberDelete(context.Context, *MemberDeleteReq) (*MemberDeleteResp, error)
	// 会员列表
	MemberList(context.Context, *MemberListReq) (*MemberListResp, error)
	mustEmbedUnimplementedUmsServer()
}

// UnimplementedUmsServer must be embedded to have forward compatible implementations.
type UnimplementedUmsServer struct {
}

func (UnimplementedUmsServer) SaveOrUpdateMember(context.Context, *SaveOrUpdateMemberReq) (*SaveOrUpdateMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveOrUpdateMember not implemented")
}
func (UnimplementedUmsServer) MemberDelete(context.Context, *MemberDeleteReq) (*MemberDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberDelete not implemented")
}
func (UnimplementedUmsServer) MemberList(context.Context, *MemberListReq) (*MemberListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberList not implemented")
}
func (UnimplementedUmsServer) mustEmbedUnimplementedUmsServer() {}

// UnsafeUmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UmsServer will
// result in compilation errors.
type UnsafeUmsServer interface {
	mustEmbedUnimplementedUmsServer()
}

func RegisterUmsServer(s grpc.ServiceRegistrar, srv UmsServer) {
	s.RegisterService(&Ums_ServiceDesc, srv)
}

func _Ums_SaveOrUpdateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveOrUpdateMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).SaveOrUpdateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ums_SaveOrUpdateMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).SaveOrUpdateMember(ctx, req.(*SaveOrUpdateMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_MemberDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).MemberDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ums_MemberDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).MemberDelete(ctx, req.(*MemberDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_MemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).MemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ums_MemberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).MemberList(ctx, req.(*MemberListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ums_ServiceDesc is the grpc.ServiceDesc for Ums service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ums_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ums.Ums",
	HandlerType: (*UmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveOrUpdateMember",
			Handler:    _Ums_SaveOrUpdateMember_Handler,
		},
		{
			MethodName: "MemberDelete",
			Handler:    _Ums_MemberDelete_Handler,
		},
		{
			MethodName: "MemberList",
			Handler:    _Ums_MemberList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ums.proto",
}