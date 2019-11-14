package main

import (
	"fmt"
)

const goroutines int = 10

func main() {

	q:=make(chan int)

	c := gen(q)

	receiver(c)
}

func receiver(c <-chan int) {
	for v:=0; v<100;v++{
		fmt.Println("Printing from channel: \t",<-c)
	}
}

func gen(c chan int) <-chan int{
	for v:=0; v < goroutines; v++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}
	return c
}