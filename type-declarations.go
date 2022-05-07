package main

import "fmt"

func main() {

	type NoKTP string

	var ktpEko NoKTP = "122121121212"

	fmt.Println(ktpEko)
	fmt.Println(NoKTP("dsasdasdasd"))

}
