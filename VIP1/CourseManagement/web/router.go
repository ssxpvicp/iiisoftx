package web

import "net/http"

func init() {
	http.HandleFunc("/add/teacher", AddTeacher)
}
