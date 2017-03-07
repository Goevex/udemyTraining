package main

import "fmt"

func main() {
	foo(1, 2, 9)
	foo(0, 7)
	aSlice := []int{987, 446, 87, 2}
	foo(aSlice...)
}

func foo(slice ...int) {
	fmt.Println(slice)
}
