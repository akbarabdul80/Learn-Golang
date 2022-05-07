package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februsari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = months[4:7]

	fmt.Println(months)
	fmt.Println(slice)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Akbar"
	newSlice[1] = "Abdul"

	fmt.Println(newSlice)
}
