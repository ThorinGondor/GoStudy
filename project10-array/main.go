package main

import "fmt"

func main() {
	// 数组创建后，如果没有赋值，有默认值。
	// 数值类型数组：默认值为 0
	// 字符串数组： 默认值为 “”
	// bool 数组： 默认值为 false

	var arr = [5]int{1, 2, 3}
	fmt.Println(arr)

	// ... 自动推断数组长度
	var arr1 = [...]int{2: 1, 3: 2, 7: 3}
	fmt.Println(arr1)

	var arr2 = [5]string{3: "Jason"}
	fmt.Println(arr2)

	for k, v := range arr1 {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
}
