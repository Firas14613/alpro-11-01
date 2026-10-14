package main

import "fmt"

func main() {
	var transaksi, jumlahTiket int
	var hari, jenisFilm string
	var diskon int

	totalPendapatan := 0

	fmt.Scanln(&transaksi)
	for i := 1; i <= transaksi; i++ {
		fmt.Printf("Transaksi ke-%d\n : ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0;

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenisFilm == "Regular" {
				hargaSatuan = 40000
			} else if jenisFilm == "3D" {
				hargaSatuan = 60000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 80000
			}
		} else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenisFilm == "Regular" {
				hargaSatuan = 50000
			} else if jenisFilm == "3D" {
				hargaSatuan = 70000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 90000
			}
		}

		totalHarga := hargaSatuan * jumlahTiket

		fmt.Printf("Harga Sebelum Diskon: %d\n", totalHarga)

		if jumlahTiket >= 5 {
			diskon = (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}
		fmt.Printf("Diskon sebesar: %d\n", diskon)
		fmt.Printf("Total Harga : %d\n", totalHarga)


		totalPendapatan += totalHarga
	}

	fmt.Println("Total Pendapatan : ", totalPendapatan)

}