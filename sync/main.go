package main

import (
	_case "basis/sync/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	//_case.WaitGroupCase()
	//_case.WaitGroupCase1()
	//_case.CondCase()
	//_case.CondQueueCase()
	//_case.MutexCase()
	//_case.MapCase()
	//_case.MapCase1()
	//_case.PoolCase()
	//_case.OnceCase()
	//_case.AtomicCase()
	//_case.AtomicCase1()
	_case.AtomicCase2()
	//设置手动停止主程序
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
