package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main() {

	var g sync.WaitGroup
	var inc int64
	var gs = 100
	g.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&inc,1)
			fmt.Println(atomic.LoadInt64(&inc))
			g.Done()
		}()
	}
	g.Wait()
	fmt.Println("End Value of Incrementer: \t", inc)
}