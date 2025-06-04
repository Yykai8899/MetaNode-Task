package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"task-go/task-go/go_base_3/constant"
)

type students struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Grade string `gorm:"column:grade"`
}

type accounts struct {
	ID      uint    `gorm:"primaryKey"`
	Balance float64 `gorm:"column:balance"`
}

type transactions struct {
	ID            uint    `gorm:"primaryKey"`
	FromAccountId uint    `gorm:"column:from_account_id"`
	ToAccountId   uint    `gorm:"column:to_account_id"`
	Amount        float64 `gorm:"column:amount"`
}

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db := ConnectDB()
	err := db.AutoMigrate(&students{}, &accounts{}, &transactions{}, &employee{}, &book{}, &User{}, &Post{}, &Comment{})
	if err != nil {
		panic(err)
	}
	return db
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(constant.DBPATH), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	// 初始化数据库
	db := InitDB()
	err := db.AutoMigrate(&students{})
	if err != nil {
		return
	}
	// 题目一
	//student(db)

	// 题目二
	//transaction(db)

	// sqlx题目1
	//sqlx1(db)

	// sqlx题目2
	//sqlx2(db)

	// gorm进阶题目1
	grom1(db)
}

func student(db *gorm.DB) {
	db.Create(&students{Name: "张三", Age: 20, Grade: "三年级"})

	var students []students
	db.Find(&students)

	// 查询数据
	db.Where("age > ?", 18).Find(&students)
	fmt.Println("年龄大于18的的学生：", students)

	db.Model(&students).Where("name = ?", "张三").Update("grade", "四年级")

	db.Where("name = ?", "张三").Find(&students)
	fmt.Println("张三年级修改后：", students)

	db.Where("age < ?", 15).Delete(&students)
	var delStudent = students
	db.Unscoped().Where("age = ?", 15).Find(&delStudent)
	fmt.Println("删除小于15岁的同学:", delStudent)
}

func transaction(db *gorm.DB) {
	// 转账 account2 -> account1 100  失败
	msg, ok := transferMoney(2, 1, 100, db)
	fmt.Println(msg, ok)

	// 转账 account1 -> account2 100  成功
	msg1, ok1 := transferMoney(1, 2, 100, db)
	fmt.Println(msg1, ok1)

}

func transferMoney(fromID, toID uint, money float64, db *gorm.DB) (string, bool) {

	var tMsg string
	var fromAccount, toAccount accounts

	err := db.Transaction(func(tx *gorm.DB) error {
		// 加锁查询
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Take(&fromAccount, fromID).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 加锁失败.", fromID)
			return err
		}

		// 检查余额
		if fromAccount.Balance < money {
			tMsg = fmt.Sprintf("账户: %d 余额不足.", fromID)
			return errors.New(fmt.Sprintf("用户%d 余额不足", fromAccount.ID))
		}

		// 扣款
		if err := tx.Model(&fromAccount).
			Update("balance", gorm.Expr("balance - ?", money)).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 扣款失败.", fromID)
			return err
		}

		// 转账
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Take(&toAccount, toID).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 加锁失败.", toID)
			return err
		}

		if err := tx.Model(&toAccount).
			Update("balance", gorm.Expr("balance + ?", money)).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 加款失败.", toID)
			return err
		}
		return nil
	})
	if err != nil {
		return tMsg, false
	}
	return "转账成功", true
}

type employee struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"column:name"`
	Department string  `gorm:"column:department"`
	Salary     float64 `gorm:"column:salary"`
}

// sqlx题目1
func sqlx1(db *gorm.DB) {
	// 写入测试数据
	//empTest := []employee{{Name: "小明", Department: "技术部"}, {Name: "小李", Department: "企划部"}, {Name: "小红", Department: "采购部"}, {Name: "小白", Department: "人事部"}}

	//db.Create(&empTest)

	var employees []employee
	db.Raw("select id,name,department,salary from employees where department = ?", "技术部").Scan(&employees)
	fmt.Println("所有技术部人员:", employees)

	var employeesMaxFee []employee
	db.Raw("select id,name,department,MAX(salary) as salary from employees").Scan(&employeesMaxFee)
	fmt.Println("工资最高人员:", employeesMaxFee)
}

type book struct {
	ID     uint    `gorm:"primaryKey"`
	Title  string  `gorm:"column:title"`
	Author string  `gorm:"column:author"`
	Price  float64 `gorm:"column:price"`
}

func sqlx2(db *gorm.DB) {
	var books []book
	// 测试数据
	for i := 0; i < 10; i++ {
		book := book{Title: "标题" + strconv.Itoa(i), Author: "作者" + strconv.Itoa(i), Price: float64(45 + i)}
		books = append(books, book)
	}
	db.Create(&books)
	var greaterThan50Book []book
	db.Raw("select title,author,price from books where price > ?", 50).Scan(&greaterThan50Book)
	fmt.Println("书本大于50元列表:", greaterThan50Book)
}

type User struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Posts []Post `gorm:"foreignKey:id;references:post_id"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"column:title"`
	post     string    `gorm:"column:post"`
	UserId   uint      `gorm:"column:user_id"`
	Comments []Comment `gorm:"foreignKey:id;references:comment_id"`
}

type Comment struct {
	gorm.Model
	Context string `gorm:"column:context"`
	PostID  uint   `gorm:"column:post_id"`
}

func grom1(db *gorm.DB) {
	// 测试数据
	userDB := User{
		Name: "王五",
		Posts: []Post{
			{
				Title:    "文章6",
				Comments: []Comment{{Context: "评论1111"}},
			},
			{
				Title:    "文章7",
				Comments: []Comment{{Context: "评论3333"}},
			},
		},
	}
	db.Create(&userDB)
	//db.Create(&Post{
	//	Title:  "文章7",
	//	UserId: 3,
	//})

	//查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Preload("Posts.Comments").Take(&user, 3)
	fmt.Println(user)

	//查询评论数量最多的文章信息。
	var post Post

	sub := db.Model(&Comment{}).
		Select("post_id").
		Group("post_id").
		Order("COUNT(*) DESC").
		Limit(1)
	err := db.Preload("Comments").
		Where("id = (?)", sub).
		First(&post).Error
	if err != nil {
		panic(fmt.Sprintf("method3 error: %v", err))
	}
	fmt.Println(post)

	var comment Comment
	db.Take(&comment, 9)
	db.Clauses(clause.Returning{}).Delete(&comment)
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	//g更新用户的文章数量统计字段
	err = tx.Model(&User{}).Where("id = ?", p.UserId).UpdateColumn("post_count", gorm.Expr("post_count+1")).Error
	return
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	println("AfterDelete", c.ID, c.Context, c.PostID)
	var count int64
	err = tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("status", "无评论").Error
	}
	return nil
}
