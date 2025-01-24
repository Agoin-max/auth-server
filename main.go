package main

import (
	"auth-server/conf"
	"auth-server/dbs"
	"auth-server/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	conf.Setup()
	dbs.Setup()
}

// @title		Auth-Server API
// @version		1.0
// @description	Auth-server
func main() {
	gin.SetMode(conf.ServerSetting.RunMode)

	routersInit := router.InitRouter()
	readTimeout := conf.ServerSetting.ReadTimeout
	writeTimeout := conf.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", conf.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
