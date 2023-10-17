package _case

import (
	"context"
	"fmt"
	"time"
)

func ContextCase() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "dest", "ContextCase")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	data := [][]int{{1, 2}, {3, 2}}
	ch := make(chan []int)
	go calculate(ctx, ch)
	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}

}

func calculate(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			fmt.Println(item)
			ctx := context.WithValue(ctx, "desc", "calculate")
			ch := make(chan []int)
			go sumContext(ctx, ch)
			ch <- item

			ch1 := make(chan []int)
			go multiContext(ctx, ch1)
			ch1 <- item
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("calculate协程退出", desc, ctx.Err())
			return
		}

	}
}
func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			fmt.Println(item)
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Println(a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("sumContext协程退出", desc, ctx.Err())
			return
		}

	}
}

func multiContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			fmt.Println(item)
			a, b := item[0], item[1]
			res := multi(a, b)
			fmt.Println(a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("multiContext协程退出", desc, ctx.Err())
			return
		}

	}
}
func sum(a, b int) int {
	return a + b
}
func multi(a, b int) int {
	return a * b
}

// GPT例子
func GPT例子() {
	// 创建一个根上下文
	ctx := context.Background()

	// 根上下文可以用来取消整个任务
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// 子操作上下文，超时时间为 3 秒钟
	ctx, timeout := context.WithTimeout(ctx, 3*time.Second)
	defer timeout()
	// 子协程执行
	go func(ctx context.Context) {
		// 模拟任务需要耗时 5 秒钟
		//time.Sleep(5 * time.Second)

		// 判断任务是否需要取消
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			fmt.Println("Task completed")
		}
	}(ctx)

	// 等待 4 秒钟
	time.Sleep(6 * time.Second)

	// 取消任务
	cancel()

	// 等待子协程执行完毕或超时
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout exceeded")
	}
}
