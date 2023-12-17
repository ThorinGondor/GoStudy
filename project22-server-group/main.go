package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// gf run main.go 启动项目
// 本章讲述分组路由 group，我们前面注册路由时每次都手动写入前缀 /api，这非常麻烦且不便维护，因此使用 group 方法可以统一添加前缀
// group 方法包在路由注册的外层即可
func main() {
	server := g.Server("server1")
	server.Group("/api", func(group *ghttp.RouterGroup) {
		group.GET("/query", func(request *ghttp.Request) {
			request.Response.Write("query success")
		})
		group.PUT("/create", func(request *ghttp.Request) {
			request.Response.Write("create success")
		})
	})
	server.SetPort(7750) // localhost:7750/api
	server.Run()

}
