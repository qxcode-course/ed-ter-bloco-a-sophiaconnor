package main

import "fmt"

func main() {
	var Q int
	var D string
	fmt.Scanf("%d %s", &Q, &D)
	x := make([]int, Q)
	y := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Scanf("%d %d", &x[i], &y[i])
	}
	if D == "L" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Printf("%d %d\n", x[i]-1, y[i])
			} else {
				fmt.Printf("%d %d\n", x[i-1], y[i-1])
			}
		}
	}
	if D == "R" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Printf("%d %d\n", x[i]+1, y[i])
			} else {
				fmt.Printf(" %d %d\n", x[i-1], y[i-1])
			}
		}
	}
	if D == "U" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Printf("%d %d\n", x[i], y[i]-1)
			} else {
				fmt.Printf("%d %d\n", x[i-1], y[i-1])
			}
		}
	}
	if D == "D" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Printf("%d %d\n", x[i], y[i]+1)
			} else {
				fmt.Printf(" %d %d\n", x[i-1], y[i-1])
			}
		}
	}
}
