package test

import (
	"afaf-tech/belajar-golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceErr(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
