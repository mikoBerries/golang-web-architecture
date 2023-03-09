package main

import "fmt"

type meow int

func (m meow) error() string {
	return fmt.Sprint("error from meow ", m)
}
func main() {
	var m meow = 1
	fmt.Println(m.error())
}
