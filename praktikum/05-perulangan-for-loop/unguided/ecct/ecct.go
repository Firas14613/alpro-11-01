package main

import "fmt"

func main () {
	var m, p, v, ga,fs, ci, tc, lulus int
	var rataRata float64

	fmt.Scan(&m)

	for i := 1; i <= m; i++ {
		fmt.Printf("Mahasiswa ke-%d: ", i)
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		if (p < 1 || p > 4) || (v < 1 || v > 4) || (ga < 1 || ga > 4) || (fs < 1 || fs > 4) || (ci < 1 || ci > 4) || (tc < 1 || tc > 4) {
			fmt.Println("Input tidak valid! Harus antara 1 sampai 4.")
			return
		}

		rataRata = float64(p + v + ga + fs + ci + tc) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Ya")
			lulus++
		} else {
			fmt.Println("Tidak")
		}
	}

	fmt.Println(lulus)
}