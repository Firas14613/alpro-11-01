package main

import "fmt"

func main () {
	var n, jam, i, jumlahJam int
	var rataRata float64
	jumlahJam = 0
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&jam)
		jumlahJam += jam
	}
	rataRata = float64(jumlahJam) / float64(n)
	fmt.Printf("%3f\n", rataRata)
}