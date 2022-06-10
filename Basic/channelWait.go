package main
import (
	"fmt"
	"time"
)
func processing(n int,done chan bool) {
	
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println("Processing.......")
	}
	done<-true
}
func main(){
	done:=make(chan bool)
	go processing(5,done)
	<-done
}