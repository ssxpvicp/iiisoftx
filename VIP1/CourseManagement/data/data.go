package data

import (
	"encoding/json"
	"fmt"
	"sync"
)

// 别名
type TableName string

const (
	Stu     TableName = "stu"
	Course  TableName = "course"
	Teacher TableName = "teacher"
	Class   TableName = "class"
)

type item struct {
	data   map[int]interface{}
	count  int
	maxId  int
	locker sync.RWMutex
}

// 这样就不是一个简单的mp了
var dataStorage map[TableName]*item

func init() {
	dataStorage = map[TableName]*item{
		Stu: &item{
			data:  make(map[int]interface{}),
			count: 0,
			maxId: 0,
		},
		Teacher: &item{
			data:  make(map[int]interface{}),
			count: 0,
			maxId: 0,
		},
		Class: &item{
			data:  make(map[int]interface{}),
			count: 0,
			maxId: 0,
		},
		Course: &item{
			data:  make(map[int]interface{}),
			count: 0,
			maxId: 0,
		},
	}
}
func Add(tableName TableName, obj interface{}) (id int, err error) {
	//锁一下,写锁
	dataStorage[tableName].locker.Lock()
	defer dataStorage[tableName].locker.Unlock()
	//dataStorage[tableName].locker.RLock()//读锁
	//defer dataStorage[tableName].locker.RUnlock()//读锁
	//
	id = dataStorage[tableName].maxId + 1
	dataStorage[tableName].count += 1
	dataStorage[tableName].maxId = id
	dataStorage[tableName].data[id] = obj
	return
}
func Edit(tableName TableName, id int, obj interface{}) (err error) {
	//锁一下,写锁
	dataStorage[tableName].locker.Lock()
	defer dataStorage[tableName].locker.Unlock()
	//dataStorage[tableName].locker.RLock()//读锁
	//defer dataStorage[tableName].locker.RUnlock()//读锁
	//

	dataStorage[tableName].data[id] = obj
	return
}
func Del(tableName TableName, id int) (err error) {
	//锁一下,写锁
	dataStorage[tableName].locker.Lock()
	defer dataStorage[tableName].locker.Unlock()
	//dataStorage[tableName].locker.RLock()//读锁
	//defer dataStorage[tableName].locker.RUnlock()//读锁
	//

	dataStorage[tableName].count -= 1
	delete(dataStorage[tableName].data, id)

	return
}
func Get(tableName TableName, id ...int) (mp map[int]interface{}, err error) {
	dataStorage[tableName].locker.RLock()
	defer dataStorage[tableName].locker.RUnlock()

	if len(id) > 0 { //大于0,说明他是切片了
		mp = make(map[int]interface{}, len(id))
		for _, i := range id {
			item, ok := dataStorage[tableName].data[i]
			if ok {
				mp[i] = item
			}
		}
	} else {
		mp = make(map[int]interface{})
		count := 100 //默认返回100条
		j := 0
		for key, item := range dataStorage[tableName].data {
			if j >= count {
				break
			}
			mp[key] = item
			j++
		}
	}
	return
}

func ShowAllData(tableName TableName) {
	item := dataStorage[tableName]
	str, _ := json.Marshal(item.data)
	fmt.Println("-------------start--------------")
	fmt.Println("tableName:", tableName)
	fmt.Println("count:", item.count)
	fmt.Println("maxId", item.maxId)
	fmt.Println("data", string(str))
	fmt.Println("-------------end--------------")
}
