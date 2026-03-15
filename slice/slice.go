package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// Slice from index 2 until index before 5
	slice1 := days[2:5]
	fmt.Println(slice1)

	// Slice all from index 1
	slice2 := days[1:]
	fmt.Println(slice2)

	// Slice all until index 5
	slice3 := days[:5]
	fmt.Println(slice3)

	// Slice all
	slice4 := days[:]
	fmt.Println(slice4)
}
