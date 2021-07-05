package data

import (
	"context"
	"log"
	"os"
	"time"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"kratos-blog/app/moment/interface/internal/conf"

	userv1 "kratos-blog/api/user/service/v1"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewUserServiceClient,
	NewUserRepo,
	NewMomentRepo,
	NewTagRepo,
	NewCardRepo,
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	log *klog.Helper
	uc  userv1.UserClient
}

var dbLogger = logger.New(
	log.New(os.Stdout, "\r\n", 0), // io writer
	logger.Config{
		SlowThreshold: time.Second,   // Slow SQL threshold
		LogLevel:      logger.Silent, // Log level
		Colorful:      false,         // Disable color
	},
)

// NewData .
func NewData(conf *conf.Data, logger klog.Logger, uc userv1.UserClient) (*Data, func(), error) {
	log := klog.NewHelper(klog.With(logger, "module", "server-service/data"))

	//mysql init
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}

	if err := db.AutoMigrate(&Moment{}, &Tag{}); err != nil {
		panic(err)
	}

	//redis init
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		log: log,
		db:  db,
		uc:  uc,
	}
	return d, func() {
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

// NewDiscovery connect consul
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

// NewUserServiceClient connect user.service
func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///blog.user.service"),
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
