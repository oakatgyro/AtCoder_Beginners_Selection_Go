package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		n, a, b     int
		total       int
		stringNum   string
		indexNumber int
		indexArray  []string
	)

	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 0; i <= n; i++ {
		stringNum = strconv.Itoa(i)
		indexArray = strings.Split(stringNum, "")
		indexSum := 0
		for _, index := range indexArray {
			indexNumber, _ = strconv.Atoi(index)
			indexSum += indexNumber
		}

		if a <= indexSum && indexSum <= b {
			total += i
		}

	}
	fmt.Println(total)
}
