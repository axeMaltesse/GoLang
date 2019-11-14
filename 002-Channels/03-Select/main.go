package main

import (
	"fmt"
)

func main()  {
	even:= make(chan int)
	odd:= make(chan int)
	quit := make(chan int)

	//send
	go send(even,odd,quit)

	//receive
	receive(even,odd,quit)

	fmt.Println("About to exit")
}

func send(e,o,q chan <- int)  {
	for i:=0; i<100; i++{
		if i % 2 == 0 {
			e <- i
		} else if i % 2 != 0 {
			o <- i
		} else {
			q <- 0
		}
	}
	close(e)
	close(o)
	close(q)
}

func receive(e,o,q <- chan int){
	for {
		select {
		case v := <-e:
			fmt.Println("From the Even channel:\t", v)
		case v := <-o:
			fmt.Println("From the Odd channel:\t", v)
		case v := <-q:
			fmt.Println("From the Quit channel:\t", v)
			return
		}
	}
}