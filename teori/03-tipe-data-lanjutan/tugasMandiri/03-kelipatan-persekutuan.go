package main

import "fmt"

func main() {
	var n, a, b int
	var hasil bool

	fmt.Scan(&n, &a, &b)
	hasil = n % a == 0 && n % b == 0

	fmt.Println(hasil)
}