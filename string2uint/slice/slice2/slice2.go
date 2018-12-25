package main

import "fmt"

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func main() {

	slice1 := make([]int, 0, 10)
	//slice1 := []int{}
	// for i := 0; i < len(slice1); i++ {
	// 	slice1[i] = i + 1
	// }
	for j := 0; j < 10; j++ {
		fmt.Println(slice1)
		fmt.Println("before", slice1, " ", &slice1, "(len, cap)", len(slice1), cap(slice1))
		fmt.Printf("addr %p\n", &slice1)
		slice1 = append(slice1, j+1)
		fmt.Println("append", slice1, " ", &slice1, "(len, cap)", len(slice1), cap(slice1))
		fmt.Printf("addr %p\n", &slice1)
	}

	var iBuffer [10]int
	slice := iBuffer[0:0]

	for i := 0; i < 20; i++ {
		fmt.Println(i, "--")
		slice = Extend(slice, i)
		fmt.Println(i, slice)
	}
}
