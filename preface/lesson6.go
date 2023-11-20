package main

import (
	"fmt"
)

func main() {
	// for 循环
	var i = 1
	for i < 10 {
		fmt.Printf("iterator: %d \n", i)
		i++
	}

	var str = "Hello go!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println()

	// 使用 for range 方式
	for index, val := range str {
		fmt.Printf("%d, %c \n", index, val)
	}
}
