// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CaCertDeviceId is the golang structure for table ca_cert_deviceId.
type CaCertDeviceId struct {
	Id             uint64      `json:"id"             description:""`          //
	CustomerId     int64       `json:"customerId"     description:"客户id"`      // 客户id
	CustomerIdcard string      `json:"customerIdcard" description:"用户身份证号"`    // 用户身份证号
	Identity       string      `json:"identity"       description:"用户定义的证书标识"` // 用户定义的证书标识
	DeviceId       string      `json:"deviceId"       description:"设备号"`       // 设备号
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`      // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`      // 更新时间
}
