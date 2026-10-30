package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var layanan string
	var tarif, zona int

	fmt.Print("Masukkan Berat (kg): ")
	fmt.Scan(&berat)
	fmt.Print("Masukkan Zona (1/2/3): ")
	fmt.Scan(&zona)
	fmt.Print("Masukkan Layanan (Regular/Express): ")
	fmt.Scan(&layanan)

	hasilBulat := math.Ceil(berat)
	berat = hasilBulat

	if berat >= 30 {
		fmt.Println("-")
		return
	}

	if layanan != "Regular" && layanan != "Express" {
		fmt.Println("Layanan tidak valid")
		return
	}

	switch zona {
	case 1:
		tarif = 10000
	case 2:
		tarif = 15000
	case 3:
		tarif = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}

	tarif *= int(berat)

	if layanan == "Express" {
			tarif += tarif / 2
	}

	fmt.Printf("%.0f\n%d", berat, tarif)
}
