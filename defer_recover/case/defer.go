package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

// defer 关键词用来声明一个延迟调用的函数
// 该函数可以有匿名函数也可以是具名函数
// defer 延迟函数的顺序为后进先出,同栈原理
func DeferCase1() {
	fmt.Println("开始执行DeferCase1")
	defer func() {
		fmt.Println("调用了匿名函数1")
	}()
	defer f1()
	defer func() {
		fmt.Println("调用了匿名函数2")
	}()
	fmt.Println("结束执行DeferCase1")

}

// 参数预计算
func DeferCase2() {
	i := 1
	//传参 加入deferList的的参数值
	defer func(j int) {
		fmt.Println("defer j:", j) //print 2
	}(i + 1)
	//闭包
	defer func() { fmt.Println("defer", i) }() //print 99
	i = 99
	fmt.Println(i) //print 99
}

// 返回值
// defer 函数执行在return之后
func DeferCase3() {
	var i, j = f2()
	fmt.Println("f2返回:", i, *j, g)
}

// 人为panic制造异常
func ExceptionCase() {
	defer func() {
		//捕获异常,恢复协程
		err := recover() //这里的err就是panic(内容)
		//异常处理
		if err != nil {
			fmt.Println("异常处理 defer recover", err)
		}
	}()
	fmt.Println("开始执行ExceptionCase")
	panic("ExceptionCase抛出异常")
	fmt.Println("结束执行ExceptionCase")

}

// 资源释放
func FileReadCase() {
	file, err := os.Open("C:\\Users\\Administrator\\Desktop\\韩文路径.e")

	if err != nil {
		log.Fatal(err)
		return
	}
	//通过defer 调用资源释放的方法
	defer func() {
		fmt.Println("释放文件资源")
		file.Close()
	}()

	fmt.Println(file.Name())
	buf := make([]byte, 1024)
	//读取数据
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

	}
	fmt.Println(buf)

}

var g = 100

// f2这个例子可以证明,defer在 return 后,在返回函数返回结果之前
func f2() (int, *int) {
	defer func() {
		g = 200
	}()
	fmt.Println("f2 g:", g)
	return g, &g
}

func f1() {
	fmt.Println("调用了具名函数f1")
}
