package main

import "fmt"

func main() {
	var a, b, n, total int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&n)

	for i := 1; i <=n; i++ {
		total += a
		a += b
	}
	fmt.Println(total)
}