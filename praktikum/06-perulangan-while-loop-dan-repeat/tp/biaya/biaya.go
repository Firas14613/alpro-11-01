package main

import "fmt"

func main () {
	var n int
	var ukt float64 = 10000000

	fmt.Scanln(&n)

	for i := 2; i <= n; i++ {
		ukt *= 1.06
	}

	fmt.Printf("%0.f", ukt)
}