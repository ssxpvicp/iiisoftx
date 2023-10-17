package _case

import "fmt"

func VarDeclareCase() {
	var i int
	var j int = 100
	var f float32 = 100.32
	//通过:=以推断的方式定义变量并赋值只能在局部变量的定义
	b := true
	//数组 Golang中一般是不用数组的.这里的数组是定长的.
	var arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{2, 3, 4, 5, 6}
	var arr2 [5]int
	arr2[2] = 4
	arr2[3] = 5
	fmt.Println(i, j, f, b, arr, arr1, arr2)

	//指针类型,用于表示变量地址的类型
	var intPtr *int
	var floatPtr *float64
	var i1 = 100
	//f1的参数为指针类型,那传参数的时候.要加上&  ,变成  &i1
	f1(&i1)

	//接口类型
	var inter interface{}
	inter = 123 //可以随便赋值
	fmt.Println(intPtr, floatPtr, i1, inter)

}
func f1(i *int) {
	//指针类型前面加上 *  那就是取指针类型i的值
	*i = *i + 1
}
