package stu

import (
	"basis/VIP1/CourseManagement/common"
	"basis/VIP1/CourseManagement/data"
	"time"
)

type Stu struct {
	ID           int
	Name         string
	Gender       common.Gender
	Birthday     time.Time
	CourseIDList map[int]int
	ClassIDList  map[int]int
}

type StuData struct {
	tableName data.TableName
}

func NewStuData() *StuData {
	return &StuData{
		tableName: data.Stu,
	}
}
func (d *StuData) Add(stu *Stu) (int, error) {
	id, err := data.Add(d.tableName, stu)
	stu.ID = id
	data.ShowAllData(d.tableName)
	return stu.ID, err
}
func (d *StuData) Edit(stu *Stu) error {
	err := data.Edit(d.tableName, stu.ID, stu)
	data.ShowAllData(d.tableName)
	return err
}
func (d *StuData) del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}
func (d *StuData) get(id ...int) ([]*Stu, error) {
	//使用了...int,下面的id参数是int 就要使用id...,打散语法
	list := make([]*Stu, 0)
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
			stu, ok := v.(*Stu)
			if !ok {
				continue
			}
			stu.ID = i
			list = append(list, stu) //append(list,stu...)append也支持打散语法
		}
	} else {
		for k, v := range mp {
			stu, ok := v.(*Stu)
			if !ok {
				continue
			}
			stu.ID = k
			list = append(list, stu)
		}
	}
	return list, err
}
