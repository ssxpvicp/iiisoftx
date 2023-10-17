package _case

import (
	"fmt"
	"sort"
)

type sortUser struct {
	ID   int64
	Name string
	Age  uint8
}
type ByID []sortUser

// 获取长度
func (a ByID) Len() int {
	return len(a)
}

// 交换位置
func (a ByID) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 返回i位置的ID是否小于j位置的ID
func (a ByID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

func SortCase() {
	list := []sortUser{
		{ID: 1, Name: "nick", Age: 19},
		{ID: 2, Name: "nick", Age: 18},
		{ID: 11, Name: "nick", Age: 17},
		{ID: 4, Name: "nick", Age: 16},
		{ID: 5, Name: "nick", Age: 15},
		{ID: 33, Name: "nick", Age: 14},
		{ID: 8, Name: "nick", Age: 13},
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	fmt.Println(list)

	sort.Sort(ByID(list))
	fmt.Println(list)
}
