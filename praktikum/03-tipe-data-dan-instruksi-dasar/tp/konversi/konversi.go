package main

import "fmt"

func main() {
	var mil float64
	const kilo float64 = 1.6

	fmt.Scan(&mil)
	fmt.Printf("%.1f", mil * kilo)
}