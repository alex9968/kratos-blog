package data

import (
	"context"

	"kratos-blog/app/user/service/internal/biz"
	"kratos-blog/app/user/service/internal/data/ent/user"
	"kratos-blog/app/user/service/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetUsername(u.Username).
		SetAge(u.Age).
		SetPasswordHash(ph).
		Save(ctx)
	return &biz.User{Id: po.ID, Age: po.Age, Username: po.Username}, err
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, err
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (int64, error) {
	po, err := r.data.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(ctx)
	if err != nil {
		return 0, err
	}
	if util.CheckPasswordHash(u.Password, po.PasswordHash) {
		return po.ID, nil
	}
	return 0, nil
}
