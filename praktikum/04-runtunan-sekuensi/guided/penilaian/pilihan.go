package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var pilihan int
	var nilai float64
	var nama string
	var grade string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("============== Menu ==============")
	fmt.Println("1. Sistem Penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukan Nama: ")
		scanner.Scan()
		nama = scanner.Text()
		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "E"
		}

		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
	} else if pilihan == 0 {
		fmt.Println("Keluar dari program.")
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}
