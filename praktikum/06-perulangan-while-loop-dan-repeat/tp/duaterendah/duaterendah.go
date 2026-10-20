package main

import "fmt"

func main() {
	var n, nilai, nilai1 int
	var nama, nama1 string

	i := 2

	fmt.Scan(&n)

	fmt.Scan(&nama1, &nilai1)

	min1Nilai := nilai1
	min1Nama := nama1

	min2Nilai := 9999999
	min2Nama := ""

	for i <= n {
		fmt.Scan(&nama, &nilai)

		if nilai < min1Nilai {
			min2Nilai = min1Nilai
			min2Nama = min1Nama
			min1Nilai = nilai
			min1Nama = nama
		} else if nilai < min2Nilai && nilai != min1Nilai {
			min2Nilai = nilai
			min2Nama = nama
		}
		i++
	}
	fmt.Println(min1Nama)
	fmt.Println(min2Nama)
}
