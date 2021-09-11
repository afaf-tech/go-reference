package main

import "fmt"

func logging() {
	fmt.Println("finish calling logging")
}

func runApplication(value int) {
	// function will always be executed error or not
	defer logging()
	fmt.Println("start application")
	result := 10 / value
	fmt.Println("result: ", result)
}

func main() {
	runApplication(0)
}
