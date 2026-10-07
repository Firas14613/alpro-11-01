package main

import "fmt"

func main() {
	var nilaiMtk, nilaiBhsIng int

	fmt.Print("Masukkan nilai Matematika: ")
	fmt.Scan(&nilaiMtk)
	fmt.Print("Masukkan nilai Bahasa Inggris: ")
	fmt.Scan(&nilaiBhsIng)
	rataRata := (nilaiMtk + nilaiBhsIng) / 2
	fmt.Println(rataRata)

	switch {
		case rataRata >= 90:
			fmt.Println("A")
		case rataRata >= 80:
			fmt.Println("B")
		case rataRata >= 70:
			fmt.Println("C")
		default:
			fmt.Println("D")
	}
}
