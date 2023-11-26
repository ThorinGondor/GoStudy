package main

// 导入其他包的代码
import (
	"awesomeProject/user"
	"fmt"
	"ms/dao"
	"ms/view"
)

func main() {
	var name string = "Kimi"
	// 执行同一模块内 其他包的函数
	user.PrintUsername(name)

	// 执行外部模块/依赖的函数
	var result = dao.Select(1)
	fmt.Println("DAO select: ", result)
	result = view.View()
	fmt.Println("VIEW execute: ", result)
}
