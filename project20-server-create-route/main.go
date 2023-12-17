package main

import (
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/frame/g"
)
import "github.com/gogf/gf/v2/net/ghttp"

// gf run main.go 启动项目
// 构建了 web 服务，创建了多个 server 服务，使用了多种路由注册规则
func main() {
	server := g.Server("server1")
	server.BindHandler("/api", func(req *ghttp.Request) {
		req.Response.Write("server 1 status: up")
	})
	server.SetIndexFolder(true) // 默认不允许列出 server 主目录的静态资源，此处我们将其设置为允许
	server.SetPort(7970, 7990)  // 多端口监听 http://localhost:7970/api; http://localhost:7990/api
	_ = server.Start()

	// 我们还可以建立多个 server 并行
	server2 := g.Server("server2")
	// domain 方法绑定域名
	server2.Domain("127.0.0.1").BindHandler("/api", func(req *ghttp.Request) {
		req.Response.Write("server 2 status: up")
	})
	// 多种路由注册规则 详细规则见页尾
	server2.BindHandler("POST:/api/:controller/update", func(req *ghttp.Request) {
		req.Response.Writef("Server 2 %v %v %v", req.Router.Uri, req.Get("controller"), "update")
	})
	server2.BindHandler("GET:/api/:controller/query", func(req *ghttp.Request) {
		req.Response.WriteJson(req.Router)
	})

	// ！！！ 上述路由注册全都是 函数注册路由 还可以 对象方法注册路由 ！！！
	ctrl := &FluxController{number: gtype.NewInt(0)}
	server2.BindHandler("POST:/api/add", ctrl.Add)

	// ！！！ 还可以使用对象注册路由 ！！！此种注册方式将自动使用 对象的方法名 为请求 URL 尾部添加对应的 path
	c := new(WebController)
	server2.BindObject("GET:/api/web/", c) // http://127.0.0.1:8990/api/web/show http://127.0.0.1:8990/api/web/index
	// 如果结构体的方法名为驼峰，则默认全部转为小写，且用 "-" 连接 比如 UserList -> /user-list 若要改变转path的方法，请使用 ghttp.UriType

	server2.SetIndexFolder(true) // 默认不允许列出 server 主目录的静态资源，此处我们将其设置为允许
	server2.SetPort(8970, 8990)  // 多端口监听 http://127.0.0.1:8970/api; http://127.0.0.1:8990/api
	_ = server2.Start()

	// 若多个 server 并行，则需要 wait
	g.Wait()

}

// FluxController 结构体
type FluxController struct {
	// 某个静态资源
	number *gtype.Int
}

// Add 接口
func (c *FluxController) Add(req *ghttp.Request) {
	c.number.Add(1)
	req.Response.Write("total access flux: ", c.number)
}

type WebController struct {
}

func (wc WebController) Index(req *ghttp.Request) {
	req.Response.Write("Index Controller")
}

func (wc WebController) Show(req *ghttp.Request) {
	req.Response.Write("Show Controller")
}

/**
多种路由规则：
	rule: /user/:user

	/user/john                match
	/user/you                 match
	/user/john/profile        no match
	/user/                    no match

	rule: /src/*path

	/src/                     match
	/src/someFile.go          match
	/src/subDir/someFile.go   match
	/user/                    no match
	/user/john                no match

	rule: /order/list/{page}.php

	/order/list/1.php          match
	/order/list/666.php        match
	/order/list/2.php5         no match
	/order/list/1              no match
	/order/list                no match
*/
