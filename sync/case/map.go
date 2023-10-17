package _case

import (
	"fmt"
	"sync"
)

func MapCase() {
	mp := sync.Map{}
	//设置键值对
	mp.Store("name", "nick")
	mp.Store("email", "nick@163.com")
	//通过key获取value,如果不存在,则返回nil,ok,否则返回false
	fmt.Println(mp.Load("name"))
	fmt.Println(mp.Load("email"))
	fmt.Println(mp.Load("email111"))
	//通过key获取value,如果不存在,则指定的value,并返回
	//ok为true,表示key存在并返回值,为false表示key不存在并设置后返回
	fmt.Println(mp.LoadOrStore("hobby", "篮球"))
	fmt.Println(mp.LoadOrStore("hobby", "羽毛球")) //这次ok肯定是true
	//根据key获取value后,删除该key
	//ok为true表示key存在,为false表示key不存在
	fmt.Println(mp.LoadAndDelete("hobby"))
	fmt.Println(mp.LoadAndDelete("hobby"))
	//遍历
	mp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

}

// 有问题的.还是推荐用Mutex加Map
func MapCase1() {

	mp := sync.Map{}
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, item := range list {
				value, ok := mp.Load(item)
				if !ok {
					mp.Store(item, 1)
				} else {
					val := value.(int)
					val += 1
					mp.Store(item, val)
				}
			}
		}()
	}
	wg.Wait()
	mp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
