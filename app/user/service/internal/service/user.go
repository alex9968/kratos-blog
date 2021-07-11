package service

import (
	"context"

	v1 "kratos-blog/api/user/service/v1"
	"kratos-blog/app/user/service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &v1.CreateUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, err
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	return &v1.GetUserReply{
		Id:       rv.Id,
		Username: rv.Username,
	}, err
}

func (s *UserService) GetUserMap(ctx context.Context, req *v1.GetUserMapReq) (*v1.GetUserMapReply, error) {
	data, err := s.uc.GetMap(ctx, req.Ids)
	res := make(map[int64]*v1.GetUserMapReply_User,0)
	for k, v := range data {
		res[k] = &v1.GetUserMapReply_User{
			Id: v.Id,
			Username: v.Username,
		}
	}
	return &v1.GetUserMapReply{ Users: res}, err
}


func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	id, err := s.uc.VerifyPassword(ctx, &biz.User{Username: req.Username, Password: req.Password})
	return &v1.VerifyPasswordReply{
		Ok: id != 0,
		Id: id,
	}, err
}
