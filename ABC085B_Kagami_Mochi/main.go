package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	mochiNum := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&mochiNum[i])
	}

	sort.Ints(mochiNum)

	keys := make(map[int]bool)
	var list []int
	for _, mochi := range mochiNum {
		if _, value := keys[mochi]; !value {
			keys[mochi] = true
			list = append(list, mochi)
		}
	}
	fmt.Println(len(list))

}
