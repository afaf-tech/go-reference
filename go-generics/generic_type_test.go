package go_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[E any](bag Bag[E]) {
	for _, e := range bag {
		fmt.Println(e)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Eko", "Kurniawan", "Khannedy"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 212, 212, 21}
	PrintBag(numbers)
}
