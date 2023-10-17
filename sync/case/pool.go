package _case

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

func PoolCase() {
	target := "192.168.239.149"
	pool, err := GetPool(target)
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				conn := pool.Get()
				fmt.Println(conn.ID)
				pool.Put(conn)
			}
		}()
	}
	wg.Wait()
}

const (
	NO  = 1
	OFF = 0
)

type Conn struct {
	ID     int64
	Target string
	Status int
}

func NewConn(target string) *Conn {
	return &Conn{
		ID:     rand.Int63(),
		Target: target,
		Status: NO,
	}
}
func (c *Conn) Getstatus() int {
	return c.Status
}

type ConnPool struct {
	sync.Pool
}

func (c *ConnPool) Get() *Conn {
	conn := c.Pool.Get().(*Conn)
	if conn.Status == OFF {
		conn = c.Pool.New().(*Conn)
	}
	return conn
}

func (c *ConnPool) Put(conn *Conn) {
	if conn.Getstatus() == OFF {
		return
	}
	c.Pool.Put(conn)
}

func GetPool(target string) (*ConnPool, error) {
	return &ConnPool{
		Pool: sync.Pool{
			New: func() any {
				return NewConn(target)
			},
		},
	}, nil
}
