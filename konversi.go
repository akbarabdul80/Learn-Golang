package main

import "fmt"

func main() {
	var nilai32 int32 = 32675
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var name = "Akbar"
	var e = name[0]
	var eString = string(e)
	fmt.Println(eString)
}
