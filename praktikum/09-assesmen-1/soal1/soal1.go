package main

import "fmt"

func main() {
	var a, i int
	var terakhir, total, tertinggi, terendah int

	fmt.Scan(&a)

	terakhir = a
	tertinggi = a
	terendah = a
	i = 1
	total = a

	for {
		fmt.Scan(&a)

		if a == terakhir {
			break
		}

		total += a
		i++

		if a < terendah {
			terendah = a
		}

		if a > tertinggi {
			tertinggi = a
		}

		terakhir = a
	}

	rataRata := float64(total) / float64(i)
	fmt.Printf("Tertinggi: %d\nTerendah: %d\nRata-rata: %.2f\n", tertinggi, terendah, rataRata)
}