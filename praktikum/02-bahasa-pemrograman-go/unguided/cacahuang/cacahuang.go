package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	lembar1 := a / 10000
	sisa1 := a % 10000

	lembar2 := sisa1 / 5000
	sisa2 := sisa1 % 5000

	lembar3 := sisa2 / 1000

	fmt.Println(lembar1, " ", lembar2, " ", lembar3)

	// pertama menghitung jumlah lembar uang 10000, lalu sisa dari pembagian tersebut disimpan di sisa1
	// kedua menghitung jumlah lembar uang 5000, lalu sisa dari pembagian tersebut disimpan di sisa2
	// ketiga menghitung jumlah lembar uang 1000, lalu sisa dari pembagian tersebut disimpan di sisa3
}
