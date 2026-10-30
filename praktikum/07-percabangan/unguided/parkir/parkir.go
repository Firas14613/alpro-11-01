package main

import "fmt"

func main() {
	var kendaraan string
	var waktu, tarif int

	fmt.Scanln(&kendaraan)
	fmt.Scanln(&waktu)

	tarifMobil := 5000
	tarifMotor := 3000

	if kendaraan == "mobil" {
		if waktu > 1 {
			tarif = tarifMobil + (3000 * (waktu - 1))
		} else {
			tarif = tarifMobil
		}
	} else if kendaraan == "motor" {
		if waktu > 1 {
			tarif = tarifMotor + (2000 * (waktu - 1))
		} else {
			tarif = tarifMotor
		}
	}

	if waktu > 24 {
		if kendaraan == "mobil" {
			tarif += 50000
		} else if kendaraan == "motor" {
			tarif += 30000
		}
	}

	fmt.Println(tarif)
}