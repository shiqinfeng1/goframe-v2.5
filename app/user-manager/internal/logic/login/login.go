package login

import (
	"context"

	"goframe-v2.5/app/user-manager/internal/dao"
	"goframe-v2.5/app/user-manager/internal/model"
	"goframe-v2.5/app/user-manager/internal/model/do"
	"goframe-v2.5/app/user-manager/internal/service"
)

func init() {
	service.RegisterLogin(New())
}

type sLogin struct{}

func New() *sLogin {
	return &sLogin{}
}
func (l *sLogin) Login(ctx context.Context, ui *model.UserLoginInput) error {
	ci := &do.CaCertInfo{
		UniqueId: ui.Passport,
	}
	_, err := dao.CaCertInfo.Ctx(ctx).Insert(ci)
	return err
}
