package main

import (
	"fmt"
	"strconv"
	"strings"
)

//verifica se esse valor pode ser usado nesse índice, ou seja, se não existe o mesmo número a uma distância menor ou igual a distância dada
func (p *problem) fit(index int, value rune) bool {
    _, _ = index, value
    return false
}
func (p *problem) solve(index int) bool {
    //se chegou ao fim, retorne true
    //se o valor atual não for um ponto, continue para o próximo índice
    //se for um ponto:
    //faça um laço e chame a recursão para cada número possível (0 a distância)
    //se algum der certo, retorne true
    //se nenhum der certo, recoloque vazio e retorne false
    _ = index
    return false
}
func posso_inserir(line []string, index int, distancia int, digito string) bool {
    for i := 1; i <= distancia; i++ {
        if index-i >= 0 && line[index-i] == digito {
            return false
        }
    }
    for i := 1; i <= distancia; i++ {
        if index+i < len(line) && line[index+i] == digito {
            return false
        }
    }
    return true
}
func backtrack(sequencia []string, index int, distancia int) bool {
    if index == len(sequencia) {
        return true
    }
    if sequencia[index] != "." {
        return backtrack(sequencia, index + 1, distancia)
    }
    for num := 0; num <= distancia; num++ {
        digito := strconv.Itoa(num)
        if posso_inserir(sequencia, index, distancia, digito) {
            sequencia[index] = digito
            if backtrack(sequencia, index + 1, distancia) {
                return true
            }
            sequencia[index] = "."
        }
    }
    return false
}

func main() {
    var sequencia string
    fmt.Scan(&sequencia)
    var distancia int
    fmt.Scan(&distancia)
    s := strings.Split(sequencia, "")
    if backtrack(s, 0, distancia) {
        fmt.Println(strings.Join(s, ""))
    }
    
}