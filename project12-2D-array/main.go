package main

import "fmt"

/*
*
本章主要讲解 go 的二维数组
*/
func main() {
	var array [5][10]string
	array[3][6] = "Kimi"
	array[4][2] = "Jason"
	fmt.Println(array)

	// 二维数组遍历
	for k, v := range array {
		for k1, v1 := range v {
			var value string
			if v1 != "" {
				value = v1
			} else {
				value = "/"
			}
			fmt.Println("layer = ", k, " index = ", k1, " val = ", value)
		}
	}
}
