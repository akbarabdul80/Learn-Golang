package main

func main() {
	var nilai1 = 90
	var nilai2 = 10

	if nilai1 > nilai2 {
		println(nilai1)
	} else if nilai1 == nilai2 {
		println("nilai1 sama dengan dari nilai2")
	} else {
		println("nilai1 lebih kecil dari nilai2")
	}

	if kkm := 60; nilai1 < kkm {
		println("Nilai tidak mencapai kkm")
	} else {
		println("Nilai cukup")
	}

}
