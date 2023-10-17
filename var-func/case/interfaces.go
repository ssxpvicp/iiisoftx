package _case

import "fmt"

// 基本接口,可用于变量的定义
type ToStrings interface {
	Strings() string
}

func (u user) String() string {
	return fmt.Sprintf("ID:%d,Name:%s,Age:%d", u.Id, u.Name, u.Age)
}

func (addr address) String() string {
	return fmt.Sprintf("ID:%d,Province:%s,City:%s", addr.ID, addr.Province, addr.City)
}

// 泛型接口
type GetKey[T comparable] interface {
	any
	GET() T
}

func (u user) Get() int64 {
	return u.Id
}
func (addr address) Get() int {
	return addr.ID
}

// 列表转集合
func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, data := range list {
		mp[data.GET()] = data
	}
	return mp
}

//func InterfaceCase() {
//	userList := []GetKey[int]{
//		user{Id: 1, Name: "nick", Age: 18},
//		user{Id: 2, Name: "king", Age: 18},
//	}
//	addList := []GetKey[int]{
//		address{
//			ID:       1,
//			Province: "江苏",
//			City:     "苏州",
//		},
//		address{
//			ID:       2,
//			Province: "江苏",
//			City:     "南京",
//		},
//	}
//	userMp:=listToMap(userList)
//
//
//}
