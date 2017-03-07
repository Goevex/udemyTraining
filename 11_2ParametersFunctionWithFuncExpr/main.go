package main

import "fmt"

func main() {
	half := func(number int) (float32, bool) {
		return float32(number) / 2, number%2 == 0
	}

	h, even := half(75)
	fmt.Println(h, even)
}
