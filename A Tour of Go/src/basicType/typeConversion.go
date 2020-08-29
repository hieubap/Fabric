package basicType

import (
"fmt"
"math"
)

func main(){
	var i = 64;
	var f float64 = math.Sqrt(i)
	var u uint = uint(f)
	
	fmt.Println(i,f,u)
}