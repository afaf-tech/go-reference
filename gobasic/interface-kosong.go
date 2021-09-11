package main

import "fmt"

// it is like any in typescript
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	}
	if i == 2 {
		return true
	}

	return "Ups"
}

func main() {
	data := Ups(3)
	fmt.Println(data)
}
