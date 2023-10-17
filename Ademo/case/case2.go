package _case

import (
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"time"
)

func RegCase() {
	// 创建一个正则表达式匹配规则对象
	// reg := regexp.MustCompile(规则字符串)
	// 利用正则表达式匹配规则对象匹配指定字符串
	// 会将所有匹配到的数据放到一个字符串切片中返回
	// 如果没有匹配到数据会返回nil
	// res := reg.FindAllString(需要匹配的字符串, 匹配多少个)

	str := "Hello 李南江 1232"
	reg := regexp.MustCompile("2")
	res := reg.FindAllString(str, -1) //-1就是不限制数量
	fmt.Println(res)                  // [2 2]
	for i, re := range res {
		fmt.Println(i, re)
	}
	res = reg.FindAllString(str, 1)
	fmt.Println(res) // [2]

	//*时间和日期函数
	var t = time.Now()
	fmt.Println(t) //2023-10-04 14:54:31.6132947 +0800 CST m=+0.008828601
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Printf("当前的时间是: %d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// 生成 10 位时间戳（秒级别）
	timestamp10 := time.Now().Unix()
	fmt.Println(timestamp10)
	// 生成 13 位时间戳（毫秒级别）
	timestamp13 := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Println(timestamp13)
	// 2006/01/02 15:04:05这个字符串是Go语言规定的, 各个数字都是固定的
	// 字符串中的各个数字可以只有组合, 这样就能按照需求返回格式化好的时间
	str1 := t.Format("2006/01/02 15:04:05")
	fmt.Println(str1)
	str2 := t.Format("2006/01/02")
	fmt.Println(str2)
	str3 := t.Format("15:04:05")
	fmt.Println(str3)

	// 获取1970年1月1日距离当前的时间(秒)
	fmt.Println(t.Unix())
	// 获取1970年1月1日距离当前的时间(纳秒)
	fmt.Println(t.UnixNano())

	// 创建随机数种子
	fmt.Println("rand.Intn", rand.Intn(100))
	rand.Seed(time.Now().UnixNano()) // 生成一个随机数
	fmt.Println(rand.Intn(10))

	//*Url编码
	urls := url.QueryEscape("郑健加油")
	fmt.Println(urls)
	fmt.Println(url.QueryUnescape(urls))

}
