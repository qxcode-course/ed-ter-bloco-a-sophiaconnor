package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	consulta := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &consulta[i])
	}

	var M int
	fmt.Scanf("%d", &M)

	busca := make([]string, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%s", &busca[i])
	}
    fmt.Println(N, M, consulta, busca)
	/*freq := make(map[string]int)
	for _, s := range consulta {
		freq[s]++
	}

	for i, b := range busca {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(freq[b])
	}
	fmt.Println()*/
}
