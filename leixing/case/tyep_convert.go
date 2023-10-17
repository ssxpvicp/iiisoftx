package _case

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

// 类型转换
func ConvertCase() {
	//同一类数据转换 数字和数字  字符串和字符和字节

	//不同类型的数据转换 数字和字符

	//接口类型转其它类型

	//数字类型的转换  高精度转低进度的话.需要自己注意
	var num1 int = 1099
	fmt.Println(int64(num1))
	var num2 int64 = 100
	fmt.Println(int(num2))

	//字符串与数字类型转换
	var num3 int = 100
	fmt.Println(strconv.Itoa(num3) + "abc")
	var str1 = "100"
	fmt.Println(strconv.Atoi(str1))
	var num4 int64 = 10010
	fmt.Println(strconv.FormatInt(num4, 10))
	var str2 = "1010"
	fmt.Println(strconv.ParseInt(str2, 10, 64))

	//字符串与[]byte 转换
	var str3 = "今天心情不是很好"
	var bytes1 = []byte(str3)
	fmt.Println(bytes1)
	fmt.Println(string(bytes1))

	//字符串与rune转换 rune是go专门处理unicode字符类型的
	// rune 其实是 int32类型
	//将字符串转为rune切片,实际上rune切片中存储了字符串的Unicode码点
	var rune1 = []rune(str3) //转成unicode的码点
	fmt.Println(rune1)
	fmt.Println(string(rune1))
	fmt.Println(string(rune1[3])) //第4个字

	//接口类型,转具体类型
	//就是断言,通过接口断言来转
	var inf interface{} = 100
	var infStruct interface{} = user{}
	fmt.Println(inf, infStruct)
	i, err := inf.(int)
	fmt.Println(i, err)
	u, ok := infStruct.(user)
	fmt.Println(u, ok)

	//时间类型转字符串
	var t time.Time
	t = time.Now()
	timeStr := t.Format("2006-01-02 15:04:05:Z07:00") //go写死的6-1-2-15-4-5
	fmt.Println(timeStr)

	//字符串转时间
	t2, _ := time.Parse("2006-01-02 15:04:05:Z07:00", timeStr)
	fmt.Println(t2)

	//uintptr
	u1 := user{}
	fmt.Println(&u1) //&{ 0 { }}
	//unsafe.Pointer 是一个通用的指针类型,不能用于计算
	uPtr := unsafe.Pointer(&u1)
	fmt.Println(uPtr) //0xc0000a0000
	namePtr := unsafe.Pointer(uintptr(uPtr) + unsafe.Offsetof(u1.Name))
	fmt.Println(namePtr)
	*(*string)(namePtr) = "郑健 加油学习GoLang!"
	fmt.Println(u1)

}
