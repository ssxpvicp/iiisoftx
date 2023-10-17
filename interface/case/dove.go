package _case

import "fmt"

type Dove struct {
}

func NewDove() AnimalI {
	return &Dove{}
}

func (a *Dove) Each() {
	fmt.Println("鸽子吃肉包子")
}
func (a *Dove) Drink() {
	fmt.Println("鸽子喝白开水")
}
func (a *Dove) Sleep() {
	fmt.Println("鸽子睡觉")
}
func (a *Dove) Run() {
	fmt.Println("鸽子运动")
}
