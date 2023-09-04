// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CaCertInfo is the golang structure for table ca_cert_info.
type CaCertInfo struct {
	Id                uint64      `json:"id"                description:""`          //
	UniqueId          int64       `json:"uniqueId"          description:"证书id"`      // 证书id
	Version           uint        `json:"version"           description:"版本号"`       // 版本号
	Identity          string      `json:"identity"          description:"用户定义的证书标识"` // 用户定义的证书标识
	CertType          uint        `json:"certType"          description:"证书类型"`      // 证书类型
	IssuerIdentity    string      `json:"issuerIdentity"    description:"颁发机构标识"`    // 颁发机构标识
	Subject           string      `json:"subject"           description:"证书主体信息"`    // 证书主体信息
	ValidityNotbefore *gtime.Time `json:"validityNotbefore" description:"证书有效期开始"`   // 证书有效期开始
	ValidityNotafter  *gtime.Time `json:"validityNotafter"  description:"证书有效期结束"`   // 证书有效期结束
	OperatorName      string      `json:"operatorName"      description:"申请者姓名"`     // 申请者姓名
	HolderName        string      `json:"holderName"        description:"持有者姓名"`     // 持有者姓名
	HolderIdcard      string      `json:"holderIdcard"      description:"持有者证件号码"`   // 持有者证件号码
	CustomerId        int64       `json:"customerId"        description:"客户id"`      // 客户id
	SerialNumber      string      `json:"serialNumber"      description:"证书序列号"`     // 证书序列号
	Status            string      `json:"status"            description:"证书状态"`      // 证书状态
	Dealing           uint        `json:"dealing"           description:"审批处理中"`     // 审批处理中
	ApplyTime         *gtime.Time `json:"applyTime"         description:"证书申请时间"`    // 证书申请时间
	CreatedAt         *gtime.Time `json:"createdAt"         description:"创建时间"`      // 创建时间
	UpdatedAt         *gtime.Time `json:"updatedAt"         description:"更新时间"`      // 更新时间
}
