package _case

import (
	"fmt"
	"time"
)

// 协程间通信
func Communication() {
	//定义的可读写的chan
	ch := make(chan int, 0)
	go communicationF1(ch)
	go communicationF2(ch)
}

// F1接收一个只写通道
func communicationF1(ch chan<- int) {
	//通过协程循环向通道写入0-99
	for i := 0; i < 100; i++ {
		ch <- i
	}

}

// F2接收一个只写通道
func communicationF2(ch <-chan int) {
	//通过协程循环向通道写入0-99
	for i := range ch {
		fmt.Println(i)
	}
}

// 并发场景下,同步机制
func ConcurrentSync() {
	//带缓冲的通道
	ch := make(chan int, 10)
	//向ch 写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	//向ch 写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	//从ch 读取消息并打印
	go func() {
		for i := range ch {
			fmt.Println(i)
			//time.Sleep(100 * time.Millisecond)
		}
	}()
}

// 通知协程退出与多路利用
func NoticeAndMultiplexing() {
	ch := make(chan int, 0)
	strCh := make(chan string, 0)
	done := make(chan struct{}, 0)

	go noticeAndMultiplexingF1(ch)
	go noticeAndMultiplexingF2(strCh)
	go noticeAndMultiplexingF3(ch, strCh, done)

	time.Sleep(5 * time.Second)
	close(done)

}

func noticeAndMultiplexingF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
func noticeAndMultiplexingF2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("数字: %d", i)
	}
}

// select 子句作为一个整体阻塞
func noticeAndMultiplexingF3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
	for {
		select {
		case _int := <-ch:
			fmt.Println(_int)

		case str := <-strCh:
			fmt.Println(str)
		case <-done:
			fmt.Println("收到退出通知,退出当前协程")
			return
		}
		i++
		fmt.Println("累计执行次数", i)
	}
}
