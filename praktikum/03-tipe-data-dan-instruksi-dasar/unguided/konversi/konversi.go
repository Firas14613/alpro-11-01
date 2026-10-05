package main

import "fmt"

func main () {
	var hari int

	fmt.Scan(&hari)
	tahun := hari / 360
	sisaTahun := hari % 360
	bulan := sisaTahun / 30
	sisaBulan := sisaTahun % 30
	minggu := sisaBulan / 7
	hari = sisaBulan % 7
	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(hari)
}