package main

import "fmt"

func main() {
	var a float32 = 1.0                  // 1.18x10^-38 ~ 3.34x10^38
	var b float64 = 3.0                  // 2.23x10^-308 ~ 1.80x10^308
	var c complex64 = complex(1.0, 2.0)  // complex number with float32 real and imaginary parts
	var d complex128 = complex(3.0, 4.0) // complex number with float64 real and imaginary parts

	fmt.Println(a, b, c, d)
}
