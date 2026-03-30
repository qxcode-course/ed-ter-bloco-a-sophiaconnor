package main

import "fmt"

func main() {
	var T, R int
	fmt.Scan(&T, &R)

	vetor := make([]int, T)
	for i := 0; i < T; i++ {
		fmt.Scan(&vetor[i])
	}
    
	if T == 0 {
		fmt.Println("[ ]")
		return
	}
    
	r := R % T
	if r < 0 {
		r += T
	}
    
    
	vetor_rotacionado := make([]int, 0, T)
	for i := T - r; i < T; i++ {
		vetor_rotacionado = append(vetor_rotacionado, vetor[i])
	}
	for i := 0; i < T-r; i++ {
		vetor_rotacionado = append(vetor_rotacionado, vetor[i])
	}

	fmt.Print("[ ")
	for i, v := range vetor_rotacionado {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println(" ]")
   // fmt.Println(vetor_rotacionado)
}
