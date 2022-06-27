package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return a / b, nil
}
