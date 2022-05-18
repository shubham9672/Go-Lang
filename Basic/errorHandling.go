package main

import (
	"errors"
	"fmt"
	"math"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}
func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

type argError struct {
	arg  float64
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func Sqrt1(value float64) (float64, error) {
	if value < 0 {
		return -1, &argError{value, "Math: negative number passed to Sqrt"}
	}
	return math.Sqrt(value), nil
}

func main() {
	for k, i := range map[int]int{1: 12, 0: 12} {
		if r, e := divide(i, k); e != nil {
			fmt.Println("error", e)
		} else {
			fmt.Println("result", r)
		}
	}
	for _, i := range []float64{9, -9} {
		if r, e := Sqrt(i); e != nil {
			fmt.Println("error:", e)
		} else {
			fmt.Println("result:", r)
		}
	}
	_, e := Sqrt1(-4)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
