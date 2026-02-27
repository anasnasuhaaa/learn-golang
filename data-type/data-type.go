package main

import "fmt"

func main() {
	// String
	var string1 string = "Kopi Susu"

	// Numeric
	// int8: -128 to 127
	// int16: -32768 to 32767
	// etc
	// const phi := 3.234324

	const (
		dayName string = "senin"
		dayNum  uint8  = 1
	)

	fmt.Println(dayName, dayNum)
	const myAge int8 = 18
	totalDays := 365

	fmt.Println(totalDays)
	fmt.Println(string1)
	fmt.Println("")
}
