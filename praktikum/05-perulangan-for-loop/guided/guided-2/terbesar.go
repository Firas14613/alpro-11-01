package main

import "fmt"

func main() {
	var n, max, num int

	fmt.Scanln(&n)

	// for i := 1; i <= n; i++ {
	// 	fmt.Scan(&num)
	// 	if i == 1 {
	// 		max = num
	// 	} else {
	// 		if num > max {
	// 			max = num
	// 		}
	// 	}
	// }
	// fmt.Println(max)

	fmt.Scan(&num)
	max = num
	for i := 2; i <= n; i++ {
		fmt.Scan(&num)
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
}