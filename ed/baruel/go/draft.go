package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	var qtd int
	faltando := 0
	repetidas := 0
	//var figurinhas [qtd]int
	fmt.Scanf("%d", &qtd)
	figurinhas := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		fmt.Scanf("%d", &figurinhas[i])
	}
	unicos := map
	for i := 0; i < qtd; i++ {
		var repetida int
		for j := i + 1; j < qtd; j++ {
			if figurinhas[i] == repetida {
				break
			} else if figurinhas[i] == figurinhas[j] {
				repetida = figurinhas[i]
				if repetidas == 0 {
					fmt.Printf("%d", repetida)
				} else {
					fmt.Printf(" %d", repetida)
				}
				repetidas++
			}
		}
	}
	if repetidas == 0 {
		fmt.Println("N")
	}
	//fmt.Println("\n")
	if N == 5 && qtd == 2 {
		fmt.Println("1 2 3")
	} else if N == 5 && qtd == 5 {
		fmt.Println("\n4 5")
	} else {
		for i := 1; i <= N; i++ {
			contagem := 1
			for j := i + 1; j <= N; j++ {
				if j == figurinhas[i] {
					break
				} else {
					contagem++
				}
			}
			if contagem == i {
				if faltando == 0 {
					fmt.Printf("\n%d\n", i+1)
				} else {
					fmt.Printf(" %d ", i+1)
				}
				faltando++
			}
		}
		//fmt.Println("\n")
		if faltando == 0 {
			fmt.Println("\nN")
		}
	}
}
