package main

func main() {
	var name = "Akbar"

	switch name {
	case "Akbar":
		println("Nama Akbar")
	case "Joko":
		println("Nama Joko")
	default:
		println("Nama tidak dikenal")
	}

	switch {
	case name == "Akbar":
		println("Nama Akbar")
	case name == "Joko":
		println("Nama Joko")
	default:
		println("Nama tidak dikenal")
	}

}
