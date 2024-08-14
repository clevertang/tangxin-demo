package main

import (
	"fmt"
	"log"
	"net/http"
	"tangxin-demo/pkg/logging"
	"tangxin-demo/pkg/setting"
	"tangxin-demo/routers"
)

func main() {
	setting.Setup()
	logging.Setup()
	addr := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	router := routers.InitRouter()
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
