package main

import (
	"fmt"
	"runtime"
	"sync"
)

var inc int
var g = sync.WaitGroup{}

func main() {

	g.Add(3)

	go func() {
		runtime.Gosched()
		inc1 := inc
		for i := 0; i < 4; i++ {
			inc1++
			inc = inc1
			fmt.Println(inc)
		}
		g.Done()
	}()

	go func() {
		runtime.Gosched()
		inc1 := inc
		for i := 0; i < 4; i++ {
			inc1++
			inc = inc1
			fmt.Println(inc)
		}
		g.Done()
	}()

	go func() {
		inc1 := inc
		for i := 0; i < 4; i++ {
			inc1++
			runtime.Gosched()
			inc = inc1
			fmt.Println(inc)
		}
		g.Done()
	}()

	g.Wait()
}
