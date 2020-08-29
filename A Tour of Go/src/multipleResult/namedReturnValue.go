package multipleResult

import (
"fmt"
)

func divide(a,b int)(x,y int){
	x = a/b;
	y = a%b;
	return
}

func main(){
	a,b := divide(5,2);
	fmt.Println("5 chia 2 duoc %g du %g",a,b)
}