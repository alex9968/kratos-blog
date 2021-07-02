package data

import (
	"context"

	"kratos-blog/pkg/util/pagination"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/app/comment/service/internal/biz"
)

var _ biz.CommentRepo = (*commentRepo)(nil)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/comment")),
	}
}

func (r *commentRepo) CreateComment(ctx context.Context, b *biz.Comment) (*biz.Comment, error) {
	po, err := r.data.db.Comment.
		Create().
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		SetImages(b.Images).
		Save(ctx)
	return &biz.Comment{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, err
}

func (r *commentRepo) GetComment(ctx context.Context, id int64) (*biz.Comment, error) {
	po, err := r.data.db.Comment.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Comment{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, err
}

func (r *commentRepo) UpdateComment(ctx context.Context, b *biz.Comment) (*biz.Comment, error) {
	po, err := r.data.db.Comment.
		Create().
		SetName(b.Name).
		SetDescription(b.Description).
		SetCount(b.Count).
		SetImages(b.Images).
		Save(ctx)
	return &biz.Comment{
		Id:          po.ID,
		Description: po.Description,
		Count:       po.Count,
		Images:      po.Images,
	}, err
}

func (r *commentRepo) ListComment(ctx context.Context, pageNum, pageSize int64) ([]*biz.Comment, error) {
	pos, err := r.data.db.Comment.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Comment, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Comment{
			Id:          po.ID,
			Description: po.Description,
			Count:       po.Count,
			Images:      po.Images,
		})
	}
	return rv, err
}
