package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"shopping/config"
)

var DB *gorm.DB

//连接数据库
func InitDB() {
	mysqlSetting := config.GetSqlSetting()
	connStr := mysqlSetting.UserName + ":" +
		mysqlSetting.Password +
		"@/" + mysqlSetting.DataName + "?charset=utf8&parseTime=true&loc=Local"


	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	DB = db
}
func CloseDB() {
	DB.Close()
}
