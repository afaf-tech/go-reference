package main

import "fmt"

func random() interface{} {
	return 323
}

func main() {
	/* 	var result interface{} = random()
	   	var resultString string = result.(string)
	   	fmt.Println(resultString) */

	resultSwitch := random()

	switch value := resultSwitch.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
