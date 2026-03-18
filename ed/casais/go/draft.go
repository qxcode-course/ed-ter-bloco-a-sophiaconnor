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
	verifica := make([]bool, N)
	for i := 0; i < N; i++ {
		verifica[i] = false
	}
	var contagem int
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if animais[i] == -animais[j] && verifica[i] == false && verifica[j] == false {
				contagem++
				verifica[i] = true
				verifica[j] = true
			}
		}
	}
	fmt.Println(contagem)
}
