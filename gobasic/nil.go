package main

import "fmt"

// nil just works in interface, function, map, slice, pointer and channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	person := NewMap("akm")
	fmt.Println(person)
}
