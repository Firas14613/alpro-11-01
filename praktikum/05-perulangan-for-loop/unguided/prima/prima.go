package main

import "fmt"

func main () {
	var n, count int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Printf("%d ", i)
			count++
		}
	}
	fmt.Println()
	
	if count == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}