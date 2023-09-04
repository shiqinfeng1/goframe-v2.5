// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CaIssuerInfo is the golang structure of table ca_issuer_info for DAO operations like Where/Data.
type CaIssuerInfo struct {
	g.Meta             `orm:"table:ca_issuer_info, do:true"`
	Id                 interface{} //
	Identity           interface{} // 自定义的机构标识
	Country            interface{} // 国家
	Province           interface{} // 省份
	City               interface{} // 城市
	Locality           interface{} // 区
	StreetAddress      interface{} // 街道
	Organization       interface{} // 组织
	OrganizationalUnit interface{} // 子组织
	PostalCode         interface{} // 邮编
	CommonName         interface{} // 名称
	ExtraName          interface{} // 额外名称
	LegalPerson        interface{} // 法人
	CredentialsType    interface{} // 证件类型 1:统一社会信用代码,2:组织机构代码证
	CredentialsNumber  interface{} // 证件号码
	ContactNumber      interface{} // 机构联系方式
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
}
