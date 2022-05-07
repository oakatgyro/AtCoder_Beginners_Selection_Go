package main

import (
	"fmt"
)

func main() {
	var N, Y int
	x, y, z := -1, -1, -1
	fmt.Scan(&N, &Y)

	for i := 0; i <= N; i++ {
		for j := 0; j <= N-i; j++ {
			k := N - i - j
			if 10000*i+5000*j+1000*k == Y {
				x, y, z = i, j, k
			}
		}
	}

	fmt.Println(x, y, z)
}
