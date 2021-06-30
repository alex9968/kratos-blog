package service

import (
	v1 "github.com/go-kratos/beer-shop/api/moment/interface/v1"
	"github.com/go-kratos/beer-shop/app/moment/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMomentInterface)

type MomentInterface struct {
	v1.UnimplementedMomentInterfaceServer

	uc *biz.UserUseCase

	log *log.Helper
}

func NewMomentInterface(uc *biz.UserUseCase, logger log.Logger) *MomentInterface {
	return &MomentInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
	}
}
