package math

import "errors"

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	if a == 0 {
		return 0, nil
	}

	return a / b, nil
}
