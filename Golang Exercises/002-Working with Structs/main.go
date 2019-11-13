package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	saySomething(&p1)

}

func (p *person) speak() {
	fmt.Println("I am saying smth")
}

func saySomething(h human)  {
	h.speak()
}