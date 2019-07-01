package routers

import (
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	r.POST("/api/auth", api.GetAuth)
	ginConfig := ginSwagger.Config{}
	ginConfig.URL = "doc.json"
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginConfig, swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	r.Use(cors.Default())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
