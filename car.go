package main

import (
	"fmt"
)

type car struct {
	engine string,
	colour string,
	sunroof bool,
	carnumber string

}

func main() {
	fmt.Print("ho gaya")

	alto := car {
		engine : "four stroke",
		colour : "white",
		sunroof : false,
		carnumber " JA 007 4R "
	}

	baleno := car {
		engine : "four stroke",
		colour : "grey",
		sunroof : false,
		carnumber " UP 007 4R ",
	}
	fmt.Print("ho gaya", alto)
	fmt.Print("ho gaya", baleno)

}
