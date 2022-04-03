package main

import "fmt"

func main() {
	var (
		n     int
		count int
	)

	fmt.Scanf("%d", &n)
	aArray := make([]int, n)

	for i := 0; i < len(aArray); i++ {
		fmt.Scan(&aArray[i])
	}

	for true {
		for i := 0; i < len(aArray); i++ {
			if aArray[i]%2 != 0 {
				fmt.Println(count)
				return
			} else {
				aArray[i] = aArray[i] / 2
			}
			if i == len(aArray)-1 {
				count++
			}
		}
	}
}
