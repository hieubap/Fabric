package main

import(
	"fmt"
	"strings"
)
func main() {
	fmt.Println("--------------------  Slices  ----------------------------------------") //slices
	number := [6]int{1,2,3,4,5,6}

	var s []int = number[1:4]
	fmt.Println(s)
	
	fmt.Println("----------------  Slices are like reference to arrays  -----------------------") //slices are like reference to arrays
	names := [4]string{
		"A",
		"B",
		"C",
		"D",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	
	fmt.Println("-----------  Slice literals  ------------") // Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	str := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(str)
	
	fmt.Println(" -----------------  Slice default  --------------------- ") // slice default
	defaul := []int{0,1,2,3,4,5,6,7,8,9}

	defaul = defaul[1:7]
	fmt.Println(defaul)

	defaul = defaul[:8]
	fmt.Println(defaul)

	defaul = defaul[3:]
	fmt.Println(defaul)
	
	fmt.Println(" ------------------  Slice length and capacity ---------------------- ")
	s2 := []int{0,1,2,3,4,5,6,7,8,9}
	printSlice(s2)

	s2 = s2[:0]
	printSlice(s)

	s2 = s2[:5]
	printSlice(s)

	s2 = s2[2:]
	printSlice(s)
	
	fmt.Println(" -----------------  Nil slice ---------------------------------------")
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil!")
	}
	
	fmt.Println(" -----------------  Creating a slice with make ---------------------------------------")
	e := make([]int, 5)
	printSlices("a", e)

	f := make([]int, 0, 5)
	printSlices("b", f)

	g := f[:2]
	printSlices("c", g)

	h := g[2:5]
	printSlices("d", h)
	
	fmt.Println(" ------------------  Slices of slices  -----------------------")
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "X", "_"},
		[]string{"O", "X", "O"},
		[]string{"_", "X", "_"},
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	
	fmt.Println(" ----------------  Appending to a slice----------------------")
	var s4 []int
	printSlice(s4)

	// append works on nil slices.
	s4 = append(s4, 0)
	printSlice(s4)

	// The slice grows as needed.
	s4 = append(s4, 1)
	printSlice(s4)

	// We can add more than one element at a time.
	s4 = append(s4, 2, 3, 4)
	printSlice(s4)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
