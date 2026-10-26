package main

import "fmt"

func main() {
	var pengunjung, totalPengunjung int

	for i := 1; i <= 5; i++ {
		for {
			fmt.Printf("Hari %d : ", i)
			fmt.Scan(&pengunjung)

			if pengunjung >= 0 && pengunjung <= 200 {
				break
			}
		}
		totalPengunjung += pengunjung
	}
	fmt.Printf("Jumlah Pengunjuung : %d", totalPengunjung)

}
