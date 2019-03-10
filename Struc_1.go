package main

import (
	"fmt"
)

type person struct {
	first, last string
	age         int
}

func main() {

	p1 := person{
		first: "Jmes ",
		last:  "bond",
		age:   32,
	}

	p2 := person{
		first: "Miss ",
		last:  " Monneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
