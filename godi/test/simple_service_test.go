package test

import (
	"afaf-tech/belajar-golang-restful-api/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService()
	fmt.Println(err)
	if err == nil {
		fmt.Println(simpleService.SimpleRepository)

	} else {
		fmt.Println(simpleService)
	}
}
