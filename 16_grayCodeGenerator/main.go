package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	var count int
	var countSlice []int

	for true {
		fmt.Println("Welche Bitlänge soll der Grey Code haben?")
		fmt.Scan(&input)
		if c, err := strconv.Atoi(input); err == nil {
			count = c
			if count > 0 { //strconv.ParseInt() also works
				break
			} else {
				fmt.Println("Gib eine Zahl größer 0 ein!")
			}
		}
	}

	countSlice = make([]int, int(math.Pow(float64(2), float64(count))))
	for i := range countSlice {
		countSlice[i] = i
	}

	fmt.Println(countSlice)
}
