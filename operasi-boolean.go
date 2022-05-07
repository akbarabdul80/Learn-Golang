package main

func main() {
	var nilaiUTS = 90
	var nilaiAkhir = 80

	var lulusUTS bool = nilaiUTS >= 80
	var lulusAkhir bool = nilaiAkhir >= 80

	var lulus bool = nilaiUTS >= 80 && nilaiAkhir >= 90

	println(lulusUTS)
	println(lulusAkhir)
	println(lulus)
}
