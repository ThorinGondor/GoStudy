package main

import (
	"awesomeProject/functions"
	"awesomeProject/typeUsage"
	"fmt"
	"time"
)

// init 函数将会在 main 函数执行前就被调用
func init() {
	fmt.Println("init main package......")
}

func judgeDataType(param interface{}) string {
	// 使用 param.(type) 判断数据类型，必须要配合 switch 进行使用
	switch param.(type) {
	case int8:
		return "int 8"
	case int32:
		return "int 32"
	case float32:
		return "float 32"
	case time.Time:
		return "time"
	case bool:
		return "bool"
	default:
		return "unknown type"
	}
}

// Go 的函数支持可变参数
func addAll(args ...int32) int32 {
	var sum int32
	for _, value := range args {
		sum += value
	}
	return sum
}

/**
 * 本章主要讲：函数的多重用法，type 关键字的使用
 */
func main() {
	// function 返回多重参数
	var n1, n2 = 15.678, 45.92
	var r1, r2 = functions.PrintNumbers(n1, n2)
	fmt.Println(r1, r2)

	// type 关键字的使用
	// 可以使用 type 指定同一种函数类型，如下：
	typeUsage.TryType(45, 18)

	// 可以使用 type 指代数据类型，注意指代的类型和原数据类型不是同一种，需要转换
	type 字符串 string
	var str 字符串
	str = "test me"
	fmt.Printf("The type of 字符串 is %T, and value is %v\n", str, str)

	// 使用 param.(type) 判断数据类型，必须要配合 switch 进行使用
	var number float32 = 17.68
	var flag = true
	fmt.Println(judgeDataType(number))
	fmt.Println(judgeDataType(flag))

	// Go 的函数支持可变参数
	fmt.Println("ADD ALL: ", addAll(10, 20, 30, 40))

	// init 函数将会在 main 函数执行前就被调用
	// 如果导入其他包，而包里的源码也有 init 函数，则先调用其他包的 init 方法

	// 匿名函数，如果我们只希望调用该函数一次，则可以考虑使用匿名函数
	result := func(args ...int32) int32 {
		var sum int32 = 0
		for _, value := range args {
			sum += value
		}
		return sum
	}(10, 20)
	fmt.Println("anonymous function add result: ", result)

	// 但如果把 匿名函数 赋值给一个函数变量，就可以多次调用该匿名函数：
	anonymousFunction := func(args ...float32) float32 {
		var sum float32 = 0
		for _, v := range args {
			sum += v
		}
		return sum
	}
	fmt.Println("anonymous function add result: ", anonymousFunction(1.5, 6.34, 5.19))
}
