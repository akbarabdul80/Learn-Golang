package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	a -= b
	a++

	var c = a + b
	fmt.Println(c)

}
