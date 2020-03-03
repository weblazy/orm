package Model

import (
	"time"
)

// Api 对外提供的api列表
type Api struct {
	Id          int       `gorm:"primary_key;column:id" json:"-"`
	GmtCreate   time.Time `gorm:"column:gmt_create" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
	Router      string    `gorm:"column:router" json:"router"`    // 路由地址
	ApiName     string    `gorm:"column:api_name" json:"apiName"` // api对应的中文名称
	Method      string    `gorm:"column:method" json:"method"`
	Host        string    `gorm:"column:host" json:"host"` // 业务接口host
}

// TableName get sql table name.获取数据库表名
func (m *Api) TableName() string {
	return "api"
}

/**
  @desc 创建
*/
func (m *Api) Create(data *Api, dbs ...*gormx.DB) error {
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
func (m *Api) Update(id int64, data *Api, dbs ...*gormx.DB) error {
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
func (u *Api) GetById(id int64) (*Api, error) {
	var resp Api
	where := "id = ?"
	err := gorm.GetORMByName("openapi").Where(where, id).Take(&resp).Error
	return &resp, err
}
