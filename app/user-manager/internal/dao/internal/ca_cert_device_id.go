// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CaCertDeviceIdDao is the data access object for table ca_cert_deviceId.
type CaCertDeviceIdDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns CaCertDeviceIdColumns // columns contains all the column names of Table for convenient usage.
}

// CaCertDeviceIdColumns defines and stores column names for table ca_cert_deviceId.
type CaCertDeviceIdColumns struct {
	Id             string //
	CustomerId     string // 客户id
	CustomerIdcard string // 用户身份证号
	Identity       string // 用户定义的证书标识
	DeviceId       string // 设备号
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// caCertDeviceIdColumns holds the columns for table ca_cert_deviceId.
var caCertDeviceIdColumns = CaCertDeviceIdColumns{
	Id:             "id",
	CustomerId:     "customer_id",
	CustomerIdcard: "customer_idcard",
	Identity:       "identity",
	DeviceId:       "deviceId",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewCaCertDeviceIdDao creates and returns a new DAO object for table data access.
func NewCaCertDeviceIdDao() *CaCertDeviceIdDao {
	return &CaCertDeviceIdDao{
		group:   "default",
		table:   "ca_cert_deviceId",
		columns: caCertDeviceIdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CaCertDeviceIdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CaCertDeviceIdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CaCertDeviceIdDao) Columns() CaCertDeviceIdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CaCertDeviceIdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CaCertDeviceIdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CaCertDeviceIdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
