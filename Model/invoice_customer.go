package Model

import (
	"time"
)

// InvoiceCustomer 客户表
type InvoiceCustomer struct {
	Id               int       `gorm:"primary_key;column:id" json:"-"`                    // 表主键流水
	BuyerName        string    `gorm:"column:buyer_name" json:"buyerName"`                // 购买者名称
	BuyertaxpayerNum string    `gorm:"column:buyertaxpayer_num" json:"buyertaxpayerNum"`  // 购买者纳税人识别号
	BuyerAddress     string    `gorm:"column:buyer_address" json:"buyerAddress"`          // 购买者详细地址
	BuyerTel         string    `gorm:"column:buyer_tel" json:"buyerTel"`                  // 购买者电话
	BuyerBankName    string    `gorm:"column:buyer_bank_name" json:"buyerBankName"`       // 购买者开户银行名称
	BuyerBankAccount string    `gorm:"column:buyer_bank_account" json:"buyerBankAccount"` // 购买者开户行账号
	TakerTel         string    `gorm:"column:taker_tel" json:"takerTel"`                  // 发票接收人手机号
	TakerEmail       string    `gorm:"column:taker_email" json:"takerEmail"`              // 发票接收人邮箱
	WxOpenid         string    `gorm:"column:wx_openid" json:"wxOpenid"`                  // 微信openid
	AlipayOpenid     string    `gorm:"column:alipay_openid" json:"alipayOpenid"`          // 支付宝openid
	CreatedAt        time.Time `gorm:"column:created_at" json:"createdAt"`                // 创建日期
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updatedAt"`                // 更新日期
}

// TableName get sql table name.获取数据库表名
func (m *InvoiceCustomer) TableName() string {
	return "invoice_customer"
}

/**
  @desc 创建
*/
func (m *InvoiceCustomer) Create(data *InvoiceCustomer, dbs ...*gormx.DB) error {
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
func (m *InvoiceCustomer) Update(id int64, data *InvoiceCustomer, dbs ...*gormx.DB) error {
	var db *gormx.DB
	if len(dbs) > 0 {
		db = dbs[0]
	} else {
		db = gorm.GetORMByName("electronic_invoice")
	}
	where := "id = ?"
	return db.Where(where, id).Update(data).Error
}

/**
  @desc 根据id获取单条数据信息
*/
func (u *InvoiceCustomer) GetById(id int64) (*InvoiceCustomer, error) {
	var resp InvoiceCustomer
	where := "id = ?"
	err := gorm.GetORMByName("electronic_invoice").Where(where, id).Take(&resp).Error
	return &resp, err
}
