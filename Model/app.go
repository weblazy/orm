package Model

import (
	"time"
)

// App app列表
type App struct {
	Id          int       `gorm:"primary_key;column:id" json:"-"`
	Name        string    `gorm:"column:name" json:"name"`                // 应用名称
	AppId       string    `gorm:"column:app_id" json:"appId"`             // app id
	AppKey      string    `gorm:"column:app_key" json:"appKey"`           // app key
	DId         int       `gorm:"column:dId" json:"dId"`                  // 公司dId
	CId         int       `gorm:"column:cId" json:"cId"`                  // 公司cId
	State       int8      `gorm:"column:state" json:"state"`              // 状态 1-正常 2-停用 3-删除
	PubKey      string    `gorm:"column:pub_key" json:"pubKey"`           // 接收回调方rsa公钥
	SmPriKey    string    `gorm:"column:sm_pri_key" json:"smPriKey"`      // 商米rsa私钥
	SmPubKey    string    `gorm:"column:sm_pub_key" json:"smPubKey"`      // 商米rsa公钥
	CallbackUrl string    `gorm:"column:callback_url" json:"callbackUrl"` // 接收回调的地址
	IsInside    int8      `gorm:"column:is_inside" json:"isInside"`       // 是否内部回调，0 表示是，1 表示外部调用
	GmtCreate   time.Time `gorm:"column:gmt_create" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
	WhiteList   string    `gorm:"column:white_list" json:"whiteList"` // ip白名单列表
}

// TableName get sql table name.获取数据库表名
func (m *App) TableName() string {
	return "app"
}

/**
  @desc 创建
*/
func (m *App) Create(data *App, dbs ...*gormx.DB) error {
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
func (m *App) Update(id int64, data *App, dbs ...*gormx.DB) error {
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
func (u *App) GetById(id int64) (*App, error) {
	var resp App
	where := "id = ?"
	err := gorm.GetORMByName("openapi").Where(where, id).Take(&resp).Error
	return &resp, err
}
