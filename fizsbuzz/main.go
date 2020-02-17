package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 変数vはifのスコープ内でのみ有効
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(4, 4, 15),
		pow(2, 4, 20),
	)
}
