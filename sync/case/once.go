package _case

import (
	"fmt"
	"sync"
)

type onceMap struct {
	sync.Once
	data map[string]int
}

func (m *onceMap) LoadData() {
	m.Once.Do(func() {
		//这个方法内部,无论多个线程操作,只有一个线程会执行
		list := []string{"A", "B", "C", "D"}
		for _, item := range list {
			_, ok := m.data[item]
			if !ok {
				m.data[item] = 0
			}
			m.data[item] += 1
		}
	})
}
func OnceCase() {
	o := &onceMap{
		Once: sync.Once{},
		data: make(map[string]int),
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.LoadData()
		}()
	}
	wg.Wait()
	fmt.Println(o.data)
}
