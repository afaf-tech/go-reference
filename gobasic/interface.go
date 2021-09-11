package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(name HasName) {
	fmt.Printf("Hello %s\n", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	var fikri Person
	fikri.Name = "John"

	sayHello(fikri)

	cat := Animal{
		Name: "Puss",
	}
	sayHello(cat)
}
