package main

import "fmt"

func main() {
	var n, m, x, y, cangkir int

	fmt.Scan(&n, &m, &x, &y)

	for x <= n && y <= m {
		n -= x
		m -= y
		cangkir++
	}
	fmt.Println(cangkir)
}

// program cangkiKopi
// kamus
// 	n, m, x, y, cangkir : integer
// algoritma
// 	input(n, m, x, y)

// 	while x <= n and y <= m do
// 		n <- n - x
// 		m <- m - y
// 		cangkir <- cangkir + 1
// 	endwhile

// 	output(cangkir)
// endprogram