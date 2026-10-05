package main

import "fmt"

func main() {
	var x, low, high int
	var hasil bool

	fmt.Scan(&x, &low, &high)

	hasil = x >= low && x <= high
	fmt.Println(hasil)
}