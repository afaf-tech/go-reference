package govalidation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidati(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}
