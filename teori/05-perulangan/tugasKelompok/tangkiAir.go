package main

import "fmt"

func main() {
	var t, v, totalT int

	fmt.Scan(&t)

	for totalT <= t {
		fmt.Scan(&v)
		totalT += v

		fmt.Println(totalT >= t)
	}
}

// program tangkiAir
// kamus
// 	t, v, totalT : integer
// algoritma
// 	input(t)

// 	for totalT <= t do
// 		input(v)
// 		totalT <- totalT + v
// 		output(totalT >= t)
// 	endfor
// endprogram