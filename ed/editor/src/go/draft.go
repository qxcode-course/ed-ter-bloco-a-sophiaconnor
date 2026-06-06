package main

import (
	"bufio"
	"fmt"
	"os"
)

func editar(comando string) string {
	var texto []rune
	cursor := 0

	for _, c := range comando {
		if c == 'R' {
			texto = append(texto[:cursor], append([]rune{'\n'}, texto[cursor:]...)...)
			cursor++
		} else {
			texto = append(texto[:cursor], append([]rune{c}, texto[cursor:]...)...)
			cursor++
		}
	}
    resultado := append(texto[:cursor], append([]rune{'|'}, texto[cursor:]...)...)
    return string(resultado)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		entrada := scanner.Text()
		resultado := editar(entrada)
		fmt.Println(resultado)
	}
}
