package main

import (
"fmt"
"math"
)

func sqrt(input float64) string{ // if
	if input < 0 {
		return sqrt(-input)+"i"
	}
	return fmt.Sprint(math.Sqrt(input))
}

func pow(x, n, lim float64) float64 { // if with short statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func parity(x int) string{// if else
	if x % 2 == 0{
		return "even number"
	} else {
		return "odd number"
	}
}


func main(){
	fmt.Println(sqrt(9),sqrt(-4))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
	fmt.Println(parity(2))
	fmt.Println(parity(1))
	fmt.Println(parity(0))
	
}
