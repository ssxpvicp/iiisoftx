package _case

import (
	"fmt"
	"strings"
)

func StringsCase() {
	//*查找子串在字符串中出现的位置
	//查找'字符'在字符串中的第一次出现的位置,找不到返回-1
	res := strings.IndexByte("hello 郑健", 'l')
	fmt.Println("IndexByte", res)
	// 查找`汉字`OR`字符`在字符串中第一次出现的位置, 找不到返回-1
	res = strings.IndexRune("hello 你好郑健", '郑')
	fmt.Println("IndexRune", res)
	res = strings.IndexRune("hello 你好郑健", 'o')
	fmt.Println("IndexRune", res)

	// 查找`汉字`OR`字符`中任意一个在字符串中第一次出现的位置, 找不到返回-1
	res = strings.IndexAny("hello 你好郑健", "你好")
	fmt.Println("IndexAny", res)
	// 会把wmhl拆开逐个查找, w、m、h、l只要任意一个被找到, 立刻停止查找
	res = strings.IndexAny("hello 你好郑健", "wmhl")
	fmt.Println("IndexAny", res)

	// 查找`子串`在字符串第一次出现的位置, 找不到返回-1
	res = strings.Index("hello 李南江", "llo")
	fmt.Println(res) // 2
	// 会把lle当做一个整体去查找, 而不是拆开
	res = strings.Index("hello 李南江", "lle")
	fmt.Println(res) // -1
	// 可以查找字符也可以查找汉字
	res = strings.Index("hello 李南江", "南")
	fmt.Println(res) // 6

	// 会将字符串先转换为[]rune, 然后遍历rune切片逐个取出传给自定义函数
	// 只要函数返回true,代表符合我们的需求, 既立即停止查找
	res = strings.IndexFunc("hello 李南江", custom)
	fmt.Println(res) // 6

	// 倒序查找`子串`在字符串第一次出现的位置, 找不到返回-1
	res = strings.LastIndex("hello 李南江", "l")
	fmt.Println(res) // 3

	//*判断字符串中是否包含子串
	// 查找`子串`在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.Index函数
	resBool := strings.Contains("hello 代码情缘", "llo")
	fmt.Println(resBool) // true

	// 查找`汉字`OR`字符`在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.IndexRune函数
	resBool = strings.ContainsRune("hello 代码情缘", 'l')
	fmt.Println(resBool) // true
	resBool = strings.ContainsRune("hello 代码情缘", '李')
	fmt.Println(resBool) // true

	// 查找`汉字`OR`字符`中任意一个在字符串中是否存在, 存在返回true, 不存在返回false
	// 底层实现就是调用strings.IndexAny函数
	resBool = strings.ContainsAny("hello 代码情缘", "wmhl")
	fmt.Println(resBool) // true

	// 判断字符串是否已某个字符串开头
	resBool = strings.HasPrefix("lnj-book.avi", "lnj")
	fmt.Println(resBool) // true

	// 判断字符串是否已某个字符串结尾
	resBool = strings.HasSuffix("lnj-book.avi", ".avi")
	fmt.Println(resBool) // true

	//*字符串比较
	// 比较两个字符串大小, 会逐个字符地进行比较ASCII值
	// 第一个参数 >  第二个参数 返回 1
	// 第一个参数 <  第二个参数 返回 -1
	// 第一个参数 == 第二个参数 返回 0
	res1 := strings.Compare("bcd", "abc")
	fmt.Println(res1) // 1
	res1 = strings.Compare("bcd", "bdc")
	fmt.Println(res1) // -1
	res1 = strings.Compare("bcd", "bcd")
	fmt.Println(res1) // 0

	// 判断两个字符串是否相等, 可以判断字符和中文
	// 判断时会忽略大小写进行判断
	res2 := strings.EqualFold("abc", "def")
	fmt.Println(res2) // false
	res2 = strings.EqualFold("abc", "abc")
	fmt.Println(res2) // true
	res2 = strings.EqualFold("abc", "ABC")
	fmt.Println(res2) // true
	res2 = strings.EqualFold("代码情缘", "代码情缘")
	fmt.Println(res2) // true

	//*字符串转换
	res3 := strings.ToLower("ABC")
	fmt.Println(res3)

	// 将字符串转换为小写
	res3 = strings.ToLower("ABC")
	fmt.Println(res3) // abc

	// 将字符串转换为大写
	res3 = strings.ToUpper("abc")
	fmt.Println(res3) // ABC

	// 将字符串转换为标题格式, 大部分`字符`标题格式就是大写
	res3 = strings.ToTitle("hello world")
	fmt.Println(res3) // HELLO WORLD
	res3 = strings.ToTitle("HELLO WORLD")
	fmt.Println(res3) // HELLO WORLD

	//*字符串拆合
	// 按照指定字符串切割原字符串
	// 用,切割字符串
	arr1 := strings.Split("a,b,c", ",")
	fmt.Println(arr1)
	arr2 := strings.Split("ambmc", "m")
	fmt.Println(arr2) // [a b c]

	// 按照指定字符串切割原字符串, 并且指定切割为几份
	// 如果最后一个参数为0, 那么会范围一个空数组
	arr3 := strings.SplitN("a,b,c", ",", 2)
	fmt.Println(arr3)
	arr4 := strings.SplitN("a,b,c", ",", 0)
	fmt.Println(arr4) // []

	// 按照指定字符串切割原字符串, 切割时包含指定字符串
	arr5 := strings.SplitAfter("a,b,c", ",")
	fmt.Println(arr5) // [a, b, c]

	// 按照指定字符串切割原字符串, 切割时包含指定字符串, 并且指定切割为几份
	arr6 := strings.SplitAfterN("a,b,c", ",", 2)
	fmt.Println(arr6) // [a, b,c]

	// 按照空格切割字符串, 多个空格会合并为一个空格处理
	arr7 := strings.Fields("a  b c    d")
	fmt.Println(arr7) // [a b c d]

	// 将字符串转换成切片传递给函数之后由函数决定如何切割
	// 类似于IndexFunc
	arr8 := strings.FieldsFunc("a,b,c", custom)
	fmt.Println(arr8) // [a b c]

	// 将字符串切片按照指定连接符号转换为字符串
	sce := []string{"aa", "bb", "cc"}
	str1 := strings.Join(sce, "-")
	fmt.Println(str1) // aa-bb-cc

	// 返回count个s串联的指定字符串
	str2 := strings.Repeat("abc", 2)
	fmt.Println(str2) // abcabc

	// 第一个参数: 需要替换的字符串
	// 第二个参数: 旧字符串
	// 第三个参数: 新字符串
	// 第四个参数: 用新字符串 替换 多少个旧字符串
	// 注意点: 传入-1代表只要有旧字符串就替换
	// 注意点: 替换之后会生成新字符串, 原字符串不会受到影响
	str3 := "abcdefabcdefabc"
	str4 := strings.Replace(str3, "abc", "mmm", -1)
	fmt.Println(str3) // abcdefabcdefabc
	fmt.Println(str4) // mmmdefmmmdefmmm

	//*字符串清理
	// 去除字符串两端指定字符
	str11 := strings.Trim("!!!abc!!!def!!!", "!")
	fmt.Println(str11) // abc!!!def
	// 去除字符串左端指定字符
	str22 := strings.TrimLeft("!!!abc!!!def!!!", "!")
	fmt.Println(str22) // abc!!!def!!!
	// 去除字符串右端指定字符
	str33 := strings.TrimRight("!!!abc!!!def!!!", "!")
	fmt.Println(str33) // !!!abc!!!def
	// // 去除字符串两端空格
	str44 := strings.TrimSpace("   abc!!!def ")
	fmt.Println(str44) // abc!!!def

	// 按照方法定义规则,去除字符串两端符合规则内容
	str5 := strings.TrimFunc("!!!abc!!!def!!!", custom)
	fmt.Println(str5) // abc!!!def
	// 按照方法定义规则,去除字符串左端符合规则内容
	str6 := strings.TrimLeftFunc("!!!abc!!!def!!!", custom)
	fmt.Println(str6) // abc!!!def!!!
	//  按照方法定义规则,去除字符串右端符合规则内容
	str7 := strings.TrimRightFunc("!!!abc!!!def!!!", custom)
	fmt.Println(str7) // !!!abc!!!def

	// 取出字符串开头的指定字符串
	str8 := strings.TrimPrefix("lnj-book.avi", "lnj-")
	fmt.Println(str8) // book.avi

	// 取出字符串结尾的指定字符串
	str9 := strings.TrimSuffix("lnj-book.avi", ".avi")
	fmt.Println(str9) // lnj-book

	//文本取出中间文本
	str := "<html><head><title>Page Title</title></head><body><h1>This is a Heading</h1><p>This is a paragraph.</p></body></html>"
	start := strings.Index(str, "<title>")
	if start == -1 {
		fmt.Println("没有找到<title>")
		return
	}
	end := strings.Index(str, "</title>")
	if end == -1 {
		fmt.Println("没有找到</title>")
		return
	}
	title := str[start+len("<title>") : end]
	fmt.Println(title) // Output: Page Title

	//*正则表达式

}
func custom(r rune) bool {
	fmt.Printf("被调用了, 当前传入的是%c\n", r)
	if r == 'o' {
		return true
	}
	return false
}
