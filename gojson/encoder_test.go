package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// decoder simply decode a JSON string into a struct without saving it in any additional variables
func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_out.json")
	decoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "John",
		Age:       24,
		Hobbies:   []string{"Doe", "Volley", "Futsal"},
		Addresses: []Address{
			{
				Street:     "Jl Makam Pahlawan",
				Country:    "Kosovo",
				PostalCode: "0",
			},
			{
				Street:     "Jl Kusuma Bangsa",
				Country:    "Built",
				PostalCode: "43344",
			},
		},
	}
	decoder.Encode(customer)

	fmt.Println(customer)
}
