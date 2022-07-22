// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: proto/student.proto

package pb

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

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*Student, error)
	GetStudentById(ctx context.Context, in *GetStudentByIdRequest, opts ...grpc.CallOption) (*Student, error)
	GetAllStudent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStudent, error)
	UpdateStudentAddress(ctx context.Context, in *UpdateStudentAddressRequest, opts ...grpc.CallOption) (*Student, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/student.StudentService/CreateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentById(ctx context.Context, in *GetStudentByIdRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/student.StudentService/GetStudentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetAllStudent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStudent, error) {
	out := new(ListStudent)
	err := c.cc.Invoke(ctx, "/student.StudentService/GetAllStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudentAddress(ctx context.Context, in *UpdateStudentAddressRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/student.StudentService/UpdateStudentAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	CreateStudent(context.Context, *CreateStudentRequest) (*Student, error)
	GetStudentById(context.Context, *GetStudentByIdRequest) (*Student, error)
	GetAllStudent(context.Context, *emptypb.Empty) (*ListStudent, error)
	UpdateStudentAddress(context.Context, *UpdateStudentAddressRequest) (*Student, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) CreateStudent(context.Context, *CreateStudentRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentById(context.Context, *GetStudentByIdRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentById not implemented")
}
func (UnimplementedStudentServiceServer) GetAllStudent(context.Context, *emptypb.Empty) (*ListStudent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStudent not implemented")
}
func (UnimplementedStudentServiceServer) UpdateStudentAddress(context.Context, *UpdateStudentAddressRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudentAddress not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/CreateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateStudent(ctx, req.(*CreateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/GetStudentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentById(ctx, req.(*GetStudentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetAllStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetAllStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/GetAllStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetAllStudent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateStudentAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateStudentAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/UpdateStudentAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateStudentAddress(ctx, req.(*UpdateStudentAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudent",
			Handler:    _StudentService_CreateStudent_Handler,
		},
		{
			MethodName: "GetStudentById",
			Handler:    _StudentService_GetStudentById_Handler,
		},
		{
			MethodName: "GetAllStudent",
			Handler:    _StudentService_GetAllStudent_Handler,
		},
		{
			MethodName: "UpdateStudentAddress",
			Handler:    _StudentService_UpdateStudentAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/student.proto",
}
