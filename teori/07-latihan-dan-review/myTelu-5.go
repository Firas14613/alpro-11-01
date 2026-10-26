package main

import "fmt"

func main() {
	var gold, platinum, silver, totalPoin int
	var goldPoin, platinumPoin, silverPoin int
	var rataGold, rataPlatinum, rataSilver float64

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
			goldPoin += poin
		} else if poin >= 100 {
			platinum++
			platinumPoin += poin
		} else {
			silver++
			silverPoin += poin
		}

		if totalPoin >= 500 {
			break
		}
	}

	if gold > 0 {
		rataGold = float64(goldPoin) / float64(gold)
	}
	if platinum > 0 {
		rataPlatinum = float64(platinumPoin) / float64(platinum)
	}
	if silver > 0 {
		rataSilver = float64(silverPoin) / float64(silver)
	}

	fmt.Printf("%.2f %.2f %.2f", rataGold, rataPlatinum, rataSilver)
}
