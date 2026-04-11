package main

import (
	"fmt"
)

func main() {

	var n int = 0
	fmt.Scanln(&n)
	var number []int

	for i := 0; i < n; i++ {
		var appendNumber int
		fmt.Scan(&appendNumber)
		number = append(number, appendNumber)
	}

	for i := 0; i < (len(number) - 1); i++ {
		var counter int
		for j := 0; j < (len(number) - i - 1); j++ {
			if number[j] < number[j+1] {
				temp := number[j]
				number[j] = number[j+1]
				number[j+1] = temp
				counter++
			}
		}
	}

	result := 0
	for i := 0; i < 3; i++ {
		if len(number) > 0 {
			result += number[i]
		}
	}
	fmt.Println(result)

}
