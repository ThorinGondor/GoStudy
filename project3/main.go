package main

import "fmt"

/**
 * 本章主要讲 闭包 的使用
 *  其实就是用一个专门的函数来创建并返回一个指定函数 f，然后将该返回的函数 f 执行若干次
 */

// 使用该函数创建并返回一个函数
func lazySum(arr []int) func() int {
	var count = 0
	var function = func() int {
		count++
		fmt.Printf("第 %d次执行函数 \n", count)
		var sum = 0
		for _, v := range arr {
			sum += v
		}
		fmt.Printf("结果为 %d \n", sum)
		return sum
	}
	return function
}

func main() {
	// 创建函数
	var function = lazySum([]int{1, 2, 3, 4})
	// 函数执行三次
	function()
	function()
	function()
}
