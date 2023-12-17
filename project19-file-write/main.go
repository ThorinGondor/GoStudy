package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * 本章主要讲文件的写入
 */
func main() {
	file, err := os.OpenFile("./file/readme", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file error: %v", err)
		return
	}

	// 带缓冲的写入操作（需要Flush操作）
	writer := bufio.NewWriter(file)

	// 延迟执行，关闭文件流
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("close file error: %v", err)
		}
	}(file)

	// writer 开始写入
	_, err = writer.WriteString("This is GTX760\n")
	if err != nil {
		fmt.Printf("write file error: %v", err)
		return
	}
	_, err = writer.WriteString("This is GTX1080\n")
	if err != nil {
		fmt.Printf("write file error: %v", err)
		return
	}
	_, err = writer.WriteString("This is RTX4090\n")
	if err != nil {
		fmt.Printf("write file error: %v", err)
		return
	}

	// writer 一定要记得 flush！否则写入无效！
	err = writer.Flush()
	if err != nil {
		fmt.Printf("flush data to file error: %v", err)
		return
	}
}
