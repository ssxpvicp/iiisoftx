package web

import (
	"basis/VIP1/CourseManagement/common"
	"basis/VIP1/CourseManagement/data/teacher"
	"fmt"
	"net/http"
	"strconv"
)

func AddTeacher(w http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("name")
	GenderStr := req.FormValue("gender")
	BirthdayStr := req.FormValue("birthday")
	Gender, _ := strconv.Atoi(GenderStr)
	t := &teacher.Teacher{
		Name:   Name,
		Gender: common.Gender(Gender),
	}
	birthday, err := common.StrToTime(BirthdayStr)
	if err == nil {
		t.Birthday = birthday
	}
	_, err = tercherdata.Add(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, t.String())
}
