package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Ekos" {
		// the next code is still executed
		//t.Fail()
		// t.Error give extra information about the error
		t.Error("Result must be Eko")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldFikri(t *testing.T) {
	result := HelloWorld("Fikri")

	if result != "Hello Eko" {
		// the next code will not be executed
		//t.FailNow()
		// t.Fatal give an extra information about the error
		t.Fatal("Result must be Fikri")
	}

	fmt.Println("TestHelloWorldFikri done")

}
