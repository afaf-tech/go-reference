package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	logJson("Fikri")
	logJson(1)
	logJson(true)
	logJson([]string{"fatih", "al", "fikri"})
}
