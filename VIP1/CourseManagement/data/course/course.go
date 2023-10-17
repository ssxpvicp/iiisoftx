package course

import "basis/VIP1/CourseManagement/data"

type Course struct {
	ID            int
	Name          string
	TeacherIDList map[int]int
	StuIDList     map[int]int
	classIDList   map[int]int
	LastClassID   int
}
type CourseData struct {
	tableName data.TableName
}

func NewCourseData() *CourseData {
	return &CourseData{
		tableName: data.Course,
	}
}
func (d *CourseData) Add(Course *Course) (int, error) {
	id, err := data.Add(d.tableName, Course)
	Course.ID = id
	data.ShowAllData(d.tableName)
	return Course.ID, err
}
func (d *CourseData) Edit(Course *Course) error {
	err := data.Edit(d.tableName, Course.ID, Course)
	data.ShowAllData(d.tableName)
	return err
}
func (d *CourseData) del(id int) error {
	err := data.Del(d.tableName, id)
	data.ShowAllData(d.tableName)
	return err
}
func (d *CourseData) get(id ...int) ([]*Course, error) {
	//使用了...int,下面的id参数是int 就要使用id...,打散语法
	list := make([]*Course, 0)
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
			Course, ok := v.(*Course)
			if !ok {
				continue
			}
			Course.ID = i
			list = append(list, Course) //append(list,stu...)append也支持打散语法
		}
	} else {
		for k, v := range mp {
			Course, ok := v.(*Course)
			if !ok {
				continue
			}
			Course.ID = k
			list = append(list, Course)
		}
	}
	return list, err
}
