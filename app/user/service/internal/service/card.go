package service

import (
	"context"

	v1 "kratos-blog/api/user/service/v1"
	"kratos-blog/app/user/service/internal/biz"
)

func (s *UserService) CreateCard(ctx context.Context, req *v1.CreateCardReq) (*v1.CreateCardReply, error) {
	rv, err := s.cc.Create(ctx, &biz.Card{
		CardNo:  req.CardNo,
		CCV:     req.Ccv,
		Expires: req.Expires,
		Name:    req.Name,
	})
	return &v1.CreateCardReply{
		Id: rv.Id,
	}, err
}

func (s *UserService) GetCard(ctx context.Context, req *v1.GetCardReq) (*v1.GetCardReply, error) {
	rv, err := s.cc.Get(ctx, req.Id)
	return &v1.GetCardReply{
		Id:      rv.Id,
		CardNo:  rv.CardNo,
		Ccv:     rv.CCV,
		Expires: rv.Expires,
	}, err
}

func (s *UserService) ListCard(ctx context.Context, req *v1.ListCardReq) (*v1.ListCardReply, error) {
	rv, err := s.cc.List(ctx, req.Uid)
	rs := make([]*v1.ListCardReply_Card, 0)
	for _, x := range rv {
		rs = append(rs, &v1.ListCardReply_Card{
			Id:      x.Id,
			CardNo:  x.CardNo,
			Ccv:     x.CCV,
			Expires: x.Expires,
		})
	}
	return &v1.ListCardReply{
		Results: rs,
	}, err
}

func (s *UserService) DeleteCard(ctx context.Context, req *v1.DeleteCardReq) (*v1.DeleteCardReply, error) {
	ok, err := s.cc.Delete(ctx, req.Uid)
	if err != nil {
	}
	return &v1.DeleteCardReply{
		Ok: ok,
	}, err
}
