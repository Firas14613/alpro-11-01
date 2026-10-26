package main

import "fmt"

func main () {
	var n, m int

	fmt.Scan(&n, &m)

	for i := 1; i <= n && i <= m; i++ {
		if n % i == 0 && m % i == 0 {
			fmt.Print(i, " ")
		}
	}
}