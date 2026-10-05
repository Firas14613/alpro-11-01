package main

import "fmt"

func main () {
	var r float64
	const phi float64 = 22.0 / 7.0

	fmt.Scan(&r)

	luasPermukaan := 4* phi * r * r
	fmt.Println(luasPermukaan)
}