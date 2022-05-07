package main

func main() {
	firstName, lastName, middleName := getFullname1()
	println(firstName, lastName, middleName)
}

func getFullname1() (firstName, middleName, lastName string) {
	firstName = "Muhamad"
	middleName = "Akbar"
	lastName = "Abdul"
	return
}
