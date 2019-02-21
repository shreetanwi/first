package main

import (
	"fmt"
)

func main() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{

		first: "James",
		friends: map[string]int{
			"Monneypenny": 555,
			"Q":           777,
			"M":           666,
		},

		favDrinks: []string{
			"Martins",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}

}
