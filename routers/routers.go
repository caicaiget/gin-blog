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
	apiv1 := r.Group("/api/v1")
	//apiv2 := r.Group("/api/v1")
	//apiv2.GET("/articles", v1.GetArticles)
	apiv1.Use(jwt.JWT())
	//binding.Validator = &util.MyValidator{}
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.POST("/articles", v1.CreateArticle)
		apiv1.PUT("/articles/:id", v1.UpdateArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
