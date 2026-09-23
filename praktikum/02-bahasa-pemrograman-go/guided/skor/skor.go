package main

import "fmt"

func main() {
	var (
		nama string
		skorMtk int
		skorBhsInggris int
		totalSkor int
		rataRata int
	)

	// nama = "Hong Gil_Dong"
	// skorMtk = 96
	// skorBhsInggris = 82
	// totalSkor = skorMtk + skorBhsInggris
	// rataRata = (totalSkor) / 2
	// fmt.Println(nama)
	// fmt.Println(totalSkor)
	// fmt.Println(rataRata)

	// nama = "Rina Amalia"
	// skorMtk = 88
	// skorBhsInggris = 75
	// totalSkor = skorMtk + skorBhsInggris
	// rataRata = (totalSkor) / 2
	// fmt.Println(nama)
	// fmt.Println(totalSkor)
	// fmt.Println(rataRata)

	fmt.Scan(&nama)
	fmt.Scan(&skorMtk)
	fmt.Scan(&skorBhsInggris)

	totalSkor = skorMtk + skorBhsInggris
	rataRata = totalSkor / 2

	fmt.Println(nama)
	fmt.Println(totalSkor)
	fmt.Println(rataRata)
}
