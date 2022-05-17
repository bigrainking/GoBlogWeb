package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

// 定义一个结构映射配置文件
type tomlConfig struct {
	// 配置文件中采用模块形式映射，每个模块中的内容独属于它:比如View、Config
	Viewer Viewer //使用大写来让外部可以访问； 为了让大写与配置文件中的块名字可以匹配需要一个依赖包
	System SystemConfig
}

var Cfg *tomlConfig

// 程序启动时 会执行init方法
func init() {
	// 获取配置文件中的内容
	Cfg = &tomlConfig{}

	// 将配置文件中没有的参数，现在手动配置
	Cfg.System.AppName = "BigRainKing"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir //当前文件夹

	//将配置文件中的内容读取到Cfg中
	_, err := toml.DecodeFile("config/config.toml", &Cfg) //此处的文件地址为当前文件夹config下面的config.toml文件
	if err != nil {
		panic(err)
	}
}
