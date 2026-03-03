package main

import (
	"fmt"
)

func main() {
	var myColors [4]string
	myColors[0] = "red"
	myColors[1] = "green"
	myColors[2] = "blue"
	myColors[3] = "black"

	for i := 0; i < len(myColors); i++ {
		fmt.Println(i+1, myColors[i])
	}

	var userChoiceNumber int = 0
	fmt.Print("Pilih angka: ")
	fmt.Scan(&userChoiceNumber)

	for i := 0; i < len(myColors); i++ {
		if userChoiceNumber == i {
			fmt.Print("Your Coice is: ")
			fmt.Print(myColors[i-1])
		}
	}

}
