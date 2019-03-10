package main

import (
	"fmt"
)

type person struct {
	first, last string
}

func main() {

	p1 := person{
		first: "Tanwi",
		last:  "Shree",
	}

	p2 := person{
		first: "Ashish",
		last:  "Sinha",
	}
	fmt.Println(p1, p2)

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	for _, V := range m {
		fmt.Println(V.first)
		fmt.Println(V.last)
	}
}
