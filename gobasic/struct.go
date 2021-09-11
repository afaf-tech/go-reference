package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func main() {
	akmal := Customer{}
	akmal.Name = "akmal"
	akmal.Address = "njajong"
	akmal.age = 16
	// literal instantiation
	fikri := Customer{Name: "Fikri", Address: "amrik", age: 21}

	fmt.Println(akmal.Name)
	fmt.Println(akmal.Address)
	fmt.Println(akmal.age)
	fmt.Println(fikri.Name)
	fmt.Println(fikri.Address)
	fmt.Println(fikri.age)
}
