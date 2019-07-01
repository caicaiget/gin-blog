package main

import (
	"fmt"
	_ "gin-blog/docs"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://erry.io
// @license.name caicai
// @license.url https://erry.io
// @host 127.0.0.1:8000
func main() {

	gin.SetMode(setting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ReadTimeout
	writeTimeout := setting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()

	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
