package main

import "fmt"

func show(vet []int, esc int) {
	fmt.Printf("[ ")
	for i := 0; i < len(vet); i++ {
		if i==0 {
			if i+1 == esc {
				fmt.Printf("%d> ", vet[i])
			}
		}
		if i!=0 && i+1 == esc {
			fmt.Printf("%d> ", vet[i])
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
	pos := (esc + 1) % size
	for vet[pos] == dead {
		pos = (pos + 1) % size
	}
	return pos
}

func main() {
	var size int
	var esc int
	fmt.Scanf("%d", &size) // quantidade de pessoas
	fmt.Scanf("%d", &esc)  //quem está com a espada
	vet := make([]int, size)
	// Inicializar o vetor com os números 1 a size
	for i := 0; i < size; i++ {
		vet[i] = i + 1
	}
	show(vet, esc)
	qtd := size
	for qtd > 1 {
		gonna_die := find_person(vet, size, esc-1) 
		vet[gonna_die] = 0
		esc = find_person(vet, size, gonna_die) +1
		show(vet, esc)
		qtd--
	}
}
