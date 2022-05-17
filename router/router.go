package router

// 注册路由器

import (
	"BRK-go-Blog/api"
	"BRK-go-Blog/views"
	"net/http"
)

// 将路由集合起来
func Router() {
	// 页面：主页index路由
	http.HandleFunc("/", views.HTML.Index)
	// 数据：文章post路由
	http.HandleFunc("/api/v1/post", api.API.Post)
	// 静态资源不能让他走index这个路由，下面我们需要重新创建一个属于静态资源的路由
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}

// 路由器处理函数在其他模块
// 为了将三种不同类型的路由进行区分，将不同类型的路由处理函数交给专门的模块处理
// 1. Views 页面
// 2. api 数据 json
// 3. 静态数据
