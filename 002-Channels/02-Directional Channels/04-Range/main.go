package main

import "fmt"

func main() {
	c := make(chan int)

	// Send it to the channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}
