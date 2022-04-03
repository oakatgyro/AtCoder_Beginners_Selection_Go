package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		n           int
		left, right int
	)

	fmt.Scanf("%d", &n)
	aArray := make([]int, n)

	for i := 0; i < len(aArray); i++ {
		fmt.Scan(&aArray[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(aArray)))

	for i, a := range aArray {
		if i%2 != 0 {
			left += a
		} else {
			right += a
		}
	}

	fmt.Println(int(math.Abs(float64(right - left))))

}
