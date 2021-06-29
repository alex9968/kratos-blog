package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/api/moment/interface/v1"
	"github.com/go-kratos/beer-shop/app/moment/interface/internal/biz"
)

func (s *MomentInterface) List(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	rv, err := s.uc.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &v1.RegisterReply{
		Id: rv.Id,
	}, err
}

func (s *MomentInterface) Create(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	rv, err := s.uc.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &v1.LoginReply{
		Token: rv,
	}, err
}

func (s *MomentInterface) Delete(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	err := s.uc.Logout(ctx, &biz.User{})
	return &v1.LogoutReply{}, err
}
