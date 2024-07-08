package main

import "fmt"

func main() {
	var a byte = 1 // alias of uint8
	var b rune = 1 // alias of int32
	var c int = 1  // alias of int32
	var d uint = 1 // alias of uint32

	fmt.Println(a, b, c, d)
}
