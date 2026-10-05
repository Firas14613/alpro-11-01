package main

import "fmt"

func main () {
	var (
		p, q int
		hasilOr, hasilAnd, hasilNot bool
	)

	fmt.Scan(&p, &q)
	hasilOr = p % 2 == 0 || q % 2 == 0
	hasilAnd = p % 2 != 0 && q % 2 != 0
	hasilNot = !(p == q)
	fmt.Println(hasilOr, " ", hasilAnd, " ", hasilNot)
}