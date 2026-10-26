package main

import "fmt"

func main() {
	var poin, n, gold, platinum, silver int

	fmt.Print("Masukan Jumlah Mahasiswa: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for {
			fmt.Printf("Masukan Poin Mahasiswa ke- %d: ", i)
			fmt.Scan(&poin)
			
			if poin >= 50 {
				break
			}
		}
		if poin > 200 {
			gold++
		} else if poin >= 100 {
			platinum++
		} else {
			silver++
		}
	}

	fmt.Printf("%d %d %d", gold, platinum, silver)
}
