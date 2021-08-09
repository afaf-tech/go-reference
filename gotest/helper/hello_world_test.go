package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fani")
	require.Equal(t, "Hello Fanif", result, "Result must be 'Hello Fani'")
	fmt.Println("Test HelloWorld with Require") // will not be printed // call failnow
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fani")
	assert.Equal(t, "Hello Fani", result, "Result must be 'Hello Fani'")
	fmt.Println("Test HelloWorld with Assert") // will be printed if failed -- call fail
}
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
