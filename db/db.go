package db

import (
	"errors"
	"fmt"
	"log"
	"orm/conf"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DBMap sync.Map
var defaultName = "dbDefault"

var (
	// ErrRecordNotFound record not found error, happens when haven't find any matched data when looking up with a struct
	ErrRecordNotFound = errors.New("record not found")
	// ErrInvalidSQL invalid SQL error, happens when you passed invalid SQL
	ErrInvalidSQL = errors.New("invalid SQL")
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("no valid transaction")
	// ErrCantStartTransaction can't start transaction when you are trying to start one with `Begin`
	ErrCantStartTransaction = errors.New("can't start transaction")
	// ErrUnaddressable unaddressable value
	ErrUnaddressable = errors.New("using unaddressable value")
)

// 初始化Gorm
func NewDB(dbname string) {

	var orm *gorm.DB
	var err error

	for orm, err = openORM(dbname); err != nil; {
		fmt.Println("Database connection exception! 5 seconds to retry")
		time.Sleep(5 * time.Second)
		orm, err = openORM(dbname)
	}

	DBMap.LoadOrStore(dbname, orm)
}

// 设置获取db的默认值
func SetDefaultName(dbname string) {
	defaultName = dbname
}

// 初始化Gorm
func UpdateDB(dbname string) error {

	v, _ := DBMap.Load(dbname)

	orm, err := openORM(dbname)

	DBMap.Delete(dbname)
	DBMap.LoadOrStore(dbname, orm)

	err = v.(*gorm.DB).Close()
	if err != nil {
		return err
	}

	return nil
}

// 通过名称获取Gorm实例
func DB(dbname string) *gorm.DB {
	v, _ := DBMap.Load(dbname)
	return v.(*gorm.DB)
}

// 获取默认的Gorm实例
func GetORM() *gorm.DB {

	v, _ := DBMap.Load(defaultName)
	return v.(*gorm.DB)
}

func openORM(dbname string) (*gorm.DB, error) {
	dbConf := conf.Conf.Db
	port := dbConf.Port
	connectStr := dbConf.User + ":" + dbConf.Pwd + "@tcp(" + dbConf.Host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectStr)
	db.SingularTable(true)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, err
}
