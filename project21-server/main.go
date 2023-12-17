package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// gf run main.go 启动项目
// 本章讲述特殊的方法名，Init 和 Shut 不能作为自定义方法存在
func main() {
	server := g.Server("server1")

	ctrl := new(IndexController)

	server.BindObject("GET:/api", ctrl) // Init 方法 和 Shut 方法是特殊方法，不会被当做独立的路由来注册

	server.SetPort(6450) // localhost:6450/api/logic
	server.Run()

}

type IndexController struct {
}

// Init 是 特殊方法名，会被用于调用函数之前做一些前置操作
func (c IndexController) Init(req *ghttp.Request) {
	req.Response.Writeln("pre process")
}

// Logic 具体的业务方法
func (c IndexController) Logic(request *ghttp.Request) {
	request.Response.Writeln("processing...")
}

// Shut 是 特殊方法名，会被用于调用函数完成之后做一些后置操作
func (c IndexController) Shut(req *ghttp.Request) {
	req.Response.Writeln("post process")
}
