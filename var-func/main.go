package main

import (
	_case "basis/var-func/case"
	"fmt"
)

func main() {
	var a = 10
	var b = 20
	_case.SumCase(a, b)
	fmt.Println(a, b)
	_case.ReferenceCase(a, &b)
	fmt.Println(a, b)

	fmt.Println(_case.G)
	_case.ScopeCase(a, b)
	fmt.Println(_case.G)

	user := _case.NewUser("nick", 18)

	fmt.Println(user.GetName(), user.GetAge())
	_case.SimpleCase()
	_case.CusNumTCase()
	_case.BuiltInCase()
	_case.TTypeCase()
	_case.TTypeCase1()
	_case.ReceiverSase()

	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	//defer stop()
	//<-ctx.Done()
}
