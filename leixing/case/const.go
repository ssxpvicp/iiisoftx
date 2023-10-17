package _case

import "fmt"

const (
	//使用了iota 中间如果有一行不用.也要用_来代替
	B = 1 << (10 * iota)
	KB
	MB
	_
	TB
)

type Gender uint

const (
	FEMALE Gender = iota
	MALE
	THIRD
)

func ConstAndEnumCase() {

	const (
		A int64 = 10
		B       = 20
	)
	size := MB
	var gender Gender = MALE

	fmt.Println(A, B, gender, size)
}
