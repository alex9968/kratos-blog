package service

import (
	v1 "kratos-blog/api/moment/interface/v1"
	"kratos-blog/app/moment/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMomentInterface)

// 资料入口，被需要才会被导入
type MomentInterface struct {
	v1.UnimplementedMomentInterfaceServer

	uc *biz.UserUseCase
	cc *biz.CardUseCase
	tc *biz.TagUseCase
	mc *biz.MomentUseCase

	log *log.Helper
}

func NewMomentInterface(uc *biz.UserUseCase,cc *biz.CardUseCase,tc *biz.TagUseCase, mc *biz.MomentUseCase, logger log.Logger) *MomentInterface {
	return &MomentInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
		cc: cc,
		tc: tc,
		mc: mc,
	}
}
