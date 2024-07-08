package main

import "fmt"

func main() {
	// const adalah sebuah variabel yang tidak dapat diubah
	const (
		firstName = "Alfian"
		lastName  = "Chabib"
	)
	fmt.Println(firstName, lastName)

	// error
	// firstName = "Alfian"
	// lastName = "Chabib"
}
