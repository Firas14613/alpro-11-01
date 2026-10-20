package main

import "fmt"

func main() {
	var n, total int

	fmt.Scanln(&n)

	if n < 1 {
		fmt.Println("Nilai tidak valid")
		return
	}
	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Println(total)
}
