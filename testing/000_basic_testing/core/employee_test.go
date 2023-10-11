package core_test

import (
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/000_basic_testing/core"
	"testing"
)

func TestValidate_EmptyName(t *testing.T) {

	e := core.Employee{
		Name: "",
	}

	// isValid debe ser false
	// err debe ser != nil
	isValid, err := core.Validate(e)

	if isValid != false {
		t.Error("Expected isValid to be false")
	}
	if err == nil {
		t.Error("Expected err to be not nil")
	}
	if err.Error() != "name cannot be empty" {
		t.Errorf("Expected err.Error() to be 'name cannot be empty', got '%s'", err.Error())
	}

}

func TestValidate_WrongEmail(t *testing.T) {

	e := core.Employee{
		Name:  "Jairo",
		Age:   18,
		Email: "invalidgmail.com",
	}

	// isValid debe ser false
	// err debe ser != nil
	isValid, err := core.Validate(e)

	if isValid != false {
		t.Error("Expected isValid to be false")
	}
	if err == nil {
		t.Error("Expected err to be not nil")
	}
	if err.Error() != "email must contain @" {
		t.Errorf("Expected err.Error() to be 'email must contain @', got '%s'", err.Error())
	}

}

func TestValidate_ValidEmployee(t *testing.T) {

	e := core.Employee{
		Name:  "Jairo",
		Age:   18,
		Email: "invalid@gmail.com",
	}

	// isValid debe ser true
	// err debe ser == nil
	isValid, err := core.Validate(e)

	if isValid != true {
		t.Error("Expected isValid to be true")
	}
	if err != nil {
		t.Error("Expected nil err")
	}

}
