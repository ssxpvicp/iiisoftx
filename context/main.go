package main

import (
	_case "basis/context/case"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	//_case.ContextCase()
	_case.GPT例子()
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
	fmt.Println(111)
}
