package main

import "fmt"
func matchingStrings(consulta []string, busca []string) []int {
	freq := make(map[string]int)
	for _, s := range consulta {
		freq[s]++
	}

	resultado := make([]int, len(busca))
	for i, b := range busca {
		resultado[i] = freq[b]
	}
	return resultado
}
func main() {
	var tam_consulta int
	fmt.Scan(&tam_consulta)

	consulta := make([]string, tam_consulta)
	for i := 0; i < tam_consulta; i++ {
		fmt.Scan(&consulta[i])
	}

	var tam_busca int
	fmt.Scan(&tam_busca)

	busca := make([]string, tam_busca)
	for i := 0; i < tam_busca; i++ {
		fmt.Scan(&busca[i])
	}
	resultado := matchingStrings(consulta, busca)
	for i, r := range resultado {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(r)
	}
	fmt.Println()
}
