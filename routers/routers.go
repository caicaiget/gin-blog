package routers

import (
	"gin-blog/middleware"
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.JsonAccept())

	gin.SetMode(setting.RunMode)
	r.POST("/api/auth", api.GetAuth)
	ginConfig := ginSwagger.Config{}
	ginConfig.URL = "doc.json"
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginConfig, swaggerFiles.Handler))
	unionApi := r.Group("/api/v1")
	{
		// 上传文件
		unionApi.POST("/upload", api.Upload)
		unionApi.GET("/:user/articles", v1.GetArticles)
		unionApi.GET("/:user/tags", v1.GetTags)
	}
	loginApi := r.Group("/api/v1")
	loginApi.Use(jwt.JWT())
	{
		//新建标签
		loginApi.POST("/tags", v1.AddTag)
		//更新指定标签
		loginApi.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		loginApi.DELETE("/tags/:id", v1.DeleteTag)
		loginApi.POST("/articles", v1.CreateArticle)
		loginApi.PUT("/articles/:id", v1.UpdateArticle)
		loginApi.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
