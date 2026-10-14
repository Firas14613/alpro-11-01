package main

import "fmt"

func main() {
	var n, total int

	fmt.Scanln(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Printf("%d ", i)
		total += i
	}

	fmt.Println()
	fmt.Println(total)
}