package main

import "fmt"

type hotdog int

func (h hotdog) Cook() {
	fmt.Println("Hotdog has been cooked")
}

type hotFood interface {
	Cook()
}

func servingFood(hf hotFood) {
	hf.Cook()
	fmt.Println("serving hot hotdog")
}

func main() {
	var x hotdog = 42
	var y hotFood
	y = x
	fmt.Printf("%T\n", y)

	servingFood(x)
}
