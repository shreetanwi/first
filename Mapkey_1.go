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
		first: "James ",
		last:  " Bond",
		favflavors: []string{
			"Chocolate",
			"Strawberry",
		},
	}
	p2 := person{
		first: "Miss ",
		last:  " Monneypenny",
		favflavors: []string{
			"Vanilla",
			"Pineapple",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p1.last: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favflavors {
			fmt.Println(i, val)
		}
	}
}
