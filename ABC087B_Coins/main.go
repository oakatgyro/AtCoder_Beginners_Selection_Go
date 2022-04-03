package main

import "fmt"

func main() {
	var (
		a, b, c, x int
		count      int
	)

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	if x < 50 || 20000 < x {
		fmt.Println(count)
		return
	}

	for num500 := 0; num500 <= a; num500++ {
		for num100 := 0; num100 <= b; num100++ {
			for num50 := 0; num50 <= c; num50++ {
				if num500+num100+num50 < 1 {
					continue
				}
				if 500*num500+100*num100+50*num50 == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
