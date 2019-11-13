package main

import (
	"fmt"
	"runtime"
	"sync"
)

var g = sync.WaitGroup{}

func main() {

	var inc = 0
	var gs = 100
	g.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			runtime.Gosched()
			var inc1 = inc
			inc1++
			inc = inc1
			fmt.Println(inc)
			g.Done()
		}()
	}
	fmt.Println("End Value of Incrementer: \t",inc)
	g.Wait()
}
