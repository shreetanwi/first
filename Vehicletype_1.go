package main

import (
	"fmt"
)

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			door:  2,
			color: "White",
		},

		fourwheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			door:  4,
			color: "Black",
		},

		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.door)
	fmt.Println(s.door)
}
