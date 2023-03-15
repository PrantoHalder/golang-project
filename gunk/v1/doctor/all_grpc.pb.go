// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package doctorpb

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

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	RegisterDoctorType(ctx context.Context, in *RegisterDoctorTypeRequest, opts ...grpc.CallOption) (*RegisterDoctorTypeResponse, error)
	DoctorLogin(ctx context.Context, in *DoctorLoginRequest, opts ...grpc.CallOption) (*DoctorLoginResponse, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) RegisterDoctorType(ctx context.Context, in *RegisterDoctorTypeRequest, opts ...grpc.CallOption) (*RegisterDoctorTypeResponse, error) {
	out := new(RegisterDoctorTypeResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/RegisterDoctorType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DoctorLogin(ctx context.Context, in *DoctorLoginRequest, opts ...grpc.CallOption) (*DoctorLoginResponse, error) {
	out := new(DoctorLoginResponse)
	err := c.cc.Invoke(ctx, "/doctorpb.DoctorService/DoctorLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	RegisterDoctorType(context.Context, *RegisterDoctorTypeRequest) (*RegisterDoctorTypeResponse, error)
	DoctorLogin(context.Context, *DoctorLoginRequest) (*DoctorLoginResponse, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) RegisterDoctorType(context.Context, *RegisterDoctorTypeRequest) (*RegisterDoctorTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDoctorType not implemented")
}
func (UnimplementedDoctorServiceServer) DoctorLogin(context.Context, *DoctorLoginRequest) (*DoctorLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoctorLogin not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_RegisterDoctorType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDoctorTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).RegisterDoctorType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/RegisterDoctorType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).RegisterDoctorType(ctx, req.(*RegisterDoctorTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DoctorLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DoctorLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctorpb.DoctorService/DoctorLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DoctorLogin(ctx, req.(*DoctorLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctorpb.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDoctorType",
			Handler:    _DoctorService_RegisterDoctorType_Handler,
		},
		{
			MethodName: "DoctorLogin",
			Handler:    _DoctorService_DoctorLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.go/gunk/v1/doctor/all.proto",
}
