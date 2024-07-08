package main

import "fmt"

func main() {
	var a string = "Hello World"
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a[0])
	fmt.Println(a[0])
	fmt.Println(a[1:4])
	fmt.Println(a[:4])
	fmt.Println(a[4:])
	fmt.Println(a[len(a)-1:])
	fmt.Println(a[:len(a)-1])
}
