package main

import (
	"fmt"
	"sync"
)

func tablePrinter(wg *sync.WaitGroup, value chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(<-value)
	}
}
func tableCalculation(wg *sync.WaitGroup, value chan int, n int) {
	defer wg.Done()
	for i := 1; i < 11; i++ {
		value <- i * n
	}
}

func main() {
	fmt.Println("---go routienes using sync waitGroup---")
	wg := new(sync.WaitGroup)
	wg.Add(2)
	value := make(chan int)
	defer close(value)
	go tableCalculation(wg, value, 5)
	go tablePrinter(wg, value)
	wg.Wait()
}
