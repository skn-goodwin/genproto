// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package product_manager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProductManagerServiceClient is the client API for ProductManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductManagerServiceClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type productManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductManagerServiceClient(cc grpc.ClientConnInterface) ProductManagerServiceClient {
	return &productManagerServiceClient{cc}
}

func (c *productManagerServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/atlant.product_manager.v1.ProductManagerService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagerServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/atlant.product_manager.v1.ProductManagerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductManagerServiceServer is the server API for ProductManagerService service.
// All implementations should embed UnimplementedProductManagerServiceServer
// for forward compatibility
type ProductManagerServiceServer interface {
	Fetch(context.Context, *FetchRequest) (*emptypb.Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedProductManagerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductManagerServiceServer struct {
}

func (UnimplementedProductManagerServiceServer) Fetch(context.Context, *FetchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedProductManagerServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeProductManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductManagerServiceServer will
// result in compilation errors.
type UnsafeProductManagerServiceServer interface {
	mustEmbedUnimplementedProductManagerServiceServer()
}

func RegisterProductManagerServiceServer(s *grpc.Server, srv ProductManagerServiceServer) {
	s.RegisterService(&_ProductManagerService_serviceDesc, srv)
}

func _ProductManagerService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagerServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlant.product_manager.v1.ProductManagerService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagerServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlant.product_manager.v1.ProductManagerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagerServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atlant.product_manager.v1.ProductManagerService",
	HandlerType: (*ProductManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _ProductManagerService_Fetch_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProductManagerService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product-manager/v1/product_manager.proto",
}
