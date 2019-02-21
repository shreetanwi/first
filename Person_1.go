package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favflavors []string
}

func main() {

	p1 := person{
		first: "Tanwi ",
		last:  " Shree",
		favflavors: []string{
			"Chocolate",
			"Rum coke",
		},
	}
	p2 := person{
		first: "Miss ",
		last:  " Mona",
		favflavors: []string{
			"Vanilla",
			"Pineapple",
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favflavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favflavors {
		fmt.Println(i, v)
	}
}
