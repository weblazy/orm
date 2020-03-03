package Model

import (
	"time"
)

// AppApi app和api对应关系表
type AppApi struct {
	Id          int       `gorm:"primary_key;column:id" json:"-"`
	GmtCreate   time.Time `gorm:"column:gmt_create" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
	AppId       int       `gorm:"column:app_id" json:"appId"` // app id
	ApiId       int       `gorm:"column:api_id" json:"apiId"` // api id
}

// TableName get sql table name.获取数据库表名
func (m *AppApi) TableName() string {
	return "app_api"
}

/**
  @desc 创建
*/
func (m *AppApi) Create(data *AppApi, dbs ...*gormx.DB) error {
	var db *gormx.DB
	if len(dbs) > 0 {
		db = dbs[0]
	} else {
		db = gorm.GetORMByName("openapi")
	}
	return db.Create(data).Error
}

/**
  @desc 创建
*/
func (m *AppApi) Update(id int64, data *AppApi, dbs ...*gormx.DB) error {
	var db *gormx.DB
	if len(dbs) > 0 {
		db = dbs[0]
	} else {
		db = gorm.GetORMByName("openapi")
	}
	where := "id = ?"
	return db.Where(where, id).Update(data).Error
}

/**
  @desc 根据id获取单条数据信息
*/
func (u *AppApi) GetById(id int64) (*AppApi, error) {
	var resp AppApi
	where := "id = ?"
	err := gorm.GetORMByName("openapi").Where(where, id).Take(&resp).Error
	return &resp, err
}
