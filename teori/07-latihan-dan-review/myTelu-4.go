package main

import "fmt"

func main() {
	var gold, platinum, silver, totalPoin int

	for {
		var poin int

		for {
			fmt.Print("Masukan Poin: ")
			fmt.Scan(&poin)

			if poin >= 50 {
				break
			}
		}

		if totalPoin >= 500 {
			break
		}

		totalPoin += poin

		if poin > 200 {
			gold++
		} else if poin >= 100 {
			platinum++
		} else {
			silver++
		}

		if totalPoin >= 500 {
			break
		}
	}
	fmt.Printf("%d %d %d", gold, platinum, silver)
}
