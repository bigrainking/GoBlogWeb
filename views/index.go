package views

import (
	"BRK-go-Blog/common"
	"BRK-go-Blog/config"
	"BRK-go-Blog/models"
	"net/http"
)

// 用模板获取html页面
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) { //类下面的函数
	// 1. 模板解析
	index := common.Templates.Index //已经处理好的index模板

	// 假数据
	categorys := []models.Category{
		{1, "Go language", "20220512", "20220513"},
	}
	// 页面文章
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		}}
	//定义数据结构来填充页面上的内容
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     2,
		Page:      2,
		Pages:     []int{1, 2},
		PageEnd:   false,
	}

	// 2. 假数据填充
	index.WriteData(w, hr)
}
