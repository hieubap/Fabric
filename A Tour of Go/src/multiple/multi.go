package main

import (
"fmt"
)
func swap(x, y string) (string, string) {// multiple
	return y, x
}

func divide(a,b int)(x,y int){ // named return value
	x = a/b;
	y = a%b;
	return
}

func main() {
	// multiple
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	// named return value
	c,d := divide(5,2);
	fmt.Println("5 chia 2 duoc %v du %v",c,d)
}