package funcstion

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"task-go/task-go/go_base_3/constant"
)

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(constant.DBPATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
