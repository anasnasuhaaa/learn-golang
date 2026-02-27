package main

import "fmt"

func main() {
	var title string
	fmt.Scanln(&title)
	count := 0

	for i := 0; i < len(title); i++ {
		char := title[i]
		if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' {
			count++
		}
	}

	if count%2 != 0 {
		fmt.Println("Diterima")
	} else {
		fmt.Println("Tidak")
	}
}
