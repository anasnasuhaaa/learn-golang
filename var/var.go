package main

import (
	"fmt"
)

func main() {
	var intNum int8 = 12
	const myName string = "ansupedia"

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	githubLink := "github.com/anasnasuhaaa"

	var integer8 int8 = 20

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Lagi Belajar Golang")
	// }

	var (
		phone1 = "Samsung"
		phone2 = "Xiaomi"
		phone3 = "infinix"
	) 
const (
		city1 = "Indramayu"
		city2 = "Bogor"
		city3 = "Jakarta"
	) 

	fmt.Println(phone1)
	fmt.Println(phone2)
	fmt.Println(phone3)

	fmt.Println(integer8)
	fmt.Println(intNum)
	fmt.Println(myName)
	fmt.Println(len(myName))
	fmt.Println(first, second, third)
	fmt.Println(githubLink)
}
