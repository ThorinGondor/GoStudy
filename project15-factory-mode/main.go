package main

import (
	"awesomeProject/project15-factory-mode/pojo"
	"fmt"
)

// 本章主要讲如何使用工厂模式构建对象，尤其在结构体是小写开头不允许直接包外创建的时候，可以通过这种工厂模式解决问题
func main() {
	// 返回值类型
	var student0 = pojo.StudentFactory(3424398, "Burton")
	fmt.Println(student0)

	// 返回引用类型
	var student1 = pojo.StudentAddressFactory(34724312, "Maxy")
	fmt.Println(*student1)
}
