package main

import "fmt"

func main() {
	names := [...]string{
		"Fadjrir",
		"Herlambang",
		"Purnomo",
		"Hidayat",
		"Saputra",
		"Setiawan",
	}

	slice1 := names[2:4]
	slice2 := names[:3]
	slice3 := names[3:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// Append Slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Senin Baru"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fadjrir"
	newSlice[1] = "Herlambang"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Purnomo")
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	iniSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
