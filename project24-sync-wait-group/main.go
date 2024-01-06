package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// WaitGroup 实现同步等待
	var wg = sync.WaitGroup{} // WaitGroup 必须使用 值类型，否则容易出现死锁
	var urls = []string{
		"https://www.baidu.com/",
		"https://www.bilibili.com/",
	}
	for _, url := range urls {
		wg.Add(1)             // 添加/减少 goroutine 的数量，正数添加，负数减少
		go func(url string) { // 开启一个协程
			defer wg.Done() // 相当于 Add(-1)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println(resp)
			}
		}(url)
	}
	wg.Wait() // 同步阻塞等待，直到 wg 的值减至 0
	// 若无 wg，我们的 go func 协程很可能还未执行完毕就因为 main 函数结束而被 kill 掉
}
