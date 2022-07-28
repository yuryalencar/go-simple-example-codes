package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	yury := person{firstName: "Yury", lastName: "Alencar"}

	fmt.Println("Struct:", yury)
	fmt.Println("First name:", yury.firstName)
	fmt.Println("Last name:", yury.lastName)
}
