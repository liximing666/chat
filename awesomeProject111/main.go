package main

import (
	"awesomeProject111/bootstrap"
	"awesomeProject111/routes"
	"log"
	"net/http"
)

func main() {
	bootstrap.Init()
	server := routes.New()

	if err := http.ListenAndServe(server.Addr, server.Handler); err != nil {
		log.Println("http服务启动失败")
	}

}
