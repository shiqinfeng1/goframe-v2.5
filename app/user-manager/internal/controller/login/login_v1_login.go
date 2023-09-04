package login

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "goframe-v2.5/app/user-manager/api/login/v1"
	"goframe-v2.5/app/user-manager/internal/model"
	"goframe-v2.5/app/user-manager/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	ui := &model.UserLoginInput{
		Passport: req.Name,
		Password: req.Password,
	}
	service.Login().Login(ctx, ui)
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
