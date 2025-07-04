package models

import (
	"errors"
	"gorm.io/gorm"
	"task-go/task-go/go-base_4/dao"
	"time"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

// Blog 定义博客结构体
type Blog struct {
	BlogId    int       `form:"blogId" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Title     string    `form:"title" gorm:"type:varchar(255)"`
	Content   string    `form:"content" gorm:"type:text"`
	User      User      `form:"user" gorm:"foreignKey:UserId"` // 注意调整字段名和外键
	UserId    uint      `form:"userId"`
	UserName  string    `form:"userName"` // 用于存储User的外键
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func BlogInitDB() *gorm.DB {
	db := dao.ConnectDB()
	err := db.AutoMigrate(&Blog{})
	if err != nil {
		panic(err)
	}
	return db
}

func CreateBlog(blog *Blog) (err error) {
	db := BlogInitDB()
	var user User
	db.Where("id = ?", blog.UserId).First(&user)
	blog.User = user
	// 根据blog中的内容新建信息
	err = db.Create(&blog).Error
	if err != nil {
		return errors.New("create blog error")
	}
	return nil
}

// 修改博客
func UpdateBlog(blogId int, blog *Blog) (err error) {
	// 根据ID更新
	err = BlogInitDB().Debug().Model(&blog).Where("blog_id=?", blogId).Updates(map[string]interface{}{
		"Title":   blog.Title,
		"Content": blog.Content,
	}).Error
	if err != nil {
		return errors.New("update blog error")
	}
	return nil
}

// 删除博客
func DelBlog(blogId int) (err error) {
	// 根据blog中的内容删除blog
	err = BlogInitDB().Debug().Where("blog_id=?", blogId).Delete(&Blog{}).Error
	if err != nil {
		return errors.New("delete blog error")
	}
	return nil
}

// 获取所有博客
func GetAllBlog(blogList *[]Blog) (err error) {
	// 从数据库中读取所有的blog
	err = BlogInitDB().Debug().Find(&blogList).Error
	if err != nil {
		return errors.New("read blog error")
	}
	return nil
}

// 获取单个
func GetABlog(blogId int) (blog *Blog, err error) {
	blog = new(Blog)
	// 从数据库中读取特定的blog
	err = BlogInitDB().Debug().Where("blog_id=?", blogId).First(blog).Error
	if err != nil {
		return nil, errors.New("read blog error")
	}
	return
}

// 搜索博客
func SearchBlog(query string) (blogList []Blog, err error) {
	// 查询包含指定关键词的博客
	err = BlogInitDB().Debug().Where("content LIKE ? OR title LIKE ?", "%"+query+"%", "%"+query+"%").Find(&blogList).Error
	if err != nil {
		return nil, err
	}
	return blogList, nil
}
