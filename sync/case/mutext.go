package _case

import "C"
import (
	"fmt"
	"sync"
)

func MutexCase() {
	//singleRoutine()
	//multipleRoutine()
	//multipleSafeRoutine()
	multipleSafeRoutineByRWMutex()
}

// 单协程操作
func singleRoutine() {
	mp := make(map[string]int)
	list := []string{"A", "B", "C", "D"}
	for i := 0; i < 20; i++ {
		for _, item := range list {
			_, ok := mp[item]
			if !ok {
				mp[item] = 1
			} else {
				mp[item] += 1
			}
		}
	}
	fmt.Println(mp)
}

// 多协程操作,非协程安全,会奔溃
func multipleRoutine() {
	mp := make(map[string]int)
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, item := range list {
				_, ok := mp[item]
				if !ok {
					mp[item] = 1
				} else {
					mp[item] += 1
				}
			}
		}()

	}
	wg.Wait()
	fmt.Println(mp)
}

// 多协程操作,协程安全
func multipleSafeRoutine() {
	type safeMap struct {
		data map[string]int
		sync.Mutex
	}

	mp := safeMap{
		data:  make(map[string]int),
		Mutex: sync.Mutex{},
	}

	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mp.Lock()
			defer mp.Unlock()

			for _, item := range list {
				_, ok := mp.data[item]
				if !ok {
					mp.data[item] = 1
				} else {
					mp.data[item] += 1
				}
			}
		}()

	}
	wg.Wait()
	fmt.Println(mp.data)
}

type cache struct {
	data map[string]string
	sync.RWMutex
}

func (c *cache) Get(key string) string {
	c.RLock()         //加读锁
	defer c.RUnlock() //关闭读锁
	value, ok := c.data[key]
	if !ok {
		return ""
	}
	return value
}
func (c *cache) Set(key string, value string) {
	c.Lock()         //加写锁
	defer c.Unlock() //关闭写锁
	c.data[key] = value
}

func newCache() *cache {
	p := &cache{
		data:    make(map[string]string),
		RWMutex: sync.RWMutex{},
	}
	return p
}

// 读写锁.只对写互斥
func multipleSafeRoutineByRWMutex() {
	c := newCache()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Set("name", "nick")
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(c.Get("name"))
		}()
	}
	wg.Wait()
}
