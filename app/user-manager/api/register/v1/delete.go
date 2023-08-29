package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteReq struct {
	g.Meta     `path:"/reset" tags:"用户管理" method:"post" summary:"用户注销"`
	Name       string `p:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password   string `p:"password" v:"required#用户密码不能为空" dc:"用户密码"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空" dc:"图形验证码的值"`
	VerifyKey  string `p:"verifyKey"  v:"required#验证码key不能为空" dc:"图形验证码的key"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
	Result string `json:"result" dc:"结果。成功:success"`
}
