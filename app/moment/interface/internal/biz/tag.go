package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	URL string
}

type Tag struct {
	Id         uint
	UserId          uint
	Name        string
	Description string
	Count       int64
	Images      []Image
}

type TagRepo interface {
	GetTag(ctx context.Context, id int64) (*Tag, error)
	ListTag(ctx context.Context, pageNum, pageSize int64) ([]*Tag, error)
}

type TagUseCase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/tag"))}
}

func (uc *TagUseCase) Get(ctx context.Context, id int64) (*Tag, error) {
	return uc.repo.GetTag(ctx, id)
}

func (uc *TagUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Tag, error) {
	return uc.repo.ListTag(ctx, pageNum, pageSize)
}
