package data

import (
	"context"

	"kratos-blog/pkg/util/pagination"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-blog/app/moment/interface/internal/biz"
)

var _ biz.MomentRepo = (*momentRepo)(nil)

type momentRepo struct {
	data *Data
	log  *log.Helper
}

type Moment struct {
	gorm.Model
	Id int64
	UserId    int64   `gorm:"column:user_id"`
	Content string `gorm:"column:content"`
}

func NewMomentRepo(data *Data, logger log.Logger) biz.MomentRepo {
	return &momentRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/moment")),
	}
}

func (r *momentRepo) CreateMoment(ctx context.Context, b *biz.Moment) (*biz.Moment, error) {
	o := Moment{Content: b.Content}
	result := r.data.db.WithContext(ctx).Create(&o)
	return &biz.Moment{
		Id: o.Id,
	}, result.Error
}

func (r *momentRepo) GetMoment(ctx context.Context, id int64) (*biz.Moment, error) {
	o := Moment{}
	result := r.data.db.WithContext(ctx).First(&o, id)
	return &biz.Moment{
		Id: o.Id,
		UserId: o.UserId,
		Content: o.Content,
	}, result.Error
}

func (r *momentRepo) UpdateMoment(ctx context.Context, b *biz.Moment) (*biz.Moment, error) {
	o := Moment{}
	result := r.data.db.WithContext(ctx).First(&o, b.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	o.UserId = b.UserId
	result = r.data.db.WithContext(ctx).Save(&o)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Moment{
		Id: o.Id,
	}, nil
}

func (r *momentRepo) ListMoment(ctx context.Context, pageNum, pageSize int64) ([]*biz.Moment, error) {
	var os []Moment
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&os)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.Moment, 0)
	for _, o := range os {
		rv = append(rv, &biz.Moment{
			Id: o.Id,
		})
	}
	return rv, nil
}

func (r *momentRepo) DeleteMoment(ctx context.Context, id int64) (bool, error) {
	o := Moment{}
	result := r.data.db.WithContext(ctx).Delete(&o, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
