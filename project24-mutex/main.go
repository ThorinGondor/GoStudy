package main

import "sync"

func main() {

	var count = 0
	// 声明一个互斥锁（无需初始化）
	var mu sync.Mutex
	// 声明一个并发阻塞（无需初始化），确保 routine 执行完成
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			for j := 0; j < 10000; j++ {
				count++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()

}
