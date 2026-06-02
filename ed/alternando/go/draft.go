package main

import "fmt"

func show(vet []int, esc int) {
	fmt.Printf("[ ")
	for i := 0; i < len(vet); i++ {
		if i == 0 {
			if i+1 == esc {
				if vet[i] > 0 {
					fmt.Printf("%d> ", vet[i])
				} else {
					fmt.Printf("<%d ", vet[i])
				}
			}
		}
		if i != 0 && i+1 == esc {
			if vet[i] > 0 {
				fmt.Printf("%d> ", vet[i])
			} else {
				fmt.Printf("<%d ", vet[i])
			}
		}
		if i+1 != esc {
			if vet[i] != 0 {
				fmt.Printf("%d ", vet[i])
			}
		}
	}
	fmt.Printf("]\n")
}

func find_person(vet []int, size int, esc int) int { //achar o próximo vivo
	dead := 0
	if vet[esc] > 0 {
		pos := (esc + 1) % size
		for vet[pos] == dead {
			pos = (pos + 1) % size
		}
		return pos
	} else {
		pos := (esc - 1 + size) % size
		for vet[pos] == dead {
			pos = (pos - 1 + size) % size
		}
		return pos
	}
}

func main() {
	var size int
	var esc int
	var alternando int
	fmt.Scanf("%d", &size)       // quantidade de pessoas
	fmt.Scanf("%d", &esc)        //quem está com a espada
	fmt.Scanf("%d", &alternando) // 1 para começar com o positivo, -1 para começar com o negativo
	vet := make([]int, size)
	// Inicializar o vetor com os números 1 a size
	for i := 0; i < size; i++ {
		if alternando == 1 {
			if (i+1)%2 != 0 {
				vet[i] = i + 1
			} else {
				vet[i] = -(i + 1)
			}
		} else {
			if (i+1)%2 == 0 {
				vet[i] = i + 1
			} else {
				vet[i] = -(i + 1)
			}
		}
	}
	show(vet, esc)
	qtd := size
	for qtd > 1 {
		current := esc - 1
		gonna_die := find_person(vet, size, current)
		vet[gonna_die] = 0
		esc = find_person(vet, size, current) + 1
		show(vet, esc)
		qtd--
	}
}
