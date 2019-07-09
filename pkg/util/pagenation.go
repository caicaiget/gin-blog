package util

import (
	"gin-blog/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) (int64, int64) {
	var pageOffset int64 = 0
	page, _ := com.StrTo(c.Query("page")).Int64()
	pageSize, _ := com.StrTo(c.DefaultQuery("pageSize", strconv.Itoa(setting.PageSize))).Int64()
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}

	return pageOffset, pageSize
}
