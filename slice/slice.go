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

	colors := [...]string{"Merah", "Jingga", "Kuning", "Hijau", "Biru", "Nilla", "Ungu"}
	for i := 1; i <= len(colors); i++ {
		fmt.Print(i, ". ")
		fmt.Println(colors[i-1])
	}
	x := 0
	y := 0

	fmt.Print("Masukan index mulai dan index sebelum akhir (dipisahkan spasi): ")
	fmt.Scan(&x, &y)
	fmt.Println(x, y)

	slice1 = colors[5:]
	fmt.Println(slice1)
	slice1[0] = "Nilla Baru"
	slice1[1] = "Ungu Baru"
	// fmt.Println(colors)
}
