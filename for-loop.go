package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan 2 ke ", i)
	}

	slice := []string{
		"Muhamad",
		"Akbar",
		"Abdul",
	}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for index, name := range slice {
		println("Index ", index, ", Range ", name)
	}

	for _, name := range slice {
		println("Name : ", name)
	}
}
