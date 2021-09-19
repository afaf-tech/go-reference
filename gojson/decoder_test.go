package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// decoder simply decode a JSON string into a struct without saving it in any additional variables
func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
