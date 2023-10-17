package _case

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCase() {
	var count int64 = 5
	//对原子操作的变量设置值
	atomic.StoreInt64(&count, 10)
	//对原子操作的变量取出值
	atomic.LoadInt64(&count)
	fmt.Println(count)
	//对原子操作的变量,累加
	atomic.AddInt64(&count, 10)
	fmt.Println(count) //累加结果20
	//对原子操作的变量交换操作.意思.就是设置新值,返回旧值
	atomic.SwapInt64(&count, 30)
	fmt.Println(count)
	//对原子操作的变量,先比较再交换,如果old参数和当前变量的值不一致则不交换,反之交换
	fmt.Println(atomic.CompareAndSwapInt64(&count, 30, 40))
	fmt.Println(count)

}

type atomicCounter struct {
	count int64
}

func (a *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.count)
}

func (a *atomicCounter) Inc() {
	atomic.AddInt64(&a.count, 1)

}

func AtomicCase1() {
	var count atomicCounter
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
}
func AtomicCase2() {
	list := []string{"A", "B", "C", "D"}
	//定义一个原子值
	var atomicMp atomic.Value
	//定义一个集合
	mp := make(map[string]int)
	//
	atomicMp.Store(&mp)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		atomicLabel:
			m := atomicMp.Load().(*map[string]int)
			m1 := make(map[string]int)
			for k, v := range *m {
				m1[k] = v
			}
			for _, item := range list {

				_, ok := m1[item]
				if !ok {
					m1[item] = 0
				}
				m1[item] += 1
			}
			bool := atomicMp.CompareAndSwap(m, &m1)
			if !bool {
				goto atomicLabel

			}
		}()
	}
	wg.Wait()
	fmt.Println(*(atomicMp.Load().(*map[string]int)))
}
