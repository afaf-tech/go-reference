package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func (c Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is ", c.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huuuuuu from ", a.Name)
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
	akmal.sayHi("lola")

	fmt.Println("==============")
	fmt.Println(fikri.Name)
	fmt.Println(fikri.Address)
	fmt.Println(fikri.age)
	fikri.sayHuu()
}
