package util

import (
	"gin-blog/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) (int, int) {
	pageOffset := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	pageSize, _ := com.StrTo(c.DefaultQuery("pageSize", strconv.Itoa(setting.PageSize))).Int()
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}

	return pageOffset, pageSize
}
