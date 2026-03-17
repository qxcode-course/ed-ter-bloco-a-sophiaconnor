package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	contagem := 0
	menor := 1000
	vencedor := N-1
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scanf("%d %d", &A, &B)
		if A >= 10 && B >= 10 {
			dif := A - B
			if dif < 0 {
				dif = -dif
			}
			if dif < menor {
				menor = dif
				vencedor = i
			} else if dif == menor && i < vencedor {
				vencedor = i
			}
		}else{
			contagem++
		}
	}
	if contagem == N {
		fmt.Println("sem ganhador")
		return
	}else{
		fmt.Println(vencedor)
	}
}

