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
			if count > 0 {
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

	fmt.Println("Ausgangs-Array:")
	var binaryCountSlice []string
	binaryCountSlice = append(binaryCountSlice, intSliceToBinaryStringSclice(countSlice)...)
	fmt.Println(binaryCountSlice)
	fmt.Println("")
	fmt.Println("Gray-Code:")
	var graySlice []string
	graySlice = append(graySlice, binaryStringArrayToGrayArray(binaryCountSlice)...)
	fmt.Println(graySlice)

}

func intSliceToBinaryStringSclice(intSlice []int) []string {
	stringArray := make([]string, len(intSlice))
	for i := range stringArray {
		stringArray[i] = strconv.FormatInt(int64(intSlice[i]), 2)
	}
	for i := range stringArray {
		format := "%0" + strconv.Itoa(len(stringArray[len(stringArray)-1])) + "s"
		stringArray[i] = fmt.Sprintf(format, stringArray[i])
	}
	return stringArray
}

func binaryStringArrayToGrayArray(stringSlice []string) []string {
	returnArray := make([]string, len(stringSlice))
	for i := range stringSlice {
		returnArray[i] = string(stringSlice[i][0])
		for j := 1; j < len(stringSlice[i]); j++ {
			var firstBool, secondBool, resultBool bool
			var resultPseudoChar string
			firstBool = (stringSlice[i][j-1] == '1')
			secondBool = (stringSlice[i][j] == '1')
			resultBool = (!firstBool && secondBool) || (firstBool && !secondBool) //XOR
			if resultBool {
				resultPseudoChar = "1"
			} else {
				resultPseudoChar = "0"
			}
			returnArray[i] = returnArray[i] + resultPseudoChar
		}
	}
	return returnArray
}
