package models

import (
	"io"
	"log"
	"text/template"
	"time"
)

// 处理好所有模板：让路由函数只需要填入数据即可

// Template结构体：为了之后扩充更多的属性方法等
type TemplateBlog struct {
	*template.Template
}

// 所有模板的集合
type HtmlTemplate struct {
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Home       TemplateBlog
	Index      TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

// 解析所有html为template并返回
// 分为三个步骤：0.为html创建承接的模板 1.获取当前html路径 2.额外函数填充
// 3.解析模板以及涉及的html 4.数据填入
// Input:1. 当前路径
// Output:全部已经加载好的模板
func InitTemplate(curDir string) HtmlTemplate {
	tbs := readTemplate( //传入的文件名需要小写：与在服务器中存储的文件名相同，为了后面需要路径处理
		[]string{"category", "custom", "detail", "home", "index", "login", "pigeonhole", "writing"},
		curDir,
	)
	// 返回HtmlTemplate
	htmlTemplate := HtmlTemplate{
		tbs[0], tbs[1], tbs[2], tbs[3], tbs[4], tbs[5], tbs[6], tbs[7],
	}
	return htmlTemplate
}

// 数据填入：模块化管理，让路由只需要调用这个api即可填入数据，不用接触原生代码
// 名字错了没关系，可以后面改，放心的写。写完了才是提高
// 因为是*TemplateBlog类型调用写入数据，整个路由接触的都是TemplateBlog类型
func (tb *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	tb.Execute(w, data)
}

// 处理模板
// Input：要处理的模板名数组; 当前文件路径
// Output：处理好的模板数组; 解析模板错误
// Q：为什么要单独拧出来：以防后面再添加新的模板页，只需要精准的修改本func即可
func readTemplate(viewNames []string, curDir string) []TemplateBlog {
	var tbs []TemplateBlog //返回
	home := curDir + "home.html"
	header := curDir + "layout/header.html"
	footer := curDir + "layout/footer.html"
	personal := curDir + "layout/personal.html"
	post := curDir + "layout/post-list.html"
	pagination := curDir + "layout/pagination.html"
	for _, v := range viewNames {
		viewName := v + ".html"
		t := template.New(viewName)
		// 2. 需要的函数填充
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date}) //传入函数
		// 3. 同时解析多个html文件
		t, err := t.ParseFiles(curDir+viewName, home, header, footer, pagination, personal, post)
		if err != nil {
			log.Println("解析模板出错：", err)
		}
		var templateBlog TemplateBlog
		templateBlog.Template = t
		tbs = append(tbs, templateBlog)
	}
	return tbs
}

// 处理首页的相关函数

type IndexData struct {
	Title string `json:title`
	Desc  string `json:describe`
}

// 待传入函数:判断偶数
func IsODD(num int) bool {
	return num%2 == 0
}

// 待传入函数:获取导航条中元素
func GetNextName(strs []string, num int) string {
	return strs[num+1] //返回的当前名字对应的url路径
}

// 待传入函数:获取当前时间并格式化
func Date(layout string) string {
	// 传入一个layout格式化时间，返回一个对应的当前时间的格式化形式
	return time.Now().Format(layout)
}
