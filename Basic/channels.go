package main

import "fmt"

func fact(n int, res chan<- int) {
	factt := 1
	for i := n; i > 0; i-- {
		factt *= i
	}
	res <- factt
}
func main() {
	//created channel
	res := make(chan int)
	go fact(5, res)
	fmt.Println(<-res)
	/*
		Output
		120
	*/
}
