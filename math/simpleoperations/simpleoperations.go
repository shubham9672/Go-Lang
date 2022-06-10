package simpleoperations

import (
	"log"
)

func Add(variable ...int) int{
	var c int = 0
	for _, v := range variable {
		c += v
	}
	return c
}
func Multiply(a int,b int) int{
	return a*b
}
func Divide(quotient int,divident int) int{
	if divident==0{
		log.Println("Warning: Division by zero")
		return 0
	}
	return quotient/divident
}
