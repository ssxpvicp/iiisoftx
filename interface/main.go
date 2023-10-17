package main

import _case "basis/interface/case"

/*
interface隐式实现
1.golang对象实现interface 无需任何关键词，只需要该对象的方法集中包含接口定义的所有方法且方法签名一致对象实现接口可以借助struct内嵌的特性，实现接口的默认实现
2.类型T方法集包含全部receiverT方法;类型*T方法集包含receiver T + *T方法
3.类型T实例value或pointer可以调用全部的方法，编译器会自动转换
4.类型T实现接口,不管是T还是*T都实现了该接口
5.类型*T实现接口，只有T类型的指针实现了该接口
*/

func main() {
	cat := _case.NewCat()
	animalLife(cat)
	dog := _case.NewDog()
	animalLife(dog)
	dove := _case.NewDove()
	animalLife(dove)
}
func animalLife(a _case.AnimalI) {
	a.Each()
	a.Drink()
	a.Sleep()
	a.Run()

}
