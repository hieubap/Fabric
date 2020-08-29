package main

import (
"fmt"
"math"
)

func main(){
	var i = 64;
	var f float64 = math.Sqrt(float64(i))
	var u uint = uint(f)
	
	fmt.Println(i,f,u)
}