package core_test

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/000_basic_testing/core"
	"testing"
)

func TestValidate_TestTables(t *testing.T) {

	tests := map[string]struct {
		employee        core.Employee
		expectedIsValid bool
		expectedError   error
	}{
		"empty name": {
			employee: core.Employee{
				Name: "",
			},
			expectedIsValid: false,
			expectedError:   fmt.Errorf("name cannot be empty"),
		},

		"wrong email": {
			employee: core.Employee{
				Name:  "Jairo",
				Email: "invalidgmail.com",
				Age:   19,
			},
			expectedIsValid: false,
			expectedError:   fmt.Errorf("email must contain @"),
		},

		"wrong age": {
			employee: core.Employee{
				Name:  "Jairo",
				Email: "invalid@gmail.com",
				Age:   15,
			},
			expectedIsValid: false,
			expectedError:   fmt.Errorf("age must be greater than 18"),
		},
	}

	for testName, subTest := range tests {
		t.Run(testName, func(t *testing.T) {

			gotIsValid, gotErr := core.Validate(subTest.employee)
			if gotIsValid != subTest.expectedIsValid {
				t.Errorf("Expected gotIsValid to be %v, got %v", subTest.expectedIsValid, gotIsValid)
			}
			if gotErr.Error() != subTest.expectedError.Error() {
				t.Errorf("Expected gotErr to be %s, got %s", subTest.expectedError.Error(), gotErr.Error())
			}

		})
	}

}

func TestValidate(t *testing.T) {

	t.Run("empty name", func(t *testing.T) {
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
	})

	t.Run("wrong email", func(t *testing.T) {
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
	})

	t.Run("valid employee", func(t *testing.T) {
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
	})

}
