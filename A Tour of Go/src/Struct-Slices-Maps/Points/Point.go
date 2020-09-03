package main

import "fmt"

func main(){
	i,j := 42,1000
	
	p:= &i
	fmt.Println(*p)
	
	*p = 21
	fmt.Println(i)
	
	p=&j
	*p = *p/20
	fmt.Println(j)
}