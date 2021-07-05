package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	usV1 "kratos-blog/api/user/service/v1"
	"kratos-blog/app/moment/interface/internal/biz"
)

var _ biz.CardRepo = (*cardRepo)(nil)

type cardRepo struct {
	data *Data
	log  *log.Helper
}

func NewCardRepo(data *Data, logger log.Logger) biz.CardRepo {
	return &cardRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}


func (rp *cardRepo) CreateCard(ctx context.Context,a  *biz.Card) (*biz.Card, error) {
	po, err := rp.data.uc.CreateCard(ctx, &usV1.CreateCardReq{
		CardNo: a.CardNo,
		Name: a.Name,
		Ccv:     a.CCV,
		Expires: a.Expires,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Card{
		Id:      po.Id,
	}, err
}

func (rp *cardRepo) GetCard(ctx context.Context, id int64) (*biz.Card, error) {
	reply, err := rp.data.uc.GetCard(ctx, &usV1.GetCardReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Card{
		Id: reply.Id,
	}, err
}

func (rp *cardRepo) ListCard(ctx context.Context, uid int64) ([]*biz.Card, error) {
	reply, err := rp.data.uc.ListCard(ctx, &usV1.ListCardReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	cards := make([]*biz.Card, 0)
	for _, v := range reply.Results {
		cards = append(cards, &biz.Card{
			Id: v.Id,
			CardNo:	v.CardNo,
			// Name: v.Name,
			CCV: v.Ccv,
		})
	}
	return cards, err
}


func (rp *cardRepo) DeleteCard(ctx context.Context,id int64) (bool, error) {
	po, err := rp.data.uc.DeleteCard(ctx, &usV1.DeleteCardReq{
		Id: id,
	})
	if err != nil {
		return false, err
	}
	return po.Ok, err
}

//
// func (r *cardRepo) ListTag(ctx context.Context, pageNum, pageSize int64) ([]*biz.Tag, error) {
// 	reply, err := r.data.bc.ListTag(ctx, &ctV1.ListTagReq{
// 		PageNum:  pageNum,
// 		PageSize: pageSize,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	rv := make([]*biz.Tag, 0)
// 	for _, x := range reply.Results {
// 		images := make([]biz.Image, 0)
// 		for _, img := range x.Image {
// 			images = append(images, biz.Image{URL: img.Url})
// 		}
// 		rv = append(rv, &biz.Tag{
// 			Id:          x.Id,
// 			Description: x.Description,
// 			Count:       x.Count,
// 			Images:      images,
// 		})
// 	}
// 	return rv, err
// }
//
