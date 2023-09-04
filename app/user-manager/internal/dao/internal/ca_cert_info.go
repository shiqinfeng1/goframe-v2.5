// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CaCertInfoDao is the data access object for table ca_cert_info.
type CaCertInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CaCertInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CaCertInfoColumns defines and stores column names for table ca_cert_info.
type CaCertInfoColumns struct {
	Id                string //
	UniqueId          string // 证书id
	Version           string // 版本号
	Identity          string // 用户定义的证书标识
	CertType          string // 证书类型
	IssuerIdentity    string // 颁发机构标识
	Subject           string // 证书主体信息
	ValidityNotbefore string // 证书有效期开始
	ValidityNotafter  string // 证书有效期结束
	OperatorName      string // 申请者姓名
	HolderName        string // 持有者姓名
	HolderIdcard      string // 持有者证件号码
	CustomerId        string // 客户id
	SerialNumber      string // 证书序列号
	Status            string // 证书状态
	Dealing           string // 审批处理中
	ApplyTime         string // 证书申请时间
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// caCertInfoColumns holds the columns for table ca_cert_info.
var caCertInfoColumns = CaCertInfoColumns{
	Id:                "id",
	UniqueId:          "unique_id",
	Version:           "version",
	Identity:          "identity",
	CertType:          "cert_type",
	IssuerIdentity:    "issuer_identity",
	Subject:           "subject",
	ValidityNotbefore: "validity_notbefore",
	ValidityNotafter:  "validity_notafter",
	OperatorName:      "operator_name",
	HolderName:        "holder_name",
	HolderIdcard:      "holder_idcard",
	CustomerId:        "customer_id",
	SerialNumber:      "serial_number",
	Status:            "status",
	Dealing:           "dealing",
	ApplyTime:         "apply_time",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewCaCertInfoDao creates and returns a new DAO object for table data access.
func NewCaCertInfoDao() *CaCertInfoDao {
	return &CaCertInfoDao{
		group:   "default",
		table:   "ca_cert_info",
		columns: caCertInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CaCertInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CaCertInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CaCertInfoDao) Columns() CaCertInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CaCertInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CaCertInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CaCertInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
