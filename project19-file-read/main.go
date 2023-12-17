package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 本章主要讲文件的读取，有两种方式
func main() {
	usingBufferReader()
	usingOsReader()
}

// usingOsReader 一次性读取全部内容，因此容易爆内存，适合小文本的读取
func usingOsReader() {
	file := "./file/readme"
	// 返回的是一个文件地址
	fmt.Println("file address: ", file)
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file error: ", err)
		return
	}
	fmt.Printf("%v", string(content))
}

// usingBufferReader 使用带缓冲的 reader 读取文件内容，适合读取大文本
func usingBufferReader() {
	// 打开一份文件
	file, err := os.Open("./file/readme")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	// 返回的是一个文件地址
	fmt.Println("file address: ", file)

	// 最后执行关闭文件流，否则可能导致资源泄露
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close file error: ", err)
		} else {
			fmt.Println("close file resource succeeded!")
		}
	}(file)

	// 创建一个 *Reader ，是带缓冲的
	//	/*
	//		const (
	//			defaultBufSize = 4096 //默认的缓冲区为 4096
	//		)
	//	*/
	reader := bufio.NewReader(file)

	fmt.Println("read start ===>")
	// 循环读取文件的内容
	for {
		content, err := reader.ReadString('\n')
		fmt.Printf(content)
		// 读到换行结束
		if err != nil && err != io.EOF {
			fmt.Println("read file error: ", err)
			break
		}
		if err == io.EOF {
			break
		}

	}
	fmt.Println("\n===> read end")
}
