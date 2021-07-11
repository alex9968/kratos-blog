package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	usV1 "kratos-blog/api/user/service/v1"
	"kratos-blog/app/moment/interface/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (rp *userRepo) Register(ctx context.Context, u *biz.User) (*biz.User, error) {
	reply, err := rp.data.uc.CreateUser(ctx, &usV1.CreateUserReq{
		Username: u.Username,
		Password: u.Password,
	})
	return &biz.User{
		Id:       reply.Id,
		Username: reply.Username,
	}, err
}

func (rp *userRepo) Login(ctx context.Context, u *biz.User) (string, error) {
	reply, err := rp.data.uc.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return "", err
	}
	if reply.Ok {
		return "some_token", nil
	}
	return "", errors.New("login failed")
}

func (rp *userRepo) Logout(ctx context.Context, u *biz.User) error {
	return nil
}

func (rp *userRepo) GetUserMap(ctx context.Context, ids []int64) (map[int64]*biz.User, error) {
	reply, err := rp.data.uc.GetUserMap(ctx, &usV1.GetUserMapReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	res := make(map[int64]*biz.User, 0)
	for k, v := range reply.Users {
		res[k] = &biz.User{
			Id:       v.Id,
			Username: v.Username,
		}
	}
	return res, err
}

func (rp *userRepo) GetUsers(ctx context.Context, id int64) (*biz.User, error) {
	reply, err := rp.data.uc.GetUser(ctx, &usV1.GetUserReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	// images := make([]biz.Image, 0)
	// for _, x := range reply.Image {
	// 	images = append(images, biz.Image{URL: x.Url})
	// }
	return &biz.User{
		Id: reply.Id,
	}, err
}

func (rp *userRepo) GetCard(ctx context.Context, id int64) (*biz.Card, error) {
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

//
// func (r *userRepo) ListTag(ctx context.Context, pageNum, pageSize int64) ([]*biz.Tag, error) {
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
