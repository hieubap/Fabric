package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int{ // example 
	a := -1
	b := 1
	return func() int{
		if a < b {
			a+=b
			return a
		} else {
		b+=a
		return b
		}
	}
}
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(2*i),
		)
	}
	
	fmt.Println(" ---------------  fibonacci ----------")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}