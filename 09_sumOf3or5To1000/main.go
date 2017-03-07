package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	fmt.Println("The sum of all numbers, which can be devided by 3 and 5 between 0 and 1000 is ", sum, ".")
}
