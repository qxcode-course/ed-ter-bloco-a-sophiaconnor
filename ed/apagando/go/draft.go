package main

import "fmt"

func main() {
	var N int //qtd inicial de pessoas na fila
	fmt.Scanf("%d", &N)
	id := make([]int, N) //id de cada pessoa na fila
	var M int            //qtd de pessoas que sairam da fila
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &id[i])
	}
	fmt.Scan(&M)
	sairam := make([]int, M) //id de cada pessoa que saiu da fila
	for i := 0; i < M; i++ {
		fmt.Scan(&sairam[i])
	}

	//fmt.Println(N, M, id, sairam)

	lista := make([]int, 0) //id de cada pessoa que ficou na fila
	sairamCount := make(map[int]int)
	for j := 0; j < M; j++ {
		sairamCount[sairam[j]]++
	}

	for i := 0; i < N; i++ {
		if sairamCount[id[i]] > 0 {
			sairamCount[id[i]]--
			continue
		}
		lista = append(lista, id[i])
	}

	saida := ""
	for i := 0; i < len(lista); i++ {
		saida = fmt.Sprint(saida, lista[i])
		if i < len(lista)-1 {
			saida = fmt.Sprint(saida, " ")
		}
	}
	fmt.Print(saida)
	fmt.Print(" \n")
	//fmt.Printf(" ")
	/*for i := 0; i < N; i++ {
		cont := 0
		for j := 0; j < M; j++ {
			verificar := 0
			if id[i] == sairam[j] {
				break
			} else {
				cont++
			}
			if cont == M {
				if verificar == 0 {
					fmt.Printf("%d", id[i])
				} else {
					fmt.Printf(" %d", id[i])
				}
				verificar++
			}
		}
	}*/
}
