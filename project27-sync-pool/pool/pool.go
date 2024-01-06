package pool

import (
	"log"
	"math/rand"
	"sync"
)

/**
为了实现一个连接池，首先定义一个连接，然后连接池继承（匿名内置） Pool，然后重写 Pool 的 Get 和 Put 方法，从而从 连接池 获取/放回连接！
*/

const (
	ON  = 1
	OFF = 0
)

// Connection 一个可用连接（可放在连接池中）
type Connection struct {
	ID     int64
	Target string
	Status int
}

func (c *Connection) GetStatus() int {
	return c.Status
}

// NewConnection 创建一个连接
func NewConnection(target string) *Connection {
	return &Connection{
		ID:     rand.Int63(),
		Target: target,
		Status: ON,
	}
}

// ConnectionPool 连接池结构体定义
type ConnectionPool struct {
	sync.Pool
}

// InitConnectionPool 初始化连接池
func InitConnectionPool(target string) (*ConnectionPool, error) {
	return &ConnectionPool{
		Pool: sync.Pool{
			New: func() any {
				return NewConnection(target)
			},
		},
	}, nil
}

// Get 从连接池获取一个对象
func (p *ConnectionPool) Get() *Connection {
	connection := p.Pool.Get().(*Connection)
	if connection.GetStatus() == OFF {
		log.Println("No connection available, Create connection")
		connection = p.Pool.New().(*Connection)
	}
	return connection
}

// Put 将已使用完成的对象放回连接池
func (p *ConnectionPool) Put(connection *Connection) {
	if connection.GetStatus() == OFF {
		return
	} else {
		p.Pool.Put(connection)
	}
}
