package data

import (
	"context"

	"kratos-blog/app/moment/interface/internal/biz"
	"kratos-blog/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.TagRepo = (*tagRepo)(nil)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

type Image struct {
	URL string
}

type Tag struct {
	gorm.Model
	UserId     uint
	Name        string
	Description string
	Count       int64
	Images      []Image
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/tag")),
	}
}


func (tr *tagRepo) CreateTag(ctx context.Context, b *biz.Tag) (*biz.Tag, error) {
	t := Tag{Model: gorm.Model{ID: b.Id}, UserId: b.UserId}
	result := tr.data.db.WithContext(ctx).Create(t)
	return &biz.Tag{
		Id: t.ID,
	}, result.Error
}

func (tr *tagRepo) GetTag(ctx context.Context, id int64) (*biz.Tag, error) {
	t := Tag{}
	result := tr.data.db.WithContext(ctx).First(&t, id)
	return &biz.Tag{
		Id: t.ID,
	}, result.Error
}

func (r *tagRepo) UpdateTag(ctx context.Context, b *biz.Tag) (*biz.Tag, error) {
	t := Tag{}
	result := r.data.db.WithContext(ctx).First(&t, b.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	t.UserId = b.UserId
	result = r.data.db.WithContext(ctx).Save(&t)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Tag{
		Id: t.ID,
	}, nil
}

func (r *tagRepo) ListTag(ctx context.Context, pageNum, pageSize int64) ([]*biz.Tag, error) {
	var ts []Tag
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&ts)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.Tag, 0)
	for _, t := range ts {
		rv = append(rv, &biz.Tag{
			Id: t.ID,
		})
	}
	return rv, nil
}



