package main

import (
	"fmt"
	"time"
)

func table(n, i int) {
	if i == 0 {
		return
	}

	table(n, i-1)
	time.Sleep(time.Second)
	fmt.Println(n, " x ", i, " = ", i*n)

}
func processing(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println("Processing.......")
	}
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	time.Sleep(time.Second)
	return n * fact(n-1)
}
func main() {

	table(2, 10)

	go table(5, 10)
	// go table(3)
	go fmt.Println(fact(5))
	go func(str string) {
		fmt.Print(str)
	}("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum")
	go processing(5)
	fmt.Scanln()

}

/*
-----Output-----
2  x  1  =  2
2  x  2  =  4
2  x  3  =  6
2  x  4  =  8
2  x  5  =  10
2  x  6  =  12
2  x  7  =  14
2  x  8  =  16
2  x  9  =  18
2  x  10  =  20
5  x  1  =  5
5  x  2  =  10
5  x  3  =  15
5  x  4  =  20
5  x  5  =  25
120
5  x  6  =  30
Processing.......
Processing.......
5  x  7  =  35
5  x  8  =  40
Processing.......
Processing.......
5  x  9  =  45
5  x  10  =  50
Processing.......
*/
