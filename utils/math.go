package utils

import "errors"

func sum(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't be divide by 0")
	}
	return a / b, nil
}
