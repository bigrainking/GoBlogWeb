package main

import (
	"BRK-go-Blog/common"
	"BRK-go-Blog/router"
	"log"
	"net/http"
)

func init() {
	// 加载好所有template的html
	common.LoadTemplate()
}
func main() {
	// 创建一个服务器
	server := http.Server{
		Addr: "127.0.0.1:8080", //本服务器的地址
	}

	// 路由
	router.Router()
	// 启动并且监听服务器的端口：看	是否有client访问
	if err := server.ListenAndServe(); err != nil { // 如果监听到错误，则不启动
		log.Println("出现问题", err)
	}
}
