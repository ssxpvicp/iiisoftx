package _case

import "fmt"

// AnimalI 声音AnimalI接口
// 定义AnimalI行为
type AnimalI interface {
	// Each 吃
	Each()
	// Drink 喝
	Drink()
	// Sleep 睡觉
	Sleep()
	// Run 奔跑
	Run()
}

type animal struct {
}

func (a animal) Each() {
	fmt.Println("Animal Each 默认实现")
}
func (a animal) Drink() {
	fmt.Println("Animal Drink 默认实现")
}
func (a animal) Sleep() {
	fmt.Println("Animal Sleep 默认实现")
}
func (a animal) Run() {
	fmt.Println("Animal Run 默认实现")
}

func init() {
	a := animal{}
	a.Drink()
	a.Each()
}
