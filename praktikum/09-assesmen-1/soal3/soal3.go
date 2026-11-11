package main

import "fmt"

func main() {
	var n, daun, tersempit int
	fmt.Scan(&n)
	fmt.Scan(&daun)
	tersempit = daun

	for i := 2; i <= n; i++ {
		fmt.Scan(&daun)
		if daun < tersempit {
			tersempit = daun
		}
	}
	fmt.Println(tersempit)
}