package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addContinue(x ,y int) int{
	return x + y
}
func main() {
	fmt.Println(add(10, 22))
	fmt.Println(addContinue(20,30))
}

