package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
	"tangxin-demo/models"
	"tangxin-demo/pkg/logging"
	"tangxin-demo/pkg/setting"
	"tangxin-demo/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)

	log.Printf("Starting server on %s", endPoint)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
		log.Printf("Listening on address: %s", add)
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server error: %v", err)
	}
}
