package data

import (
	"context"

	"github.com/go-kratos/beer-shop/app/moment/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	ctV1 "github.com/go-kratos/beer-shop/api/moment/interface/service/v1"
)

var _ biz.TagRepo = (*tagRepo)(nil)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/tag")),
	}
}

func (r *tagRepo) GetTag(ctx context.Context, id int64) (*biz.Tag, error) {
	reply, err := r.data.bc.GetTag(ctx, &ctV1.GetTagReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	images := make([]biz.Image, 0)
	for _, x := range reply.Image {
		images = append(images, biz.Image{URL: x.Url})
	}
	return &biz.Tag{
		Id:          reply.Id,
		Name:        reply.Name,
		Description: reply.Description,
		Count:       reply.Count,
		Images:      images,
	}, err
}

func (r *tagRepo) ListTag(ctx context.Context, pageNum, pageSize int64) ([]*biz.Tag, error) {
	reply, err := r.data.bc.ListTag(ctx, &ctV1.ListTagReq{
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Tag, 0)
	for _, x := range reply.Results {
		images := make([]biz.Image, 0)
		for _, img := range x.Image {
			images = append(images, biz.Image{URL: img.Url})
		}
		rv = append(rv, &biz.Tag{
			Id:          x.Id,
			Description: x.Description,
			Count:       x.Count,
			Images:      images,
		})
	}
	return rv, err
}
