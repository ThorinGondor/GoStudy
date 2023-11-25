package main

import "fmt"

/**
 * 本章主要讲 defer 的使用
 * go 的 defer 相当于 java 的 finally 异常处理块，使用 defer 可以保证函数不论是否发生异常退出以后，都必须执行 defer 语句，从而保证资源正常释放
 * 注：catch 错误则需要 panic
 */

// 下面这个例子可以演示 defer 的执行顺序
func sum(args ...int32) int32 {

	// 执行顺序 OK3 -> OK2 -> OK1， 先执行正常语句，然后因为 defer 是栈，先进后出
	defer fmt.Println("OK1")
	defer fmt.Println("OK2")

	var sum int32 = 0

	for _, v := range args {
		sum += v
	}
	fmt.Println("OK3 Sum = ", sum)
	return sum

}

func main() {
	sum(1, 2, 3, 4)
}
