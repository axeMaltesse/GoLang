package main

import (
	"fmt"
	"sync"
)


func main() {

	var g sync.WaitGroup

	var inc = 0
	var gs = 100
	g.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			var inc1 = inc
			inc1++
			inc = inc1
			fmt.Println(inc)
			m.Unlock()
			g.Done()
		}()
	}

	fmt.Println("End Value of Incrementer: \t", inc)
	g.Wait()
}