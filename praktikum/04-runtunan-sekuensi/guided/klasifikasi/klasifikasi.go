package main

import "fmt"

func main() {
	var umur, gaji int
	var keterangan string

	fmt.Scanln(&umur)
	fmt.Scanln(&gaji)

	switch {
	case umur < 18:
		keterangan = "Masih sekolah"
	case umur >= 18 && umur <= 25 && gaji >= 50:
		keterangan = "Muda Sukses"
	case umur >= 18 && umur <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case umur > 25 && umur <= 40 && gaji >= 100:
		keterangan = "Pekerja Mapan"
	case umur > 25 && umur <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karir"
	case umur > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case umur > 40 && gaji < 150:
		keterangan = "Perlu evaluasi karir"
	}

	fmt.Println(keterangan)
}
