package main

func main() {

	slice := []string{
		"Muhamad",
		"Akbar",
		"Abdul",
	}

	for _, name := range slice {
		if name == "Akbar" {
			continue
		}
		println("Name : ", name)
	}

	for _, name := range slice {
		if name == "Akbar" {
			break
		}
		println("Name : ", name)
	}
}
