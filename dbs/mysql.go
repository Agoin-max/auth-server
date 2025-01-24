package dbs

import (
	"auth-server/conf"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Setup() {
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseSetting.User,
		conf.DatabaseSetting.Password,
		conf.DatabaseSetting.Host,
		conf.DatabaseSetting.Name)))

	if err != nil {
		panic("failed to connect database")
	}
}
