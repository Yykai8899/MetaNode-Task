package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task-go/task-go/go-base_4/dao"
	"time"
)

type Comment struct {
	CommentId int       `json:"commentId" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Blog      Blog      `json:"blog" gorm:"foreignKey:BlogID"`
	BlogID    int       `json:"blogId" gorm:"index"` // 为BlogID创建索引，优化查询性能
	UserID    int       `json:"userId"`
	UserName  string    `json:"userName"` // 用于存储User的外键
	Content   string    `json:"content" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func CommentInitDB() *gorm.DB {
	db := dao.ConnectDB()
	err := db.Debug().AutoMigrate(&Comment{})
	if err != nil {
		panic(err)
	}
	return db
}

// 新增评论
func CreateComment(comment *Comment) (err error) {
	fmt.Println("新增评论:", comment)
	err = CommentInitDB().Debug().Create(&comment).Error
	if err != nil {
		return errors.New("create comment error")
	}
	return nil
}

// 获取评论列表
func GetComment(id int, commentList *[]Comment) (err error) {
	err = CommentInitDB().Debug().Where("blog_id=?", id).Find(&commentList).Error
	if err != nil {
		return errors.New("get comment error")
	}
	return nil
}

// 删除评论
func DelComment(idiot int) (err error) {
	fmt.Printf("id = %v", idiot)
	err = CommentInitDB().Debug().Where("comment_id=?", idiot).Delete(&Comment{}).Error
	if err != nil {
		return errors.New("delete comment error")
	}
	return nil
}
