package main

import (
	"fmt"
	"customMath/constant"
	"customMath/simpleoperations"

)

func main(){
	fmt.Println(constant.PI)
	fmt.Println(simpleoperations.Add(2,3))
	fmt.Println(simpleoperations.Add(2,3,5))
	fmt.Println(simpleoperations.Divide(8,2))
	fmt.Println(simpleoperations.Divide(8,0))
	fmt.Println(simpleoperations.Multiply(2,4))
}
/*---Output---
3.14
5
10
4
2022/06/08 12:14:47 Warning: Division by zero
0
8
*/