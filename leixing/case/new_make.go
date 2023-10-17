package _case

import "fmt"

func NewCase() {
	//通过NEW函数,可以创建任意类型,并返回指针
	mpPtr := new(map[string]*user)
	fmt.Println(mpPtr)

	slicePtr := new([]user)
	if *slicePtr == nil {
		fmt.Println("切片值为nil", *slicePtr)
	}
	*slicePtr = append(*slicePtr, user{Name: "nick", Age: 19})
	fmt.Println(slicePtr)

	userPtr := new(user)
	strPtr := new(string)
	fmt.Println(userPtr, strPtr)
}

// make 仅用于 切片 集合 通道的初始化
func MakeCase() {
	//初始化切片,并设置长度和容量
	slice := make([]int, 10, 20)
	slice[0] = 100
	//初始化集合,并设置集合的初始大小

	mp := make(map[string]string, 10)
	mp["A"] = "a"
	fmt.Println(mp)

	//初始化通道,设置通道的读写方向和缓冲大小
	ch := make(chan int, 10) //可读可写
	fmt.Println(ch)
	ch1 := make(chan<- int, 10) //只写
	fmt.Println(ch1)
	ch2 := make(<-chan int, 10) //只读
	fmt.Println(ch2)
}

func SliceAndMapCase() {
	//定义切片
	var slice []int
	slice = []int{1, 2, 3, 4, 5, 6}
	slice1 := make([]int, 10)
	slice1[0] = 1011
	slice1[1] = 10
	fmt.Println(slice, slice1)
	//切片的截取
	slice2 := make([]int, 5, 10)
	slice2[0] = 0
	slice2[1] = 1
	slice2[2] = 2
	slice2[3] = 3
	slice2[4] = 4

	fmt.Println(len(slice2), cap(slice2))
	slice3 := slice2[2:5] //从0个到第4个
	fmt.Println(len(slice3), cap(slice3), slice3)

	slice3 = append(slice3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(slice3), cap(slice3), slice3)

	//集合,无序
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	mp["B"] = "b"
	mp["C"] = "c"
	mp["C"] = "d"
	fmt.Println(mp)

	for k, v := range mp {
		//无序
		fmt.Println(k, v)
	}
}
