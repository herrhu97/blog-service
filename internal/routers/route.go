package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	article := v1.NewArticle()
	tag := v1.NewTag()
	// v1 获取

	apiv1 := r.Group("/api/v1")
	apiv1.Use() //middleware.JWT()
	{
		// 创建标签
		apiv1.POST("/tags", tag.Create)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiv1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiv1.GET("/tags", tag.List)

		// 创建文章
		apiv1.POST("/articles", article.Create)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiv1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiv1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiv1.GET("/articles", article.List)
	}

	return r
}
