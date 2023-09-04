// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CaIssuerInfoDao is the data access object for table ca_issuer_info.
type CaIssuerInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns CaIssuerInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CaIssuerInfoColumns defines and stores column names for table ca_issuer_info.
type CaIssuerInfoColumns struct {
	Id                 string //
	Identity           string // 自定义的机构标识
	Country            string // 国家
	Province           string // 省份
	City               string // 城市
	Locality           string // 区
	StreetAddress      string // 街道
	Organization       string // 组织
	OrganizationalUnit string // 子组织
	PostalCode         string // 邮编
	CommonName         string // 名称
	ExtraName          string // 额外名称
	LegalPerson        string // 法人
	CredentialsType    string // 证件类型 1:统一社会信用代码,2:组织机构代码证
	CredentialsNumber  string // 证件号码
	ContactNumber      string // 机构联系方式
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
}

// caIssuerInfoColumns holds the columns for table ca_issuer_info.
var caIssuerInfoColumns = CaIssuerInfoColumns{
	Id:                 "id",
	Identity:           "identity",
	Country:            "country",
	Province:           "province",
	City:               "city",
	Locality:           "locality",
	StreetAddress:      "street_address",
	Organization:       "organization",
	OrganizationalUnit: "organizational_unit",
	PostalCode:         "postal_code",
	CommonName:         "common_name",
	ExtraName:          "extra_name",
	LegalPerson:        "legal_person",
	CredentialsType:    "credentials_type",
	CredentialsNumber:  "credentials_number",
	ContactNumber:      "contact_number",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewCaIssuerInfoDao creates and returns a new DAO object for table data access.
func NewCaIssuerInfoDao() *CaIssuerInfoDao {
	return &CaIssuerInfoDao{
		group:   "default",
		table:   "ca_issuer_info",
		columns: caIssuerInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CaIssuerInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CaIssuerInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CaIssuerInfoDao) Columns() CaIssuerInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CaIssuerInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CaIssuerInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CaIssuerInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
