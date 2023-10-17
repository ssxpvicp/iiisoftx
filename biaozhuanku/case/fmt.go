package _case

import (
	"fmt"
	"os"
)

func FmtCase() {
	//打印到标准输出
	fmt.Println("今天天气不错")
	//格式化,并打印到标准输出
	fmt.Printf("%s天气很好\n", "今天")
	//格式化
	str := fmt.Sprintf("%s天气很好", "今天和明天都")
	//输出到io.writer
	_, err := fmt.Fprintf(os.Stderr, str)
	if err != nil {
		return
	}
	fmt.Println()
}
func FmtCase1() {
	type simple struct {
		value int
	}
	//通用占位符
	a := simple{value: 10}
	fmt.Printf("默认格式的值: %v \n", a)
	fmt.Printf("包含字段名的值:%+v \n", a)
	fmt.Printf("go语法表示的值:%#v \n", a)
	fmt.Printf("go语法表示的类型%T \n", a)
	fmt.Printf("输入百分号%%10 \n")

	//整数占位符
	v1 := 10
	v2 := 20170 //"今"字的码点值
	fmt.Printf("二进制: %b \n", v1)
	fmt.Printf("Unicode 码点转字符: %c \n", v2)
	fmt.Printf("十进制:%d \n", v1)
	fmt.Printf("八进制:%o \n", v1)
	fmt.Printf("Oo为前缀的八进制:%O \n", v1)
	fmt.Printf("用单引号将字符的值包起来: %q \n", v2)
	fmt.Printf("十六进制:%x \n", v1)
	fmt.Printf("十六进制大写:%X \n", v1)
	fmt.Printf("Unicode格式: %U \n", v2)
	//宽度设置
	fmt.Printf("%v 的二进制:%b;go语法表示为:%#b;指定二进制宽度为8且不足8位自动补0:%08b \n", v1, v1, v1, v1)
	fmt.Printf("%v 的十六进制:%b;go语法表示为:%#x;指定十六进制宽度为8且不足8位自动补0:%08x \n", v1, v1, v1, v1)
	fmt.Printf("%v 的字符为:%c;go语法表示为:%#c;指定宽度为5且不足5位自动补空格:%5c \n", v2, v2, v2, v2)
	//浮点占位符
	var f1 = 123.789
	var f2 = 12345678910.78999
	fmt.Printf("指数为二的幂的无小数科学计数法: %b \n", f1)
	fmt.Printf("数科学计数法: %e \n", f1)
	fmt.Printf("数科学计数法,大写: %E \n", f1)
	fmt.Printf("有小数点而无指数,即常规的浮点数格式.默认宽度和精度: %f \n", f1)
	fmt.Printf("宽度为9精度默认: %9f \n", f1)
	fmt.Printf("默认宽度,精度保留两位小数: %.2f \n", f1)
	fmt.Printf("宽度为9,精度保留两位小数: %9.2f \n", f1)
	fmt.Printf("宽度为9,精度保留0位小数: %9.f \n", f1)
	fmt.Printf("根据情况自动选择%%E或%%E来输出,以生产更紧凑的输出(未位无0): %G %G\n", f1, f2)
	fmt.Printf("以十六进制方式表示: %x\n", f1)
	fmt.Printf("以十六进制大写方式表示: %X\n", f1)

	//字符串占位符
	var str = "今天是个好日子"
	fmt.Printf("描述一下今天:%s \n", str)
	fmt.Printf("描述一下今天:%q \n", str)
	fmt.Printf("十六进制显示:%x \n", str)
	fmt.Printf("十六进制大写显示:%X \n", str)

	//指针占位符
	var str1 = "今天是个好日子"
	bytes := []byte(str1)
	fmt.Printf("切片的第0个元素的地址 %p \n", bytes)

	mp := make(map[string]int, 0)
	fmt.Printf("集合的第0个元素的地址 %p \n", mp)

	var p *map[string]int = new(map[string]int)
	fmt.Printf("指针的第0个元素的地址 %p \n", p)
}
