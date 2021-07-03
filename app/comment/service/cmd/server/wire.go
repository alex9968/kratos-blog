// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-blog/app/comment/service/internal/biz"
	"kratos-blog/app/comment/service/internal/conf"
	"kratos-blog/app/comment/service/internal/data"
	"kratos-blog/app/comment/service/internal/server"
	"kratos-blog/app/comment/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}