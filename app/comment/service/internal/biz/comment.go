package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	URL string
}

type Comment struct {
	Id          int64
	Name        string
	Description string
	Count       int64
	Images      []Image
}

type CommentRepo interface {
	CreateComment(ctx context.Context, c *Comment) (*Comment, error)
	UpdateComment(ctx context.Context, c *Comment) (*Comment, error)
	GetComment(ctx context.Context, id int64) (*Comment, error)
	ListComment(ctx context.Context, pageNum, pageSize int64) ([]*Comment, error)
}

type CommentUseCase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUseCase(repo CommentRepo, logger log.Logger) *CommentUseCase {
	return &CommentUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer"))}
}

func (uc *CommentUseCase) Create(ctx context.Context, u *Comment) (*Comment, error) {
	return uc.repo.CreateComment(ctx, u)
}

func (uc *CommentUseCase) Get(ctx context.Context, id int64) (*Comment, error) {
	return uc.repo.GetComment(ctx, id)
}

func (uc *CommentUseCase) Update(ctx context.Context, u *Comment) (*Comment, error) {
	return uc.repo.UpdateComment(ctx, u)
}

func (uc *CommentUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Comment, error) {
	return uc.repo.ListComment(ctx, pageNum, pageSize)
}
