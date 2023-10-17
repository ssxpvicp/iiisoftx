package class

import "basis/VIP1/CourseManagement/data"

type Class struct {
	ID           int
	Name         string
	StuIDList    map[int]int
	HeadmasterID int
}
type ClassData struct {
	tableName data.TableName
}

func NewClassData() *ClassData {
	return &ClassData{
		tableName: data.Class,
	}
}
func (d *ClassData) Add(Class *Class) (int, error) {
	id, err := data.Add(d.tableName, Class)
	Class.ID = id
	data.ShowAllData(d.tableName)
	return Class.ID, err
}
func (d *ClassData) Edit(Class *Class) error {
	err := data.Edit(d.tableName, Class.ID, Class)
	data.ShowAllData(d.tableName)
	return err
}
func (d *ClassData) del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}
func (d *ClassData) get(id ...int) ([]*Class, error) {
	//使用了...int,下面的id参数是int 就要使用id...,打散语法
	list := make([]*Class, 0)
	mp, err := data.Get(d.tableName, id...)
	if err != nil {
		return nil, err
	}
	if len(id) > 0 {
		for _, i := range id {
			v, ok := mp[i]
			if !ok {
				continue
			}
			Class, ok := v.(*Class)
			if !ok {
				continue
			}
			Class.ID = i
			list = append(list, Class) //append(list,stu...)append也支持打散语法
		}
	} else {
		for k, v := range mp {
			Class, ok := v.(*Class)
			if !ok {
				continue
			}
			Class.ID = k
			list = append(list, Class)
		}
	}
	return list, err
}
