package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

func main() {

	fikri := Man{Name: "Fikri"}
	fikri.Married()
	fmt.Println(fikri)
}
