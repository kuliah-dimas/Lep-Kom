package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f_name string = "Dimas"
	var l_name string = "Febriyanto"
	var age int = 18

	fmt.Println("My name is " + f_name + " " + l_name + ". I am " + strconv.Itoa(age) + " old")
}
