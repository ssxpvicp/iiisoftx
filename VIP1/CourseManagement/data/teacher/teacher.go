package teacher

import (
	"basis/VIP1/CourseManagement/common"
	"basis/VIP1/CourseManagement/data"
	"encoding/json"
	"time"
)

type Teacher struct {
	ID           int
	Name         string
	StuIDList    map[int]int
	HeadmasterID int
	Gender       common.Gender
	Birthday     time.Time
}

func (t *Teacher) String() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}

type TeacherData struct {
	tableName data.TableName
}

func NewTeacherData() *TeacherData {
	return &TeacherData{
		tableName: data.Teacher,
	}
}
func (d *TeacherData) Add(Teacher *Teacher) (int, error) {
	id, err := data.Add(d.tableName, Teacher)
	Teacher.ID = id
	data.ShowAllData(d.tableName)
	return Teacher.ID, err
}
func (d *TeacherData) Edit(Teacher *Teacher) error {
	err := data.Edit(d.tableName, Teacher.ID, Teacher)
	data.ShowAllData(d.tableName)
	return err
}
func (d *TeacherData) del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}
func (d *TeacherData) get(id ...int) ([]*Teacher, error) {
	//使用了...int,下面的id参数是int 就要使用id...,打散语法
	list := make([]*Teacher, 0)
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
			Teacher, ok := v.(*Teacher)
			if !ok {
				continue
			}
			Teacher.ID = i
			list = append(list, Teacher) //append(list,stu...)append也支持打散语法
		}
	} else {
		for k, v := range mp {
			Teacher, ok := v.(*Teacher)
			if !ok {
				continue
			}
			Teacher.ID = k
			list = append(list, Teacher)
		}
	}
	return list, err
}
