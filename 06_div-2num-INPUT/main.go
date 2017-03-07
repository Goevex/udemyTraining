package main

import "fmt"

func main() {
	var num1 float32
	var num2 float32
	fmt.Println("First number:")
	fmt.Scan(&num1)
	fmt.Println("Second number:")
	fmt.Scan(&num2)
	fmt.Printf("%.0f %s %.0f %s %.32f \n", num1, " / ", num2, " = ", num1/num2)
}
