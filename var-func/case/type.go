package _case

import "fmt"

type user struct {
	Id   int64
	Name string
	Age  uint8
}
type address struct {
	ID       int
	Province string
	City     string
}

// 集合转列表
func mapToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}
func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}

}
func TTypeCase() {
	userMp := make(map[int64]user, 0)
	userMp[1] = user{
		Id:   1,
		Name: "nick",
		Age:  18,
	}
	userMp[2] = user{
		Id:   3,
		Name: "king",
		Age:  19,
	}
	userList := mapToList[int64, user](userMp)

	ch := make(chan user)
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addMp := make(map[int64]address, 0)
	addMp[1] = address{
		ID:       1,
		Province: "江苏",
		City:     "常熟",
	}
	addMp[2] = address{
		ID:       3,
		Province: "江苏",
		City:     "苏州",
	}
	addList := mapToList[int64, address](addMp)

	ch1 := make(chan address)
	go myPrintln(ch1)
	for _, u1 := range addList {
		ch1 <- u1
	}

}

// 泛型切片的定义
type List[T any] []T

// 泛型集合的定义 //声明了两个泛型  k 和 v
type MapT[K comparable, v any] map[K]v

// 泛型通道的定义
type Chan[T any] chan T

func TTypeCase1() {
	userMp := make(MapT[int64, user], 0)
	userMp[1] = user{
		Id:   1,
		Name: "nick",
		Age:  18,
	}
	userMp[2] = user{
		Id:   3,
		Name: "king",
		Age:  19,
	}
	var userList List[user]
	userList = mapToList[int64, user](userMp)

	ch := make(Chan[user])
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addMp := make(MapT[int64, address], 0)
	addMp[1] = address{
		ID:       1,
		Province: "江苏",
		City:     "常熟",
	}
	addMp[2] = address{
		ID:       3,
		Province: "江苏",
		City:     "苏州",
	}
	var addList List[address]
	addList = mapToList[int64, address](addMp)

	ch1 := make(Chan[address])
	go myPrintln(ch1)
	for _, u1 := range addList {
		ch1 <- u1
	}

}
