package main

import "fmt"

func main() {
	var (
		n int
		genap bool
	)

	fmt.Scan(&n)

	genap = n % 2 == 0

	fmt.Println(genap)
}