package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/gorm"

	"github.com/go-kratos/beer-shop/app/moment/interface/internal/conf"

	"context"

	userv1 "github.com/go-kratos/beer-shop/api/user/service/v1"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewUserServiceClient,
	NewUserRepo,
)

// Data .
type Data struct {
	db *gorm.DB
	log *log.Helper
	uc  userv1.UserClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	uc userv1.UserClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}

func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
			recovery.Recovery()),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}
