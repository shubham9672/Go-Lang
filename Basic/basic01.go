package main

import (
	"fmt"
	"math"
	"time"
)

func basic01() {
	// var stringVar = "Shubham"
	// dynamicVar := "Hello, Variable"
	// fmt.Printf("String Var %s\n", stringVar)
	// fmt.Println(dynamicVar)
	// var i uint8 = 255
	var x, j = 3, 127
	var hexaDecimal = 0x3b
	var address = &hexaDecimal
	// fmt.Println(i+3, i, i-3)
	// fmt.Println(j+3, j, j-3)
	// fmt.Printf("The type and value of String var %v,%T\n", stringVar, stringVar)
	// fmt.Printf("The type and value of dynamic var %v,%T\n", dynamicVar, dynamicVar)
	// fmt.Printf("The type and value of i %v,%T\n", i, i)
	// fmt.Printf("The type and value of x %v,%T\n", x, x)
	// fmt.Printf("The type and value of j %v,%T\n", j, j)
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// fmt.Printf("The type and value of hexaDecimal %v,%T\n", hexaDecimal, hexaDecimal)
	fmt.Println(x + hexaDecimal)
	fmt.Println(j % x)
	fmt.Println(j & x)
	fmt.Println(j ^ x)
	fmt.Println(j << 1)
	fmt.Println(*address)
	// Constants
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	//Controll Statements
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 10; j++ {
		for k := 0; k <= j; k++ {
			fmt.Print(k + 1)
		}
		fmt.Println()
	}
	for j := 1; j < 100; j++ {
		if j%3 == 0 {
			fmt.Println(j, " is divided by 3")
		} else if j%5 == 0 {
			fmt.Println(j, " is divided by 5")
		} else if j%7 == 0 {
			fmt.Println(j, " is divided by 7")
		}
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	oddEven := func(i int) {
		switch i % 2 {
		case (0):
			fmt.Println(i, " is Even")
		case (1):
			fmt.Println(i, " is Odd")
		default:
			fmt.Println(i, " Don't know may be -1")
		}
	}
	for j := -2; j < 10; j++ {
		oddEven(j)
	}
}
