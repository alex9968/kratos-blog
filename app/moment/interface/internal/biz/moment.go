package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Moment struct {
	Id       int64
	UserId 	int64
	Content string
}

type MomentRepo interface {
	CreateMoment(ctx context.Context, u *Moment) (*Moment, error)
	GetMoment(ctx context.Context, id int64) (*Moment, error)
	ListMoment(ctx context.Context, pageNum, pageSize int64) ([]*Moment, error)
	UpdateMoment(ctx context.Context, b *Moment) (*Moment, error)
}

type MomentUseCase struct {
	repo MomentRepo
	log  *log.Helper
}

func NewMomentUseCase(repo MomentRepo, logger log.Logger) *MomentUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &MomentUseCase{
		repo: repo,
		log:  log,
	}
}

func (uc *MomentUseCase) Create(ctx context.Context, u *Moment) (*Moment, error) {
	return uc.repo.CreateMoment(ctx, u)
}

func (uc *MomentUseCase) List(ctx context.Context, u *Moment) ([]*Moment, error) {
	return uc.repo.ListMoment(ctx, 1, 10)
}
