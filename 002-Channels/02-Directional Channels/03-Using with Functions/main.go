package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)
	bar(c)

}

// Send it to the channel
func foo(c chan<- int) {
	c <- 42
}

//Receive from the channel
func bar(c <-chan int) {
	fmt.Println(<-c)
}
