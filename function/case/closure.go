package _case

import (
	"fmt"
	"log"
)

//斐波那契数列 ,X0+X1=X2...求Xn的值

func tool() func() int {
	var X0 = 0
	var X1 = 1
	var X2 = 0

	return func() int {
		X2 = X0 + X1
		X0 = X1
		X1 = X2
		return X2
	}

}

func Fib(n int) int {
	if n <= 2 {
		log.Fatal("请选择大于2的位置的")
	}
	t := tool()
	var res int
	for i := 0; i < n-2; i++ {
		res = t()
	}
	return res
}

func ClosureTrap() {
	/*错误的方式
	for i := 0; i < 10; i++ {

		go func() {
			fmt.Println(i)
		}()
	}
	*/
	for i := 0; i < 10; i++ {

		go func(j int) {
			fmt.Println(j)
		}(i)
	}
}
