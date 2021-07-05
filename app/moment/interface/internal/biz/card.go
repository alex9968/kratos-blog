package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Card struct {
	Id      int64
	CardNo  string
	Name    string
	CCV     string
	Expires string
}

type CardRepo interface {
	CreateCard(ctx context.Context, c *Card) (*Card, error)
	GetCard(ctx context.Context, id int64) (*Card, error)
	ListCard(ctx context.Context, id int64) ([]*Card, error)
	DeleteCard(ctx context.Context, id int64) (bool, error)
}

type CardUseCase struct {
	repo CardRepo
	log  *log.Helper
}

func NewCardUseCase(repo CardRepo, logger log.Logger) *CardUseCase {
	return &CardUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/card"))}
}

func (cc *CardUseCase) Create(ctx context.Context, u *Card) (*Card, error) {
	return cc.repo.CreateCard(ctx, u)
}


func (cc *CardUseCase) Get(ctx context.Context, id int64) (*Card, error) {
	return cc.repo.GetCard(ctx, id)
}

func (cc *CardUseCase) List(ctx context.Context, uid int64) ([]*Card, error) {
	return cc.repo.ListCard(ctx, uid)
}

func (cc *CardUseCase) Delete(ctx context.Context, uid int64) (bool, error) {
	return cc.repo.DeleteCard(ctx, uid)
}
