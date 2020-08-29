package controlStatement

import "fmt"

// trong Go for continue chính là while
func forEver(){
	for{
		
	}
}
func main() {
	sum := 0
	for i:=0;i<=10;i++{// for
		sum += i
	}
	fmt.Printf(sum)
	
	
	sum1 := 1
	for ; sum < 1000; {// for continue
		sum += sum
	}
	fmt.Println(sum1)
	
	
	sum2 := 1
	for sum1 < 1000 {// while
		sum1 += sum1
	}
	fmt.Println(sum2)
}