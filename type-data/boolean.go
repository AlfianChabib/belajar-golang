package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	c := !a
	d := !b

	fmt.Println(a, b, c, d)

	if a {
		fmt.Println("a is true")
	} else {
		fmt.Println("a is false")
	}

	if b {
		fmt.Println("b is true")
	} else {
		fmt.Println("b is false")
	}
}
