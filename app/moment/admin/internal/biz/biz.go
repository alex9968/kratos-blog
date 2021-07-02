package biz

import (
	"context"

	ctV1 "github.com/go-kratos/beer-shop/api/catalog/service/v1"
	usV1 "github.com/go-kratos/beer-shop/api/user/service/v1"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewUserUseCase,
	NewDiscovery,
	NewUserServiceClient,
	NewCatalogServiceClient,
)

func NewDiscovery() registry.Discovery {
	cli, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}

func NewUserServiceClient(r registry.Discovery) usV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///default/beer.user.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return usV1.NewUserClient(conn)
}


func NewCatalogServiceClient(r registry.Discovery) ctV1.CatalogClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///default/beer.catalog.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return ctV1.NewCatalogClient(conn)
}

