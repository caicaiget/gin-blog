package api

import (
	"gin-blog/pkg/setting"
	"github.com/b3log/gulu"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
)

func Upload(c *gin.Context)  {
	result := gulu.Ret.NewResult()
	result.Code = 200
	defer c.JSON(http.StatusOK, result)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Msg = err.Error()
		return
	}
	filename := header.Filename
	out, err := os.Create(path.Join(setting.StaticFilePath, filename))
	if err != nil {
		result.Code = http.StatusBadRequest
		result.Msg = err.Error()
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	result.Data = "success"
}
