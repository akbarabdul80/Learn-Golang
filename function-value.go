package main

func main() {
	goodbye := getGoodBaye
	println(goodbye("akbar"))
}

func getGoodBaye(name string) string {
	return "Good Bye " + name
}
