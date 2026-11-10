package rlspb_test

import (
	"context"
	"net"
	"testing"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/envoyproxy/ratelimit/src/rlspb"
)

func TestUnimplementedRateLimitServiceServer(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer ln.Close()

	srv := grpc.NewServer()
	rlspb.RegisterRateLimitServiceServer(srv, rlspb.UnimplementedRateLimitServiceServer{})
	go srv.Serve(ln)
	defer srv.Stop()

	conn, err := grpc.NewClient(ln.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	client := rlspb.NewRateLimitServiceClient(conn)
	req := &pb.RateLimitRequest{Domain: "d"}

	for _, call := range []func(context.Context, *pb.RateLimitRequest, ...grpc.CallOption) (*pb.RateLimitResponse, error){
		client.ShouldRateLimit,
		client.CheckRateLimit,
		client.ConsumeRateLimit,
	} {
		_, err := call(context.Background(), req)
		assert.Equal(t, codes.Unimplemented, status.Code(err))
	}
}
