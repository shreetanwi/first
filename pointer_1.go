package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Tanwi Shree",
	}

	fmt.Println(p1)
	ChangeMe(&p1)
	fmt.Println(p1)
}

func ChangeMe(p *person) {
	p.name = "James Bond"
	(*p).name = "Monneyp"
}
