package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m1 = map[string]Vertex{ // map literal
	"Golang": Vertex{
		10.101010, -20.202020,
	},
	"Java": Vertex{
		123.123123, 0.123321,
	},
}

var m2 = map[string]Vertex{ // map literal continue
	"Bell Labs": {0.123, -0.123},
	"Google":    {3.21, -9.876},
}
func main() {
	m = make(map[string]Vertex)
	m["Python"] = Vertex{
		0.123456, -9.876543,
	}
	fmt.Println(m["Python"])
	
	fmt.Println(" -------------  Map literal  -------------------")
	fmt.Println(m1)
	
	fmt.Println(" -------------  Map literal continue -------------------")
	fmt.Println(m2)
	
	fmt.Println(" -------------  Mutating Maps -------------------")
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 89
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}