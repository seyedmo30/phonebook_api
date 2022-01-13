package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seyedmo30/phonebook_api/phonebook/models"
	"github.com/seyedmo30/phonebook_api/phonebook/routers"
	"github.com/seyedmo30/phonebook_api/pkg/setting"
)

func init() {
	setting.Setup()
	models.Setup()

}
func main() {

	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	// log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
