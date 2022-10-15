package govalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func TestValidati(t *testing.T) {
	if validate == nil {
		t.Error("Validate is nil")
	}

}

func TestValidationVariable(t *testing.T) {
	var user string = ""
	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariables(t *testing.T) {
	password := "password"
	confirmPassword := "password"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestValidationMultipleTags(t *testing.T) {
	var user string = "fdsaf4"
	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTagParemeter(t *testing.T) {
	var user string = "000000"
	err := validate.Var(user, "required,numeric,min=3,max=3")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	fikri := LoginRequest{Username: "Fikri", Password: "nana"}
	err := validate.Struct(fikri)
	if err != nil {
		fmt.Println(err.Error())
	}
}
