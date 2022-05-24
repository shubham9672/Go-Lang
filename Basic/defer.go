package main

import (
	"bufio"
	"fmt"
	"os"
)

func hi() {
	fmt.Println("Hello From Go")
}
func end() {
	fmt.Print("Thank You Bye")
}

func main() {
	//Using Defer keyword it will print at the end of the execution
	defer end()
	//Print Hello from Go
	hi()
	// Taking input from user
	fmt.Println("Please Enter your name:")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	//Finally Print the input
	fmt.Print("Hello ", input)
	/*---------Output-------*/
	/*Hello From Go
	Please Enter your name:
	Shubham
	Hello Shubham
	Thank You Bye*/
}
