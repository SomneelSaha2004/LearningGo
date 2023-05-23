package main

import (
	"fmt"
)

type ZeroDivisionError struct {
	a, b float64
}

func (z *ZeroDivisionError) Error() string {
	return fmt.Sprintf("Cannot divide by zero : %v/%v", z.a, z.b)
}
func divide(lhs, rhs float64) (float64, error) {
	if rhs == 0 {
		return 0, &ZeroDivisionError{lhs, rhs}
	} else {
		return lhs / rhs, nil
	}
}
func main() {
	a, err := divide(3, 0)
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err)
	}
}
