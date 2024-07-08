package main

import "fmt"

func main() {
	// integer
	// int8 int16 int32 int64
	// uint8 uint16 uint32 uint64
	var a int8 = 127                  // -128 ~ 127
	var b int16 = 32767               // -32768 ~ 32767
	var c int32 = 2147483647          // -2147483648 ~ 2147483647
	var d int64 = 9223372036854775807 // -9223372036854775808 ~ 9223372036854775807

	fmt.Println(a, b, c, d)

	var e uint = 255                    // 0 ~ 255
	var f uint8 = 255                   // 0 ~ 255
	var g uint16 = 65535                // 0 ~ 65535
	var h uint32 = 4294967295           // 0 ~ 4294967295
	var i uint64 = 18446744073709551615 // 0 ~ 18446744073709551615

	fmt.Println(e, f, g, h, i)
}
