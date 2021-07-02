package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/api/comment/service/v1"
	"github.com/go-kratos/beer-shop/app/comment/service/internal/biz"
)

func (s *CommentService) CreateComment(ctx context.Context, req *v1.CreateCommentReq) (*v1.CreateCommentReply, error) {
	b := &biz.Comment{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Create(ctx, b)
	img := make([]*v1.CreateCommentReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.CreateCommentReply_Image{Url: i.URL})
	}
	return &v1.CreateCommentReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CommentService) GetComment(ctx context.Context, req *v1.GetCommentReq) (*v1.GetCommentReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	img := make([]*v1.GetCommentReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.GetCommentReply_Image{Url: i.URL})
	}
	return &v1.GetCommentReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CommentService) UpdateComment(ctx context.Context, req *v1.UpdateCommentReq) (*v1.UpdateCommentReply, error) {
	b := &biz.Comment{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Update(ctx, b)
	img := make([]*v1.UpdateCommentReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.UpdateCommentReply_Image{Url: i.URL})
	}
	return &v1.UpdateCommentReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CommentService) ListComment(ctx context.Context, req *v1.ListCommentReq) (*v1.ListCommentReply, error) {
	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListCommentReply_Comment, 0)
	for _, x := range rv {
		img := make([]*v1.ListCommentReply_Comment_Image, 0)
		for _, i := range x.Images {
			img = append(img, &v1.ListCommentReply_Comment_Image{Url: i.URL})
		}
		rs = append(rs, &v1.ListCommentReply_Comment{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &v1.ListCommentReply{
		Results: rs,
	}, err
}
