package main

import "fmt"

/**
 * 本章主要讲 string 的多重方法
 */

func main() {
	// 创建切片 元素长度为10 容量为40
	slice := make([]int, 10, 40)
	fmt.Println(slice)

	// 创建 map
	hashMap := make(map[string]int)
	fmt.Println(hashMap)

	// 创建通道 channel
	ch := make(chan int)
	fmt.Println(ch)
}
