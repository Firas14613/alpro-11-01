package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukkan nilai ujian: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			break
		}
		
		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else if nilai < 60 {
			fmt.Println("Anda tidak lulus ujian.")
		}
	}
	fmt.Println("Program Selesai.")
}