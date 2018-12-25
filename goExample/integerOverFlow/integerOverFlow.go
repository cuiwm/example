package main

import "fmt"

// https://golang.org/ref/spec#Integer_overflow
func main() {
	var x int8
	x = -128
	fmt.Println(x / -1)
	//result: -128 Not 128 for Integer overflow
}
