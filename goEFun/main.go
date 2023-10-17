package main

import (
	. "github.com/duolabmeng6/goefun/ecore"
	. "github.com/duolabmeng6/goefun/ehttp"
	. "github.com/duolabmeng6/goefun/etool"
)

func main() {
	E调试输出("欢迎使用 go-efun")
	http := NewHttp()
	返回文本, 访问失败 := http.Get("https://www.taobao.com/help/getip.php")
	if 访问失败 != nil {
		E调试输出("访问失败", 返回文本)
	}
	E调试输出(返回文本)

	data := New存取键值表()
	data.Set("aaa", "111")
	data.Set("bbb", "222")
	data.Set("ccc", "333")
	data.SetArray("list", "a")
	data.SetArray("list", "b")
	data.SetArray("list", "c")

	E调试输出(data.ToJson(true))

	E调试输出(E取md5从文本("1234567"))

	E调试输出(E取现行时间().E时间到文本("Y-m-d H:i:s"))
	E调试输出(E取md5从文本("123"))

}
