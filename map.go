package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Akbar",
		"address": "Sragen",
	}

	person["title"] = "Programmer"
	delete(person, "title")
	fmt.Println(person)
	fmt.Println(person["name"])
}
