package main

import "fmt"

func main() {
	var tam_consulta int
	fmt.Scanf("%d", &tam_consulta)

	consulta := make([]string, tam_consulta)
	for i := 0; i < tam_consulta; i++ {
		fmt.Scanf("%s", &consulta[i])
	}

	var tam_busca int
	fmt.Scanf("%d", &tam_busca)

	busca := make([]string, tam_busca)
	for i := 0; i < tam_busca; i++ {
		fmt.Scanf("%s", &busca[i])
	}
    fmt.Println(tam_consulta, tam_busca, consulta, busca)
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
