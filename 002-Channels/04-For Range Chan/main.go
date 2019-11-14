package main

import "fmt"

func main() {

	q:=make(chan int)

	c := gen(q)

	receiver(c)
}

func receiver(c <-chan int) {
	for v := range c {
		fmt.Println("Printing from channel: \t",v)
	}
}

func gen(c chan int) <-chan int{
	go func() {
		for i:=0; i<100;i++{
			c <- i
		}
		close(c)
	}()
	return c
}