package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		N              int
		t1, t2         int
		x1, x2, y1, y2 float64
	)
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&t2, &x2, &y2)
		dt := t2 - t1
		dist := int(math.Abs(x2-x1) + math.Abs(y2-y1))

		if dt < dist || dist%2 != dt%2 {
			fmt.Println("No")
			return
		}
		t1, x1, y1 = t2, x2, y2
	}
	fmt.Println("Yes")
}
