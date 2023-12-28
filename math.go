package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Subtraction(1, 2))
	fmt.Println(Multiplication(1, 2))
}

func Sum(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Multiplication(a, b int) int {
	return a * b
}
