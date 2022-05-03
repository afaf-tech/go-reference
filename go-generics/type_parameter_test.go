package go_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T { // any is type constraints
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result = Length[string]("Eko")
	assert.Equal(t, "Eko", result)
	var resultNum = Length[int](100)
	assert.Equal(t, 100, resultNum)
	assert.True(t, true)
}
