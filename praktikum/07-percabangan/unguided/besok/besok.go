package main

import "fmt"

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	if bulan < 1 || bulan > 12 {
		fmt.Println("Bulan tidak valid")
		return
	}

	maxHari := 31
	switch bulan {
	case 4, 6, 9, 11:
		maxHari = 30
	case 2:
		kabisat := tahun % 4 == 0 || (tahun % 400 == 0 && tahun % 100 != 0)
		if kabisat {
			maxHari = 29
		} else {
			maxHari = 28
		}
	}

	if tanggal < 1 || tanggal > maxHari {
		fmt.Println("Tidak valid")
		return
	}

	fmt.Println("Valid")

	besokTanggal := tanggal + 1
	besokBulan := bulan
	besokTahun := tahun

	if besokTanggal > maxHari {
		besokTanggal = 1
		besokBulan++
		if besokBulan > 12 {
			besokBulan = 1
			besokTahun++
		}
	}

	fmt.Printf("%d/%d/%d", besokTanggal, besokBulan, besokTahun)
}