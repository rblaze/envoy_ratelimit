package rlspb

import (
	"context"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	RateLimitService_ShouldRateLimit_FullMethodName  = "/envoy.service.ratelimit.v3.RateLimitService/ShouldRateLimit"
	RateLimitService_CheckRateLimit_FullMethodName   = "/envoy.service.ratelimit.v3.RateLimitService/CheckRateLimit"
	RateLimitService_ConsumeRateLimit_FullMethodName = "/envoy.service.ratelimit.v3.RateLimitService/ConsumeRateLimit"
)

// RateLimitServiceClient is the client API for RateLimitService service,
// extending go-control-plane's RateLimitServiceClient with CheckRateLimit and ConsumeRateLimit.
type RateLimitServiceClient interface {
	pb.RateLimitServiceClient
	// Check whether rate limiting would take place without consuming quota.
	CheckRateLimit(ctx context.Context, in *pb.RateLimitRequest, opts ...grpc.CallOption) (*pb.RateLimitResponse, error)
	// Consume rate limit quota without performing a limit check.
	ConsumeRateLimit(ctx context.Context, in *pb.RateLimitRequest, opts ...grpc.CallOption) (*pb.RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateLimitServiceClient(cc grpc.ClientConnInterface) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *pb.RateLimitRequest, opts ...grpc.CallOption) (*pb.RateLimitResponse, error) {
	out := new(pb.RateLimitResponse)
	err := c.cc.Invoke(ctx, RateLimitService_ShouldRateLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimitServiceClient) CheckRateLimit(ctx context.Context, in *pb.RateLimitRequest, opts ...grpc.CallOption) (*pb.RateLimitResponse, error) {
	out := new(pb.RateLimitResponse)
	err := c.cc.Invoke(ctx, RateLimitService_CheckRateLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateLimitServiceClient) ConsumeRateLimit(ctx context.Context, in *pb.RateLimitRequest, opts ...grpc.CallOption) (*pb.RateLimitResponse, error) {
	out := new(pb.RateLimitResponse)
	err := c.cc.Invoke(ctx, RateLimitService_ConsumeRateLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service,
// extending go-control-plane's RateLimitServiceServer with CheckRateLimit and ConsumeRateLimit.
type RateLimitServiceServer interface {
	pb.RateLimitServiceServer
	// Check whether rate limiting would take place without consuming quota.
	CheckRateLimit(context.Context, *pb.RateLimitRequest) (*pb.RateLimitResponse, error)
	// Consume rate limit quota without performing a limit check.
	ConsumeRateLimit(context.Context, *pb.RateLimitRequest) (*pb.RateLimitResponse, error)
}

// UnimplementedRateLimitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitServiceServer struct{}

func (UnimplementedRateLimitServiceServer) ShouldRateLimit(context.Context, *pb.RateLimitRequest) (*pb.RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShouldRateLimit not implemented")
}

func (UnimplementedRateLimitServiceServer) CheckRateLimit(context.Context, *pb.RateLimitRequest) (*pb.RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRateLimit not implemented")
}

func (UnimplementedRateLimitServiceServer) ConsumeRateLimit(context.Context, *pb.RateLimitRequest) (*pb.RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeRateLimit not implemented")
}

func RegisterRateLimitServiceServer(s grpc.ServiceRegistrar, srv RateLimitServiceServer) {
	s.RegisterService(&RateLimitService_ServiceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RateLimitService_ShouldRateLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*pb.RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimitService_CheckRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).CheckRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RateLimitService_CheckRateLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).CheckRateLimit(ctx, req.(*pb.RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateLimitService_ConsumeRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ConsumeRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RateLimitService_ConsumeRateLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ConsumeRateLimit(ctx, req.(*pb.RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RateLimitService_ServiceDesc is the grpc.ServiceDesc for RateLimitService service.
var RateLimitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v3.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
		{
			MethodName: "CheckRateLimit",
			Handler:    _RateLimitService_CheckRateLimit_Handler,
		},
		{
			MethodName: "ConsumeRateLimit",
			Handler:    _RateLimitService_ConsumeRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v3/rls.proto",
}
