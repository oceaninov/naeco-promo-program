// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProgramServiceClient is the client API for ProgramService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgramServiceClient interface {
	AddProgram(ctx context.Context, in *AddProgramReq, opts ...grpc.CallOption) (*ProgramRes, error)
	EditProgram(ctx context.Context, in *EditProgramReq, opts ...grpc.CallOption) (*ProgramRes, error)
	DeleteProgram(ctx context.Context, in *DeleteProgramReq, opts ...grpc.CallOption) (*ProgramRes, error)
	GetProgramByTopicID(ctx context.Context, in *GetProgramReq, opts ...grpc.CallOption) (*Programs, error)
	GetProgramDetail(ctx context.Context, in *GetProgramDetailReq, opts ...grpc.CallOption) (*Program, error)
	ProgramChangeStatus(ctx context.Context, in *ProgramStatus, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProgram(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Programs, error)
	ChangeStatusProgram(ctx context.Context, in *ProgramStatus, opts ...grpc.CallOption) (*Response, error)
	AddProgramBlacklistsBulk(ctx context.Context, in *Blacklisting, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteProgramBlacklistsBulk(ctx context.Context, in *Blacklisting, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProgramBlacklists(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*Blacklists, error)
}

type programServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProgramServiceClient(cc grpc.ClientConnInterface) ProgramServiceClient {
	return &programServiceClient{cc}
}

func (c *programServiceClient) AddProgram(ctx context.Context, in *AddProgramReq, opts ...grpc.CallOption) (*ProgramRes, error) {
	out := new(ProgramRes)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/AddProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) EditProgram(ctx context.Context, in *EditProgramReq, opts ...grpc.CallOption) (*ProgramRes, error) {
	out := new(ProgramRes)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/EditProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) DeleteProgram(ctx context.Context, in *DeleteProgramReq, opts ...grpc.CallOption) (*ProgramRes, error) {
	out := new(ProgramRes)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/DeleteProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) GetProgramByTopicID(ctx context.Context, in *GetProgramReq, opts ...grpc.CallOption) (*Programs, error) {
	out := new(Programs)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/GetProgramByTopicID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) GetProgramDetail(ctx context.Context, in *GetProgramDetailReq, opts ...grpc.CallOption) (*Program, error) {
	out := new(Program)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/GetProgramDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) ProgramChangeStatus(ctx context.Context, in *ProgramStatus, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/ProgramChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) GetProgram(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Programs, error) {
	out := new(Programs)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/GetProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) ChangeStatusProgram(ctx context.Context, in *ProgramStatus, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/ChangeStatusProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) AddProgramBlacklistsBulk(ctx context.Context, in *Blacklisting, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/AddProgramBlacklistsBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) DeleteProgramBlacklistsBulk(ctx context.Context, in *Blacklisting, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/DeleteProgramBlacklistsBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programServiceClient) GetProgramBlacklists(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*Blacklists, error) {
	out := new(Blacklists)
	err := c.cc.Invoke(ctx, "/api.v1.ProgramService/GetProgramBlacklists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProgramServiceServer is the server API for ProgramService service.
// All implementations should embed UnimplementedProgramServiceServer
// for forward compatibility
type ProgramServiceServer interface {
	AddProgram(context.Context, *AddProgramReq) (*ProgramRes, error)
	EditProgram(context.Context, *EditProgramReq) (*ProgramRes, error)
	DeleteProgram(context.Context, *DeleteProgramReq) (*ProgramRes, error)
	GetProgramByTopicID(context.Context, *GetProgramReq) (*Programs, error)
	GetProgramDetail(context.Context, *GetProgramDetailReq) (*Program, error)
	ProgramChangeStatus(context.Context, *ProgramStatus) (*emptypb.Empty, error)
	GetProgram(context.Context, *emptypb.Empty) (*Programs, error)
	ChangeStatusProgram(context.Context, *ProgramStatus) (*Response, error)
	AddProgramBlacklistsBulk(context.Context, *Blacklisting) (*emptypb.Empty, error)
	DeleteProgramBlacklistsBulk(context.Context, *Blacklisting) (*emptypb.Empty, error)
	GetProgramBlacklists(context.Context, *GetBlacklistReq) (*Blacklists, error)
}

// UnimplementedProgramServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProgramServiceServer struct {
}

func (UnimplementedProgramServiceServer) AddProgram(context.Context, *AddProgramReq) (*ProgramRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProgram not implemented")
}
func (UnimplementedProgramServiceServer) EditProgram(context.Context, *EditProgramReq) (*ProgramRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProgram not implemented")
}
func (UnimplementedProgramServiceServer) DeleteProgram(context.Context, *DeleteProgramReq) (*ProgramRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProgram not implemented")
}
func (UnimplementedProgramServiceServer) GetProgramByTopicID(context.Context, *GetProgramReq) (*Programs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgramByTopicID not implemented")
}
func (UnimplementedProgramServiceServer) GetProgramDetail(context.Context, *GetProgramDetailReq) (*Program, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgramDetail not implemented")
}
func (UnimplementedProgramServiceServer) ProgramChangeStatus(context.Context, *ProgramStatus) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProgramChangeStatus not implemented")
}
func (UnimplementedProgramServiceServer) GetProgram(context.Context, *emptypb.Empty) (*Programs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgram not implemented")
}
func (UnimplementedProgramServiceServer) ChangeStatusProgram(context.Context, *ProgramStatus) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatusProgram not implemented")
}
func (UnimplementedProgramServiceServer) AddProgramBlacklistsBulk(context.Context, *Blacklisting) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProgramBlacklistsBulk not implemented")
}
func (UnimplementedProgramServiceServer) DeleteProgramBlacklistsBulk(context.Context, *Blacklisting) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProgramBlacklistsBulk not implemented")
}
func (UnimplementedProgramServiceServer) GetProgramBlacklists(context.Context, *GetBlacklistReq) (*Blacklists, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgramBlacklists not implemented")
}

// UnsafeProgramServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgramServiceServer will
// result in compilation errors.
type UnsafeProgramServiceServer interface {
	mustEmbedUnimplementedProgramServiceServer()
}

func RegisterProgramServiceServer(s grpc.ServiceRegistrar, srv ProgramServiceServer) {
	s.RegisterService(&ProgramService_ServiceDesc, srv)
}

func _ProgramService_AddProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProgramReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).AddProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/AddProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).AddProgram(ctx, req.(*AddProgramReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_EditProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProgramReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).EditProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/EditProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).EditProgram(ctx, req.(*EditProgramReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_DeleteProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProgramReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).DeleteProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/DeleteProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).DeleteProgram(ctx, req.(*DeleteProgramReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_GetProgramByTopicID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProgramReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).GetProgramByTopicID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/GetProgramByTopicID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).GetProgramByTopicID(ctx, req.(*GetProgramReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_GetProgramDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProgramDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).GetProgramDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/GetProgramDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).GetProgramDetail(ctx, req.(*GetProgramDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_ProgramChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgramStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).ProgramChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/ProgramChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).ProgramChangeStatus(ctx, req.(*ProgramStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_GetProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).GetProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/GetProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).GetProgram(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_ChangeStatusProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgramStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).ChangeStatusProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/ChangeStatusProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).ChangeStatusProgram(ctx, req.(*ProgramStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_AddProgramBlacklistsBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blacklisting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).AddProgramBlacklistsBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/AddProgramBlacklistsBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).AddProgramBlacklistsBulk(ctx, req.(*Blacklisting))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_DeleteProgramBlacklistsBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blacklisting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).DeleteProgramBlacklistsBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/DeleteProgramBlacklistsBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).DeleteProgramBlacklistsBulk(ctx, req.(*Blacklisting))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramService_GetProgramBlacklists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramServiceServer).GetProgramBlacklists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProgramService/GetProgramBlacklists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramServiceServer).GetProgramBlacklists(ctx, req.(*GetBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProgramService_ServiceDesc is the grpc.ServiceDesc for ProgramService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgramService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.ProgramService",
	HandlerType: (*ProgramServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProgram",
			Handler:    _ProgramService_AddProgram_Handler,
		},
		{
			MethodName: "EditProgram",
			Handler:    _ProgramService_EditProgram_Handler,
		},
		{
			MethodName: "DeleteProgram",
			Handler:    _ProgramService_DeleteProgram_Handler,
		},
		{
			MethodName: "GetProgramByTopicID",
			Handler:    _ProgramService_GetProgramByTopicID_Handler,
		},
		{
			MethodName: "GetProgramDetail",
			Handler:    _ProgramService_GetProgramDetail_Handler,
		},
		{
			MethodName: "ProgramChangeStatus",
			Handler:    _ProgramService_ProgramChangeStatus_Handler,
		},
		{
			MethodName: "GetProgram",
			Handler:    _ProgramService_GetProgram_Handler,
		},
		{
			MethodName: "ChangeStatusProgram",
			Handler:    _ProgramService_ChangeStatusProgram_Handler,
		},
		{
			MethodName: "AddProgramBlacklistsBulk",
			Handler:    _ProgramService_AddProgramBlacklistsBulk_Handler,
		},
		{
			MethodName: "DeleteProgramBlacklistsBulk",
			Handler:    _ProgramService_DeleteProgramBlacklistsBulk_Handler,
		},
		{
			MethodName: "GetProgramBlacklists",
			Handler:    _ProgramService_GetProgramBlacklists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "program.proto",
}
