package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"task-go/task-go/go-base_4/models"
	"task-go/task-go/go-base_4/response"
)

// CreateBlogHandler 创建博客
func CreateBlogHandler(c *gin.Context) {

	var blog models.Blog

	err := c.ShouldBind(&blog)

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println("新增博客输入信息:", blog)

	err = models.CreateBlog(&blog)

	if err != nil {
		response.FailWithMsg(c, "blog create fail")
	} else {
		response.OkWithData(c, blog)
	}
}

// UpdateBlogHandler 更新博客
func UpdateBlogHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		response.FailWithMsg(c, "id not found")
	}
	idiot, _ := strconv.Atoi(id)
	var blog models.Blog

	err := c.ShouldBind(&blog)
	if err != nil {
		println(err.Error())
		return
	}
	err = models.UpdateBlog(idiot, &blog)
	if err != nil {
		response.FailWithMsg(c, "blog update fail")
	} else {
		response.OkWithData(c, gin.H{
			"blog": blog,
		})
	}
}

// DeleteBlogHandler 删除博客
func DeleteBlogHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		response.FailWithMsg(c, "id not found")
	}
	idiot, _ := strconv.Atoi(id)

	err := models.DelBlog(idiot)
	if err != nil {
		response.FailWithMsg(c, "blog delete fail")
	} else {
		response.OkWithMsg(c, "删除成功")
	}
}

// GetAllBlogsHandler 查看所有博客
func GetAllBlogsHandler(c *gin.Context) {
	var blogs []models.Blog

	err := models.GetAllBlog(&blogs)
	if err != nil {
		response.FailWithMsg(c, "blog get fail")
	} else {
		response.OkWithData(c, blogs)
	}
}

// GetBlogByIdHandler 查看单个博客
func GetBlogByIdHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		response.FailWithMsg(c, "id not found")
	}
	idiot, _ := strconv.Atoi(id)
	// 从数据库中读取所有博客
	blog, err := models.GetABlog(idiot)
	if err != nil {
		response.FailWithMsg(c, "blog get fail")
	} else {
		response.OkWithData(c, blog)
	}
}

// SearchBlogsHandler 博客搜索
func SearchBlogsHandler(c *gin.Context) {
	query, ok := c.Params.Get("query")
	if !ok {
		response.FailWithMsg(c, "query not found")
	}
	blogList, err := models.SearchBlog(query)
	if err != nil {
		response.FailWithMsg(c, "blog search fail")
	} else {
		response.OkWithData(c, blogList)
	}
}
