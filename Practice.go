package main

import (
	"fmt"
)

var p int

func main() {
	y := "Tannu"
	fmt.Println(y)
	z := `HelloY "Tannu"`
	fmt.Println(z)

	a := 42
	fmt.Println(a)
	p := 10
	fmt.Println(p)
	// you()

	// Default value codes
	var l int
	fmt.Println(l)
	var gaya string
	fmt.Println(gaya)
	var t bool
	fmt.Println(t)
	// Print the type of variable
	// \t is tab
	// \n is for new line
	fmt.Printf("\t %T", t)
	fmt.Printf("\t %T", gaya)
	fmt.Printf("\t %T", l)
	// Creating your own type
	type bangalore int
	var tannu bangalore
	fmt.Printf("\n %T", tannu)
	p = int(tannu)

}

func you() {
	p := 100
	fmt.Println(p)
}
