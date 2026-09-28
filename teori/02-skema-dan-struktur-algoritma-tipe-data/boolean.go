package main

import "fmt"

func main() {
	// fmt.Println("benar: ", true)
	// fmt.Println("salah: ", false)

	nilaiAkhir := 80
	nilaiAbsen := 90

	var lulusNilaiAkhir = nilaiAkhir >= 75
	var lulusNilaiAbsen = nilaiAbsen >= 75

	var lulus = lulusNilaiAkhir && lulusNilaiAbsen
	fmt.Println(lulus)

}