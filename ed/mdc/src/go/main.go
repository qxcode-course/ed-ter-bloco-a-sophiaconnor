package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}
	c := a % b
	return mdc(b, c)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
