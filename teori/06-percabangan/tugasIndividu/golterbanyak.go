package main

import "fmt"

func main() {
	var a, b, c ,d int

	fmt.Scan(&a, &b, &c, &d)

	terbanyak := a
	tersedikit := a

	if b > terbanyak {
		terbanyak = b
	}
	if c > terbanyak {
		terbanyak = c
	}
	if d > terbanyak {
		terbanyak = d
	}

	if b < tersedikit {
		tersedikit = b
	}
	if c < tersedikit {
		tersedikit = c
	}
	if d < tersedikit {
		tersedikit = d
	}

	fmt.Printf("gol terbanyak = %d, gol tersedikit = %d", terbanyak, tersedikit)
}

// program golTerbanyak
// kamus
//     a, b, c, d, terbanyak, tersedikit : integer
// algoritma
//     input(a, b, c, d)

//     terbanyak <- a
//     tersedikit <- a

//     if b > terbanyak then
//         terbanyak <- b
//     endif
//     if c > terbanyak then
//         terbanyak <- c
//     endif
//     if d > terbanyak then
//         terbanyak <- d
//     endif

//     if b < tersedikit then
//         tersedikit <- b
//     endif
//     if c < tersedikit then
//         tersedikit <- c
//     endif
//     if d < tersedikit then
//         tersedikit <- d
//     endif

//     print("gol terbanyak = ", terbanyak, ", gol tersedikit = ", tersedikit)
// endprogram