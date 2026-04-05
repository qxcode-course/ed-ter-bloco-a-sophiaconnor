package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	result := "["
	for i, v := range vet {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d", v)
	}
	return result + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	return "[" + tostrrevHelper(vet, len(vet)-1)
}

func tostrrevHelper(vet []int, i int) string {
	if i == 0 {
		return fmt.Sprintf("%d]", vet[i])
	}
	return fmt.Sprintf("%d, %s", vet[i], tostrrevHelper(vet, i-1))
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	reverseHelper(vet, 0, len(vet)-1)
}

func reverseHelper(vet []int, inicio, fim int) {
	if inicio >= fim {
		return
	}
	vet[inicio], vet[fim] = vet[fim], vet[inicio]
	reverseHelper(vet, inicio+1, fim-1)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	return sumHelper(vet, 0)
}

func sumHelper(vet []int, i int) int {
	if i >= len(vet) {
		return 0
	}
	return vet[i] + sumHelper(vet, i+1)
	//sum(vet []int)
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	_ = vet
	return multHelper(vet, 0)
}
func multHelper(vet []int, i int) int {
	if i >= len(vet) {
		return 1
	}
	return vet[i] * multHelper(vet, i+1)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice

func min(vet []int) int {
    if len(vet) == 0 {
        return -1
    }

    var rec func(i int) int
    rec = func(i int) int {
        if i == len(vet)-1 {
            return i
        }
        minIdx := rec(i + 1)
        if vet[i] <= vet[minIdx] {
            return i
        }
        return minIdx
    }

    return rec(0)
}
func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
