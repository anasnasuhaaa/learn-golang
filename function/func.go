package main

import (
	"fmt"
)

func main() {
	fmt.Print("")
	var num1 int
	var num2 int
	var opp string

	fmt.Scanf("%d %d %s", &num1, &num2, &opp)
	fmt.Println(operator(num1, num2, opp))
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
