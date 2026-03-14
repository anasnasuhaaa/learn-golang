package main

import (
	"fmt"
)

func main() {

	// Array Colors
	var myColors [4]string
	myColors[0] = "red"
	myColors[1] = "green"
	myColors[2] = "blue"
	myColors[3] = "black"

	// Array Days
	days := [7]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	// Create New mergeArr for saving value of myColors + days
	var mergeArr []string

	// Append myColors value to mergeArr
	for i := 0; i < len(myColors); i++ {
		mergeArr = append(mergeArr, myColors[i])
	}

	// Append days value to mergeArr
	for i := 0; i < len(days); i++ {
		mergeArr = append(mergeArr, days[i])
	}

	//Display mergeArr
	for j := 0; j < len(mergeArr); j++ {
		fmt.Println(j+1, mergeArr[j])
	}

	// User choice
	var userChoiceNumber int = 0
	fmt.Print("Pilih angka: ")
	fmt.Scan(&userChoiceNumber)

	// Display user choice
	for i := 0; i < len(mergeArr); i++ {
		if userChoiceNumber == i {
			fmt.Print("Your Coice is: ")
			fmt.Print(mergeArr[i-1])
		} else if userChoiceNumber > (len(mergeArr)-1) || userChoiceNumber < 1 {
			fmt.Print("Pilihan tidak tersedia")
			break
		}

	}

}
