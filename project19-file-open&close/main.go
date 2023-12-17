package main

import (
	"fmt"
	"os"
)

// 本章主要讲文件的操作：打开一份文件，关闭文件
func main() {
	// 打开一份文件
	file, err := os.Open("./file/readme")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	// 返回的是一个文件地址
	fmt.Println(file)

	// 关闭文件流，否则可能导致资源泄露
	err = file.Close()
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

}
