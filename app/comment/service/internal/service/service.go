package service

import (
	v1 "kratos-blog/api/comment/service/v1"
	"kratos-blog/app/comment/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCommentService)

type CommentService struct {
	v1.UnimplementedCommentServer

	bc  *biz.CommentUseCase
	log *log.Helper
}

func NewCommentService(bc *biz.CommentUseCase, logger log.Logger) *CommentService {
	return &CommentService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}
