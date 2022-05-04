package go_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// generic struct
type Data[T any] struct {
	First  T
	Second T
}

// generic methods
// methods can't have any generic Parameter
// its generic type is from struct inherited
func (d *Data[_]) SayHello(name string) string { // _ operator to ignore T in generic Data
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Fikri",
		Second: "Luv",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Fikri",
		Second: "Luv",
	}

	// generic already been declared on struct instance creation above
	assert.Equal(t, "Fikri", data.ChangeFirst("Fikri"))
	assert.Equal(t, "Hello Fikri", data.SayHello("Fikri"))
}
