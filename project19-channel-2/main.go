package main

import "fmt"

/**
 * 本章主要 route 协程 & channel 管道通信
 * channel 线程安全，允许被同时多个协程进行操作
 */

func produceData(channel chan int) {
	for i := 1; i < 800; i++ {
		channel <- i
	}
	close(channel)
}

func consumeData(channel chan int, primeChannel chan int, exitChannel chan bool) {
	var flag bool
	for {
		v, ok := <-channel
		if ok {
			flag = true
			for i := 2; i < v; i++ {
				if v%i == 0 {
					flag = false
					break
				}
			}

		} else {
			break
		}
		if flag {
			primeChannel <- v
		}
	}
	fmt.Println("A prime channel ended")
	exitChannel <- true

}

func main() {
	dataChannel := make(chan int, 100)
	primeChannel := make(chan int, 200)
	exitChannel := make(chan bool, 4)
	go produceData(dataChannel)
	for i := 0; i < 4; i++ {
		go consumeData(dataChannel, primeChannel, exitChannel)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChannel
		}
		close(primeChannel)
	}()

	for {
		v, ok := <-primeChannel
		if ok {
			fmt.Println("素数：", v)
		} else {
			break
		}
	}

}
