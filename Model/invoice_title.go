package Model

import (
	"time"
)

// InvoiceTitle 电子发票抬头表
type InvoiceTitle struct {
	Id             int       `gorm:"primary_key;column:id" json:"-"`               // 表主键
	TaxPayerNum    string    `gorm:"column:tax_payer_num" json:"taxPayerNum"`      // 纳税人识别号
	EnterpriseName string    `gorm:"column:enterprise_name" json:"enterpriseName"` // 企业抬头名称
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`           // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updatedAt"`           // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *InvoiceTitle) TableName() string {
	return "invoice_title"
}

/**
  @desc 创建
*/
func (m *InvoiceTitle) Create(data *InvoiceTitle, dbs ...*gormx.DB) error {
	var db *gormx.DB
	if len(dbs) > 0 {
		db = dbs[0]
	} else {
		db = gorm.GetORMByName("electronic_invoice")
	}
	return db.Create(data).Error
}

/**
  @desc 创建
*/
func (m *InvoiceTitle) Update(id int64, data *InvoiceTitle, dbs ...*gormx.DB) error {
	var db *gormx.DB
	if len(dbs) > 0 {
		db = dbs[0]
	} else {
		db = gorm.GetORMByName("electronic_invoice")
	}
	where := "id = ?"
	return db.Model(InvoiceTitle{}).Where(where, id).Update(data).Error
}

/**
  @desc 根据id获取单条数据信息
*/
func (u *InvoiceTitle) GetById(id int64) (*InvoiceTitle, error) {
	var resp InvoiceTitle
	where := "id = ?"
	err := gorm.GetORMByName("electronic_invoice").Where(where, id).Take(&resp).Error
	return &resp, err
}
