package main

import (
	"fmt"
)

func main() {

	var a, b, c int
	a = 50
	b = 20
	c = 30

	x := a + b - c
	y := a * b
	z := float64(a) / float64(b)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
