package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := &Customer{
		FirstName:  "Fatih",
		MiddleName: "Al",
		LastName:   "Fikri",
		Age:        10,
		Married:    false,
	}
	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%v]\n", string(byte))
}
