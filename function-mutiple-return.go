package main

func main() {
	println(getHello("Akbar"))
	firstName, lastName := getFullname()
	println(firstName, lastName)
}

func getHello(name string) (string, string) {
	return "Hello " + name, ", Selamat Datang"
}

func getFullname() (string, string) {
	return "Muhamad", "Akbar"
}
