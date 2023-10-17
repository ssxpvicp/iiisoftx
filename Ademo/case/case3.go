package _case

import (
	"fmt"
	"time"
)

type Pro struct {
	name string `json:"name"`
	code int    `json:"code"`
}

func Case3() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sli = append(sli[0:3], sli[4:]...)
	fmt.Println(sli)

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli())
}
