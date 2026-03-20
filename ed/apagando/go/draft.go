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

	lista := make([]int,0, N-M) //id de cada pessoa que ficou na fila
	for i := 0; i < N; i++ {
		cont :=0
		for j := 0; j < M; j++ {
			if id[i] == sairam[j] {
				break
			}else{
				cont++
			}
			if cont == M {
				lista = append(lista, id[i])
			}
		}
	}
	/*for i:=0; i<N-M; i++ {
		if i == 0 {
			fmt.Printf("%d", lista[i])
		}else{
			fmt.Printf(" %d", lista[i])
		}
		if i==N-M-1 {
			fmt.Printf(" ")
		}
	}*/
	saida := ""
	for i:=0; i<N-M; i++ {
		saida = fmt.Sprint(saida, lista[i])
		saida = fmt.Sprint(saida, " ")
	}
	fmt.Println(saida[0:len(saida)-1])
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
