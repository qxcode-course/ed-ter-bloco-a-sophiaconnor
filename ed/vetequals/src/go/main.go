package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EhVazio(vet []int) bool {
	if len(vet) == 0 {
		return true
	}
	//return fmt.Sprint(vet) == "[]"
	return false
}

// não altere a assinaturafunc equals(a []int, b []int) bool {
	// 1. Verifica se ambos terminaram ao mesmo tempo (iguais)
func equals(a []int, b []int) bool {
	//_, _ = a, b
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	// não use a função len para ver ou comparar os tamanhos
	// utilize a função EhVazio para ver se o vetor é vazio
	if EhVazio(a) && EhVazio(b) {
		return true
	}
	if EhVazio(a) || EhVazio(b) {
		return false
	}
	if a[0] != b[0] {
		return false
	}
	return 	equals(a[1:], b[1:])
}

	// você só pode consultar o primeiro elemento do vetor
	// e não pode usar nenhum tipo de laço
	// Use recursao para consultar os outros elementos equals(a[1:], b[1:])
	// não altere o protótipo da função nem crie funções auxiliares
	//return false

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	vet1 := str2slice(scanner.Text())
	scanner.Scan()
	vet2 := str2slice(scanner.Text())
	if equals(vet1, vet2) {
		fmt.Println("iguais")
	} else {
		fmt.Println("diferentes")
	}
}

func str2slice(line string) []int {
	parts := strings.Fields(line)
	nums := make([]int, 0)
	for i := 1; i < len(parts)-1; i++ {
		value, _ := strconv.Atoi(parts[i])
		nums = append(nums, value)
	}
	return nums
}
