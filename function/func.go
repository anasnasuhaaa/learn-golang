package main

import (
	"fmt"
)

func main() {
	fmt.Print("")
	fmt.Println(operator(2, 2, "^"))
}

func operator(num1 int, num2 int, opp string) int {
	result := 0
	switch opp {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "^":
		result = intPow(num1, num2)
	}
	return result
}

func intPow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
