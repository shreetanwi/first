package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := "Tanwi Shree",
}

fmt.Println(p1)
ChangeMe(&p1)
fmt.Println(p1)
ChangeMe(p1)
}

func ChangeMe(p *person){
	p.name = "James Bond"
	(*p).name = "Monneyp"
}