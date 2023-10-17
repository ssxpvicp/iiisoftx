package _case

import (
	"fmt"
	"sync"
	"time"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})

	go readList(&list, cond)
	go readList(&list, cond)
	go readList(&list, cond)

	time.Sleep(3 * time.Second)
	initList(&list, cond)
}
func initList(list *[]int, c *sync.Cond) {
	//主叫方,可以持锁,也可以不持锁
	c.L.Lock()
	defer c.L.Unlock()
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	//唤醒所有条件等待的协程
	c.Broadcast()
	//唤醒一个
	//c.Signal()
}

func readList(list *[]int, c *sync.Cond) {
	//被叫方,必须持锁
	c.L.Lock()
	defer c.L.Unlock()
	fmt.Println("我在延时中,有锁,所以只能单个进行")
	time.Sleep(time.Second * 3)
	for len(*list) == 0 {
		fmt.Println("readlist wait")
		c.Wait()
	}
	fmt.Println("list数据", *list)
}

type queue struct {
	list []int
	cond *sync.Cond
}

func newQueue() *queue {
	q := &queue{
		list: []int{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
	return q
}
func (q *queue) Put(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.list = append(q.list, item)
	//当数据写入成功以后,唤醒一个协程来处理数据
	q.cond.Signal()
}
func (q *queue) GetMany(n int) []int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.list) < n {
		//取的数据不够.就行等待
		q.cond.Wait()
	}
	list := q.list[:n]
	q.list = q.list[n:]
	return list
}
func CondQueueCase() {
	q := newQueue()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			list := q.GetMany(n)
			fmt.Printf("%d:%d \n", n, list)
		}(i)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}

	wg.Wait()

}
