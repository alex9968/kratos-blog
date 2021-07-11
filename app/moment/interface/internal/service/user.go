package service

import (
	"context"

	v1 "kratos-blog/api/moment/interface/v1"
	"kratos-blog/app/moment/interface/internal/biz"
)

func (s *MomentInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	rv, err := s.uc.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &v1.RegisterReply{
		Uid: rv.Id,
	}, err
}

func (s *MomentInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	rv, err := s.uc.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &v1.LoginReply{
		Token: rv,
	}, err
}

func (s *MomentInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	err := s.uc.Logout(ctx, &biz.User{})
	return &v1.LogoutReply{}, err
}



