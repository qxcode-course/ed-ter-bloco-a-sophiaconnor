package main

import "fmt"

func main() {
	var N int
	var E int
	vivos := make([]int, N)
	for i := 0; i < N; i++ {
		vivos[i] = i + 1
	}
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &E)
	for i := 0; i < N; i++ {
		if i == E {
			vivos[i+1] = 0
		}
		if i == E+2 {
			if vivos[i] == 0 {
				E++
			} else {
				vivos[i+1] = 0
			}
		}
		for j := 0; j < N; j++ {
			if vivos[i] != 0 {
				fmt.Printf("%d ", vivos[j])
			}
		}
	}
}
