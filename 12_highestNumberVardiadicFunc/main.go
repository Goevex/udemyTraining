package main

import "fmt"

func main() {
	coupleOfNumbers := []int{1, 2, 5, 6, 1, 9, 7}
	hNumber := max(coupleOfNumbers...)
	fmt.Println("The highest Number is ", hNumber)
}

func max(slice ...int) int {
	high := slice[0]
	for _, h := range slice {
		if h > high {
			high = h
		}
	}
	return high
}
