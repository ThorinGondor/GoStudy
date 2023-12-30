package main

import "fmt"

/**
 * 本章主要 route 协程 & channel 管道通信
 * channel 线程安全，允许被同时多个协程进行操作
 */
func main() {
	// route 协程操作的管道
	channel := make(chan int, 5)
	// 标记退出的管道
	exitChannel := make(chan bool, 1)

	// go 开启协程
	go writeData(channel)
	go readData(channel, exitChannel)

	// 阻塞读，防止协程还未完成就退出程序
	<-exitChannel
}

func writeData(channel chan int) {
	for i := 1; i <= 50; i++ {
		fmt.Println("send into channel: ", i)
		channel <- i
	}
	close(channel)
}

func readData(channel chan int, exitChannel chan bool) {
	for {
		val, ok := <-channel
		if ok {
			fmt.Println("get from channel: ", val)
		} else {
			fmt.Println("cannot get from channel, result is: ", ok)
			break
		}
	}
	exitChannel <- true
}
