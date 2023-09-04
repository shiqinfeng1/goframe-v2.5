// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CaCertInfo is the golang structure of table ca_cert_info for DAO operations like Where/Data.
type CaCertInfo struct {
	g.Meta            `orm:"table:ca_cert_info, do:true"`
	Id                interface{} //
	UniqueId          interface{} // 证书id
	Version           interface{} // 版本号
	Identity          interface{} // 用户定义的证书标识
	CertType          interface{} // 证书类型
	IssuerIdentity    interface{} // 颁发机构标识
	Subject           interface{} // 证书主体信息
	ValidityNotbefore *gtime.Time // 证书有效期开始
	ValidityNotafter  *gtime.Time // 证书有效期结束
	OperatorName      interface{} // 申请者姓名
	HolderName        interface{} // 持有者姓名
	HolderIdcard      interface{} // 持有者证件号码
	CustomerId        interface{} // 客户id
	SerialNumber      interface{} // 证书序列号
	Status            interface{} // 证书状态
	Dealing           interface{} // 审批处理中
	ApplyTime         *gtime.Time // 证书申请时间
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
}
