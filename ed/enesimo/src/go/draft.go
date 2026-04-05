package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}
	if x == div {
		return true
	}
	if x%div == 0 {
		return false
	}
	return eh_primo(x, div+1)
}
func enesimo_primo(n int, atual int) int {
	if eh_primo(atual, 2) {
		if n == 1 {
			return atual
		}
		return enesimo_primo(n-1, atual+1)
	}
	return enesimo_primo(n, atual+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(enesimo_primo(n, 2))
}
