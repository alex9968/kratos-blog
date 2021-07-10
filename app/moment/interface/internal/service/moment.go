package service

import (
	"context"

	v1 "kratos-blog/api/moment/interface/v1"
	"kratos-blog/app/moment/interface/internal/biz"
)


func (s *MomentInterface) GetMoment(ctx context.Context, req *v1.GetMomentReq) (*v1.GetMomentReply, error) {
	rv, err := s.mc.Get(ctx,  req.Id)
	return &v1.GetMomentReply{
		Id: rv.Id,
		UserId: rv.UserId,
		Content: rv.Content,
	}, err
}

func (s *MomentInterface) ListMoment(ctx context.Context, req *v1.ListMomentReq) (*v1.ListMomentReply, error) {
	rv, err := s.mc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	rs := make([]*v1.ListMomentReply_Moment, 0)
	for _, v := range rv {
		rs = append(rs, &v1.ListMomentReply_Moment{
			Content: v.Content,
			Id:      v.Id,
		})
	}
	return &v1.ListMomentReply{
		Results: rs,
	}, err
}

func (s *MomentInterface) CreateMoment(ctx context.Context, req *v1.CreateMomentReq) (*v1.CreateMomentReply, error) {
	rv, err := s.mc.Create(ctx, &biz.Moment{
		Content: req.Content,
	})
	return &v1.CreateMomentReply{
		Id: rv.Id,
	}, err
}

func (s *MomentInterface) DeleteMoment(ctx context.Context, req *v1.DeleteMomentReq) (*v1.DeleteMomentReply, error) {
	ok, err := s.mc.Delete(ctx, req.Id)
	return &v1.DeleteMomentReply{
		Ok: ok,
	}, err
}
