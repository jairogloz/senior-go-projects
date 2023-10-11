package math_test

import (
	"errors"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/chatgpt/codereview/math"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b        int
		expected    int
		expectedErr error
	}{
		{10, 2, 5, nil},   // Caso normal
		{-10, 2, -5, nil}, // Numerador negativo
		{10, -2, -5, nil}, // Denominador negativo
		{-10, -2, 5, nil}, // Numerador y denominador negativos
		{0, 2, 0, nil},    // Numerador es cero
		{10, 0, 0, errors.New("cannot divide by zero")}, // Denominador es cero
	}

	for _, tt := range tests {
		result, err := math.Divide(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("divide(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
		if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
			t.Errorf("divide(%d, %d) error = %s; expected %s", tt.a, tt.b, err.Error(), tt.expectedErr.Error())
		}
	}
}
