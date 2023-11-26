package typeUsage

import "fmt"

// 下面两个函数类型完全相同
func add(n1 int32, n2 int32) int32 {
	return n1 + n2
}

func sub(n1 int32, n2 int32) int32 {
	return n1 - n2
}

// 用 type 指代同一种类型的函数，即用 type 定义函数类型
type calc func(n1, n2 int32) int32

// 创建通用函数，即可使用函数类型作为传入形参
func calculate(specificFun calc, n1 int32, n2 int32) int32 {
	return specificFun(n1, n2)
}

func TryType(param1, param2 int32) {

	// 传入 add 的函数类型
	var result = calculate(add, param1, param2)
	fmt.Println("ADD: ", result)

	// 传入 sub 的函数类型
	result = calculate(sub, param1, param2)
	fmt.Println("SUB: ", result)
}
