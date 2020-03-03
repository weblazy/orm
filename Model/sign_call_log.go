package Model

import (
	"time"
)

// SignCallLog 对外回调日志记录
type SignCallLog struct {
	Id              int64     `gorm:"primary_key;column:id" json:"-"`                  // 自增ID
	AppId           string    `gorm:"column:app_id" json:"appId"`                      // 应用app_id
	CallbackUrl     string    `gorm:"column:callback_url" json:"callbackUrl"`          // 回调地址，请求方传入
	RequestData     string    `gorm:"column:request_data" json:"requestData"`          // 对外调用数据
	ScheduleTime    string    `gorm:"column:schedule_time" json:"scheduleTime"`        // 每次调用时间间隔,逗号隔开,单位秒
	Status          int8      `gorm:"column:status" json:"status"`                     // 0进行中1成功,2失败
	RequestTimes    int       `gorm:"column:request_times" json:"requestTimes"`        // 对外完成请求次数
	ResponseData    string    `gorm:"column:response_data" json:"responseData"`        // 对外调用响应数据
	RequestTakeTime float64   `gorm:"column:request_take_time" json:"requestTakeTime"` // 请求花费时间
	ExtData         string    `gorm:"column:ext_data" json:"extData"`                  // 扩展数据
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt"`              // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updatedAt"`              // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SignCallLog) TableName() string {
	return "sign_call_log"
}

/**
  @desc 创建
*/
func (m *SignCallLog) Create(data *SignCallLog, dbs ...*gormx.DB) error {
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
func (m *SignCallLog) Update(id int64, data *SignCallLog, dbs ...*gormx.DB) error {
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
func (u *SignCallLog) GetById(id int64) (*SignCallLog, error) {
	var resp SignCallLog
	where := "id = ?"
	err := gorm.GetORMByName("openapi").Where(where, id).Take(&resp).Error
	return &resp, err
}
