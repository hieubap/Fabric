package main

import (
"fmt"
)

type Vertex struct{
	X,Y int
}

var ( // struct literal
	v1 = Vertex{1, 2}  
	v2 = Vertex{X: 1}  // Y:0
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // *Vertex
)

func main(){
	fmt.Println(Vertex{1,2})
	
	v := Vertex{4,5} // struct field
	v.X = 6
	fmt.Println(v)
	
	vertex := Vertex{0,1} // pointer to struct
	p := &vertex
	p.X = 1e9
	fmt.Println(vertex)
	
	fmt.Println(v1, p, v2, v3)
}