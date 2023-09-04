// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CaIssuerInfo is the golang structure for table ca_issuer_info.
type CaIssuerInfo struct {
	Id                 uint64      `json:"id"                 description:""`                          //
	Identity           string      `json:"identity"           description:"自定义的机构标识"`                  // 自定义的机构标识
	Country            string      `json:"country"            description:"国家"`                        // 国家
	Province           string      `json:"province"           description:"省份"`                        // 省份
	City               string      `json:"city"               description:"城市"`                        // 城市
	Locality           string      `json:"locality"           description:"区"`                         // 区
	StreetAddress      string      `json:"streetAddress"      description:"街道"`                        // 街道
	Organization       string      `json:"organization"       description:"组织"`                        // 组织
	OrganizationalUnit string      `json:"organizationalUnit" description:"子组织"`                       // 子组织
	PostalCode         string      `json:"postalCode"         description:"邮编"`                        // 邮编
	CommonName         string      `json:"commonName"         description:"名称"`                        // 名称
	ExtraName          string      `json:"extraName"          description:"额外名称"`                      // 额外名称
	LegalPerson        string      `json:"legalPerson"        description:"法人"`                        // 法人
	CredentialsType    uint        `json:"credentialsType"    description:"证件类型 1:统一社会信用代码,2:组织机构代码证"` // 证件类型 1:统一社会信用代码,2:组织机构代码证
	CredentialsNumber  string      `json:"credentialsNumber"  description:"证件号码"`                      // 证件号码
	ContactNumber      string      `json:"contactNumber"      description:"机构联系方式"`                    // 机构联系方式
	CreatedAt          *gtime.Time `json:"createdAt"          description:"创建时间"`                      // 创建时间
	UpdatedAt          *gtime.Time `json:"updatedAt"          description:"更新时间"`                      // 更新时间
}
