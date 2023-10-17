package main

import (
	"net/http"
)

type Stu struct {
	Name string
	Age  int
}

func main() {
	/*
		d := course.NewCourseData()
		obj := &course.Course{
			Name: "nick",
		}
		d.Add(obj)
		data.ShowAllData(data.Stu)
		/*
	*/
	/*
		d := stu.NewStuData()
		obj := &stu.Stu{
			Name:   "nick",
			Gender: 0,
		}
		d.Add(obj)
		data.ShowAllData(data.Stu)
	*/
	/*
		s0 := &s{
			Name: "nick",
			Age:  80,
		}
		s1 := &s{
			Name: "nick1",
			Age:  20,
		}
		s2 := &s{
			Name: "nick2",
			Age:  10,
		}
		data.Add(data.Course, s0)
		data.Add(data.Course, s1)
		data.Add(data.Course, s2)

		data.ShowAllData(data.Course)

		s2.Name = "改一下"
		s2.Age = 99
		data.Edit(data.Course, 3, s2)
		data.ShowAllData(data.Course)

		data.Del(data.Course, 2)
		data.ShowAllData(data.Course)

		get, _ := data.Get(data.Course, 1, 3)
		for i, i2 := range get {
			fmt.Println(i, i2)
		}
	*/
	/*时间转换练习
	str := "2023-09-03 21:12:50"
	t := time.Now()
	t1, err := common.StrToTime(str, common.DataTime)
	fmt.Println(t1, err)
	fmt.Println(common.TimeToStr(t, common.DataTimeZone))
	*/
	//go自带的httpServer
	http.ListenAndServe(":8080", nil)
}
