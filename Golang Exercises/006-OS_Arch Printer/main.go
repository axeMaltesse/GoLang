package main

import (
	"fmt"
	"runtime"
)

func main() {
	const GOOS string = runtime.GOOS
	const GOARCH string = runtime.GOARCH

	fmt.Println("GOARCH is: \t",GOARCH)
	fmt.Println("GOOS is: \t",GOOS)
}