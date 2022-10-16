package govalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	fikri := LoginRequest{Username: "Fikri", Password: "nana"}
	err := validate.Struct(fikri)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println("Field error:", err.Field(), "on tag", err.Tag(), "error", err.Error())
		}
	}
}

func TestValidationCrossField(t *testing.T) {
	type LoginRequest struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	fikri := LoginRequest{Username: "Fikri", Password: "fafafaf", ConfirmPassword: "fafafaf"}
	err := validate.Struct(fikri)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println("Field error:", err.Field(), "on tag", err.Tag(), "error", err.Error())
		}
	}
}

func TestValidationNestedStruct(t *testing.T) {
	type Address struct {
		City   string `validate:"required"`
		Street string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	akira := &User{
		Id:   "32",
		Name: "akira",
	}

	err := validate.Struct(akira)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println("Field error:", err.Field(), "on tag", err.Tag(), "error", err.Error())
		}
	}
}

func TestValidationDive(t *testing.T) {
	type Address struct {
		City   string `validate:"required"`
		Street string `validate:"required"`
	}

	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
	}

	akira := &User{
		Id:   "32",
		Name: "akira",
		Address: []Address{
			{
				City:   "akira",
				Street: "jl veteran",
			},
			{
				City:   "",
				Street: "",
			},
		},
	}

	err := validate.Struct(akira)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println("Field error:", err.Field(), "on tag", err.Tag(), "error", err.Error())
		}
	}
}

func TestValidationBasicColletion(t *testing.T) {
	type Address struct {
		City   string `validate:"required"`
		Street string `validate:"required"`
	}

	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobbies []string  `validate:"dive,required,min=3"` // hobbies is not required because of no "required" at the first order
	}

	akira := &User{
		Id:   "32",
		Name: "akira",
		Address: []Address{
			{
				City:   "akira",
				Street: "jl veteran",
			},
			{
				City:   "",
				Street: "",
			},
		},
		Hobbies: []string{"na", "nilna"},
	}

	err := validate.Struct(akira)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println("Field error:", err.Field(), "on tag", err.Tag(), "error", err.Error())
		}
	}
}

func TestValidationMap(t *testing.T) {
	type Address struct {
		City   string `validate:"required"`
		Street string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id      string            `validate:"required"`
		Name    string            `validate:"required"`
		Address []Address         `validate:"required,dive"`
		Hobbies []string          `validate:"dive,required,min=3"`                   // hobbies is not required because of no "required" at the first order
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"` // keys for key validations, last dive is for School validation
		Wallet  map[string]int    `validate:"dive,keys,required,min=2,endkeys,required,gt=0"`
	}

	akira := &User{
		Schools: map[string]School{
			"SD": {
				Name: "Mi Maarif tanqul",
			},
			"SMP": {
				Name: "Mts Maarif tanqul",
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"GOPAY": 10,
			"Dana":  -1,
		},
	}

	err := validate.Struct(akira)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println("Field error:", err.Field(), "on tag", err.Tag(), "error", err.Error())
		}
	}
}

func TestValidationALias(t *testing.T) {
	validate.RegisterAlias("varchar", "required,max=10")

	type Seller struct {
		Name    string `validate:"varchar"`
		Address string `validate:"varchar"`
	}

	seller1 := Seller{
		Name:    "seller1",
		Address: "nonafdfsadfasfdsfds",
	}
	err := validate.Struct(seller1)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println("Field error:", err.Field(), "on tag", err.Tag(), "error", err.Error())
		}
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidationFunction(t *testing.T) {
	validate.RegisterValidation("username", MustValidUsername)
	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required,min=5"`
	}

	fikri := LoginRequest{Username: "FIKRI", Password: "nana"}
	err := validate.Struct(fikri)
	if err != nil {
		fmt.Println(err.Error())
	}

}

var regexNumber = regexp.MustCompile(`^[0-9]+$`)

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return length == len(value)
}

func TestCustomValidationParameter(t *testing.T) {
	validate.RegisterValidation("pin", MustValidPin)

	type LoginRequest struct {
		Password string `validate:"required,min=5"`
		Pin      string `validate:"required,pin=4"`
	}

	fikri := LoginRequest{Password: "nana", Pin: "9094"}
	err := validate.Struct(fikri)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	validate.RegisterValidation("pin", MustValidPin)

	type LoginRequest struct {
		Username string `validate:"required,email|numeric"`
	}

	fikri := LoginRequest{Username: "alfa"}
	err := validate.Struct(fikri)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()

	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue

}

func TestCrossFieldValidation(t *testing.T) {
	validate.RegisterValidation("feic", MustEqualIgnoreCase) // feic: field equal ignore case (abbreviated)

	type User struct {
		Username string `validate:"required,feic=Email"`
		Email    string `validate:"required,email"`
	}

	fikri := User{Username: "feic@example", Email: "feic@example.com"}
	err := validate.Struct(fikri)
	if err != nil {
		fmt.Println(err.Error())
	}
}
