package main

import (
	_case "basis/function/case"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	sum, err := _case.Sum(1, 2)
	fmt.Println(sum, err)
	//方法赋值给变量
	f1 := _case.Sum
	fmt.Println(f1(2, 3))
	f2 := _case.LogMiddleWare(_case.Sum)
	fmt.Println(f2(5, 6))
	f3 := _case.LogMiddleWare2(_case.Sum)
	fmt.Println(f3(7, 8))
	//再次附加中间
	f3 = _case.LogMiddleWare2(f3)
	fmt.Println(f3.Accumulation(1, 2, 3, 4, 5, 6))

	var f4 = _case.SumFunc(f2)
	fmt.Println(f4.Accumulation(1, 2, 3, 4, 5, 6, 7))

	fmt.Println(_case.Fib(10))

	//闭包的陷阱
	_case.ClosureTrap()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	<-ctx.Done()
}
