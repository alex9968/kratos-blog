package server

import (
	"github.com/go-kratos/beer-shop/api/user/service/v1"
	"github.com/go-kratos/beer-shop/app/user/service/internal/conf"
	"github.com/go-kratos/beer-shop/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, s *service.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				status.Server(),
				tracing.Server(),
				logging.Server(),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, s)
	return srv
}

