package main

import (
	"fmt"
	"time"
)

func f1(c1 chan string) {
	time.Sleep(1 * time.Second)
	c1 <- "Processed Function 1"
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go f1(c1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Processed anonymous function"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
