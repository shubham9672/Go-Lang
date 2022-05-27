package main

import (
	"fmt"
	"time"
)

func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4}
	ch := make(chan int, len(arr))
	go timesThree(arr, ch)
	time.Sleep(time.Second) //  let the goroutine return all the values
	for i := 0; i < len(arr); i++ {
		result := <-ch
		fmt.Printf("The result is: %v\n", result)
	}

}
