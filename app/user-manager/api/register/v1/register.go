package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfo struct {
	QuestionIndex1 uint `p:"questionIndex1" v:"required#必须输入问题索引1" dc:"问题索引1"`
	QuestionIndex2 uint `p:"questionIndex2" v:"required#必须输入问题索引2" dc:"问题索引2"`
	QuestionIndex3 uint `p:"questionIndex3" v:"required#必须输入问题索引3" dc:"问题索引3"`
}

type RegisterReq struct {
	g.Meta     `path:"/register" tags:"用户管理" method:"post" summary:"用户注册"`
	Name       string `p:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password   string `p:"password" v:"required#用户密码不能为空" dc:"用户密码"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空" dc:"图形验证码的值"`
	VerifyKey  string `p:"verifyKey"  v:"required#验证码key不能为空" dc:"图形验证码的key"`
	UserInfo
}

type RegisterRes struct {
	g.Meta `mime:"application/json"`
	Result string `json:"result" dc:"结果。成功:success"`
}
