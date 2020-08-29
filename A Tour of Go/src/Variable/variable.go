package main

import (
"fmt"
)

var a,b bool

var d,e int = 1,2 // init

func main(){
	// variable
	var c bool
	fmt.Println(a,b,c)
	
	// variable with init
	var f,g,h = true,false,"init"
	fmt.Println(d,e,f,g,h)
}