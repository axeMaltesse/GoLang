package main

import (
	"fmt"
	"sync"
)

var wb sync.WaitGroup

func main() {

	wb.Add(2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Counter 1: \t", i)
		}
		wb.Done()
	}()

	go func() {

		for i := 0; i < 5; i++ {
			fmt.Println("Counter 2: \t", i)
		}
		wb.Done()
	}()
	wb.Wait()
}
