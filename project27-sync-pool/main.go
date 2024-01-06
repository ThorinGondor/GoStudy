package main

import (
	"awesomeProject/project27-sync-pool/pool"
	"log"
	"sync"
)

/*
pool 是一个临时对象池，缓存被重复利用的对象，代替在堆上分配内存空间，以此减少 GC 的压力

pool 用于缓存创建成本比较高、使用比较频繁的对象，比如 HTTP/JDBC 连接池

pool 的长度默认为 CPU 的线程数 (12)

缓存于 pool 的对象可能在不被通知的情况下被回收
*/
func main() {
	// 创建一个连接池
	connectionPool, err := pool.InitConnectionPool("127.0.0.1/")
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 获取一个连接对象
			connection := connectionPool.Get()
			log.Println("Get Connection", connection.ID, connection.Status)
			// 放回一个连接对象
			connectionPool.Put(connection)
			log.Println("Put back Connection", connection.ID, connection.Status)
		}()
	}

	wg.Wait()

}
