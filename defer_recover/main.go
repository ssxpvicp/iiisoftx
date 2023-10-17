package main

import _case "basis/defer_recover/case"

/*
应用结题
1.资源释放
2.异常捕获和处理

#defer
1.defer 关键词用来声明一个延迟调用函数,该函数可以是匿名函数也可以是具名函数
2.defer 延时函数执行时间(位置),在方法return之后,返回参数到调用方法之前
3.defer 延时函数可以在方法返回之后改变函数的返回值
4.defer 在方法结束(正常返回,异常结束)都会去调用defer声明的延时函数,可以有效避免因异常导致的资源无法释放的问题
5.defer 可以指定多个defer 延时函数,多个延时函数执行顺序为后进先出,同栈一样.
6.defer 通常用于资源释放 异常捕获等场景,例如:关闭连接,关闭文件等
7.defer 当defer和recover配合,可以实现异常捕获与处理逻辑
8.defer 不建义在for循环中使用defer

#recover
1.Go语言的内建函数,可以让进入宕机流程中的. goroutine 恢复过来
2.recover 仅在延迟函数 defer 中有效,在正常的执行过程中.调用 recover 会返回nil 并且没有其它任何效果
3.如果当前的 goroutine 出现panic , 调用 recover 可以捕获到 panic 的输入值,并且恢复正常的执行

#panic
1.Go语言的一种异常机制
2.可通过panic 函数主动抛出异常
*/

func main() {
	_case.DeferCase1()
	_case.DeferCase2()
	_case.DeferCase3()
	//_case.ExceptionCase()
	_case.FileReadCase()
}
