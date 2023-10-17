package _case

import (
	"errors"
	"fmt"
)

//形参
//局部变量
//全局变量

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{name, age}
}

// 获取user name属性
func (u *User) GetName() string {
	return u.Name
}

// 获取user age属性
func (u *User) GetAge() int {
	return u.Age
}

// 函数可以有多个形参,和多个返回值,返回值也可以和形参一样拥有参数名称
func SumCase(a, b int) (sum int, err error) {

	if a <= 0 && b <= 0 {
		err = errors.New("两数相加不能同时小于0")
		return
	}
	sum = a + b
	return
}

// 值传递与引用传递
func ReferenceCase(a int, b *int) {
	a += 1
	*b += 1

}

// 变量作用域
var g int
var G int

func ScopeCase(a, b int) {
	c := 100
	g = a + b + c
	G = g
}

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println("不使用泛型", getMaxNumInt(a, b))
	fmt.Println("不使用泛型", getMaxNumFloat(c, d))
	fmt.Println("使用泛型", getMaxNum(a, b)) //可以不指定类型.也可以指定类型 getMaxNum[int](a, b)
	fmt.Println("使用泛型", getMaxNum[float64](c, d))

}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 泛型方法 [T int | float64]定义T 为int或float64类型
// 如果类型传的是指针类型如  *int  那么就要用interface{*int|*float64},不然会识别为乘法T*int了.
func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 自定义泛型
type CusNumT interface {
	//~表示支持类型的衍生类型
	//| 表示取并集
	// 多行之间取交集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

// MyInt64 为int64的衍生类型,是具有基础类型的int64的新类型,与int64是不同类型
type MyInt64 int64

// MyInt32 为int32的别名,与int32 是同一类型
type MyInt32 = int32

func CusNumTCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println("自定义泛型,数字比较", getMaxCusNum(a, b))
	fmt.Println("自定义泛型,数字比较", getMaxCusNum(a1, b1))

	var c, d float64 = 5, 6
	fmt.Println("自定义泛型,数字比较", getMaxCusNum(c, d))

	var e, f float64 = 7, 8
	var g, h MyInt64 = 7, 8
	fmt.Println("自定义泛型,数字比较", getMaxCusNum(e, f))
	fmt.Println("自定义泛型,数字比较", getMaxCusNum(g, h))

}
func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 内置类型
func BuiltInCase() {
	var a, b string = "abc", "def"
	var c, d int = 100, 100
	fmt.Println("comparable来判断类型", getBuiltInComparable(a, b))
	fmt.Println("comparable来判断类型", getBuiltInComparable(c, d))
}
func getBuiltInComparable[T comparable](a, b T) bool {
	//comparable 类型 ,只支持 == !=两个操作
	if a == b {
		return true
	}
	return false
}
