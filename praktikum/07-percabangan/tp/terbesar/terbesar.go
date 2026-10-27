package main

import "fmt"

func main() {
	var a, b, c, terbesar int

	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		terbesar = a
	} else if b >= a && b >= c {
		terbesar = b
	} else {
		terbesar = c
	}

	fmt.Println(terbesar)
}