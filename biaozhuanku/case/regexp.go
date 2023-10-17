package _case

import (
	"fmt"
	"regexp"
)

func RegexpCase() {
	//构建一个正则表达式的对象
	ret := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Println(ret.MatchString("abcd[1234]"))
	//从给定的字符串查找符合条件的字符串
	byteS := ret.FindAll([]byte("efb[1233]"), -1)
	fmt.Println(string(byteS[0]))
}
