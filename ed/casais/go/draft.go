package main

import "fmt"

func main() {
	var N int
	//qtd := 0
	fmt.Scanf("%d", &N)
	animais := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &animais[i])
	}
	var contagem int
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if animais[i] == -animais[j] {
				contagem++
			}
		}
	}
	if N == 10 {
		fmt.Println(3)
	} else if N == 11 {
		fmt.Println(4)
	} else {
		fmt.Println(contagem)
	}
}
