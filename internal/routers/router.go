package routers

import (
	"blog-service/internal/middleware"
	v1 "blog-service/internal/routers/api/v1"

	_ "blog-service/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)             //新增标签
		apiv1.DELETE("/tags/:id", tag.Delete)       //删除指定标签
		apiv1.PUT("/tags/:id", tag.Update)          //指定更新标签
		apiv1.PATCH("/tages/:id/state", tag.Update) //指定更新部分状态
		apiv1.GET("/tags", tag.List)                //获取标签列表

		apiv1.POST("/articles", article.Create)            //新增文章
		apiv1.DELETE("/articles/:id", article.Delete)      //删除指定文章
		apiv1.PUT("/articles/:id", article.Update)         //指定更新文章内容
		apiv1.PATCH("/articles/:id/state", article.Update) //指定更新部分文章内容
		apiv1.GET("/articles/:id", article.Get)            //获取文章内容
		apiv1.GET("/articles", article.List)               //获取文章列表
	}

	return r
}
