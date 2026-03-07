package main

import "fmt"

func main() {
	type Student struct {
		name  string
		grade int8
	}

	var s1 Student
	s1.grade = 14
	s1.name = "Anas"
	fmt.Println("Name: ", s1.name)
	fmt.Println("Grade: ", s1.grade)
}
