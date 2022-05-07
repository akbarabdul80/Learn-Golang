package main

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
