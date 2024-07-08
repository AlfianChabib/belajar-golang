package main

import "fmt"

func main() {
	var name string
	name = "Alfian"
	fmt.Println(name)
	name = "Chabib"
	fmt.Println(name)

	firstName := "Alfian"
	lastName := "Chabib"
	fmt.Println(firstName, lastName)

	firstName = "Chabib"
	lastName = "Alfian"
	fmt.Println(firstName, lastName)

	var (
		namaPertama  = "Alfian"
		namaTerakhir = "Chabib"
	)
	fmt.Println(namaPertama, namaTerakhir)
}
