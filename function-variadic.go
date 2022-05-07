package main

func main() {
	total := sumAll(1, 2, 3, 5)
	println(total)

	slice := []int{10, 100, 12}
	total = sumAll(slice...)
	print(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
