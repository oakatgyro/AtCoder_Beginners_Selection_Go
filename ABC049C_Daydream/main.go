package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	judge := true

	for S != "" {
		if strings.HasPrefix(S, "dream") {
			_S := S[5:]
			if _S == "er" {
				break
			}
			if strings.HasPrefix(_S, "erd") || strings.HasPrefix(_S, "ere") {
				S = S[7:]
			} else {
				S = _S
			}
			continue
		} else if strings.HasPrefix(S, "erase") {
			_S := S[4:]
			if _S == "e" || _S == "er" {
				break
			}
			if strings.HasPrefix(_S, "erd") || strings.HasPrefix(_S, "ere") {
				S = S[6:]
			} else {
				S = S[5:]
			}
		} else {
			judge = false
			break
		}
	}

	if judge {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
