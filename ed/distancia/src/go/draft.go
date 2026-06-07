package main

import (
	"fmt"
	"strconv"
	"strings"
)

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