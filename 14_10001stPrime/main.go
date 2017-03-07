package main

import (
	"fmt"
	"math"
)

func main() {
	var primeCount int
	var thenThousandFirst int
	for i := 0; i <= math.MaxInt64; i++ {
		if isPrime(i) {
			primeCount++
		}
		if primeCount == 10001 {
			thenThousandFirst = i
			break
		}
	}

	fmt.Println("The 10001st prime is ", thenThousandFirst, ".")
}

func isPrime(intNumber int) bool {
	result := true
	if intNumber == 0 || intNumber == 1 {
		result = false
	} else {
		intCalcNumber := int(math.Sqrt(float64(intNumber)))
		for i := 2; i <= intCalcNumber; i++ {
			crux := float64(intNumber) / float64(i)
			if crux == math.Trunc(crux) {
				result = false
				break
			}
		}
	}
	return result
}
