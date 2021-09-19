package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
		Age:       32,
		Married:   true,
		Hobbies:   []string{"nana", "volley", "futsal"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{ "FirstName": "John", "LastName": "Doe", "Age": 12, "Married":true, "Hobbies": ["nana", "volley", "futsal"] }`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "John",
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

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{ "FirstName": "John", "Addresses": [{"Street": "Jl Makam Pahlawan", "Country": "Kosovo", "PostalCode": "0"	},{ "Street":"Jl Kusuma Bangsa","Country": "Built", "PostalCode": "43344"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArray(t *testing.T) {
	jsonString := `[{"Street": "Jl Makam Pahlawan", "Country": "Kosovo", "PostalCode": "0"	},{ "Street":"Jl Kusuma Bangsa","Country": "Built", "PostalCode": "43344"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
