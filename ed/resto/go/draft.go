package main

import "fmt"

func resultados(num int) {
	if num == 0 {
		return
	}

	resto := num % 2
	quociente := num / 2
	resultados(quociente)
	fmt.Println(quociente, resto)
}

func main() {
	var num int
	fmt.Scan(&num)
	resultados(num)
}
