package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nilai float64
	var name string
	var grade string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan Nama: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scanln(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80{
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", name, grade)
}
