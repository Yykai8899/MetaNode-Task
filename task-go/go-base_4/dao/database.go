package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"task-go/task-go/go_base_3/constant"
)

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	db, merr := MysqlDB()
	if merr != nil {
		liteDB, err := gorm.Open(sqlite.Open(constant.DBPATH), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		fmt.Println("已连接sqlLite数据库")
		return liteDB
	}
	fmt.Println("已连接mysql数据库")
	return db
}

// MysqlDB 连接mysql数据库
func MysqlDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	return db, err
}
