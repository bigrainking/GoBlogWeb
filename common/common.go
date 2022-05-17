package common

import (
	"BRK-go-Blog/config"
	"BRK-go-Blog/models"
)

// 通用文件
var Templates models.HtmlTemplate

// 加载所有模板文件
func LoadTemplate() {
	Templates = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
}
