// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CaCertDeviceId is the golang structure of table ca_cert_deviceId for DAO operations like Where/Data.
type CaCertDeviceId struct {
	g.Meta         `orm:"table:ca_cert_deviceId, do:true"`
	Id             interface{} //
	CustomerId     interface{} // 客户id
	CustomerIdcard interface{} // 用户身份证号
	Identity       interface{} // 用户定义的证书标识
	DeviceId       interface{} // 设备号
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
