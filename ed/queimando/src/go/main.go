package main

import (
	"bufio"
	"fmt"
	"os"
)
type Pos struct{
	l, c int
}

func (p Pos) getNeighboors() []Pos{
	l := p.l
	c := p.c
	return []Pos {
		Pos{l: l, c: c-1},
		Pos{l: l, c: c+1},
		Pos{l: l-1, c: c},
		Pos{l: l+1, c: c},
	}
}
func (s *Stack[T]) Empty() bool{
	if len(s.data) == 0 {
		return true
	}
	return false
}
func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	//_ , _ , _ = mat, l, c
	stack.Push(Pos{l, c})
	for !stack.Empty(){
		atual := stack.Pop()
		if atual.l <0 || atual.l >= len(grid) || atual.c < 0 || atual.c >= len(grid[0]){
			continue
			}
		if grid[atual.l][atual.c] != '#'{
			continue
		}
		grid[atual.l][atual.c] ='o'
		for _, viz := range atual.getNeighboors(){
			stack.Push(viz)
		}
	}
}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
