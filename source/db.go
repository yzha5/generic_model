package source

import (
	"GenericModel/config"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var (
	dbLock       sync.Mutex
	gormInstance *gorm.DB
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", config.MysqlDsn)
	if err != nil {
		//打印内容并退出程序
		log.Fatalf("创建数据库连接发生错误：%v", err)
	}

	//不使用复数表名
	db.SingularTable(true)

	//TODO:
	db.LogMode(true)

	//设置表前辍
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "admin_" + defaultTableName
	//}

	return db
}

func InstanceDB() *gorm.DB {
	if gormInstance != nil {
		return gormInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if gormInstance != nil {
		return gormInstance
	}
	gormInstance = NewDB()
	return gormInstance
}
