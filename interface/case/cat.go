package _case

import "fmt"

type Cat struct {
	animal
	AAA int
}

func NewCat() AnimalI {
	return &Cat{}
}
func (c *Cat) Each() {
	fmt.Println("猫吃老鼠")
}
