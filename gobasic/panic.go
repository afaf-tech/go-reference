package main

import "fmt"

func endApp() {
	// will revover the panic cathc. and still continue the application process afterwards.
	message := recover()
	if message != nil {
		fmt.Println("\n Error dengan message: ", message)
	}
	fmt.Println("\n Applikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		// will stop all execution after
		panic("Applikasi Error")
	}
	fmt.Println("Applikaasi berjalan")
}

func main() {
	runApp(true)
}
