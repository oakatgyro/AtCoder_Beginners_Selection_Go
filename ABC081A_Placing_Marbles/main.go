package main

import (
	"fmt"
	"strings"
)

func main() {
	var block string
	sum := 0
	fmt.Scanf("%s", &block)

	slice := strings.Split(block, "")

	for _, num := range slice {
		if num == "1" {
			sum++
		}
	}
	fmt.Println(sum)
}
