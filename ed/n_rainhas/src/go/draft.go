package main

import "fmt"

func solve(linha, num int, cols, diag1, diag2 []bool, count int) {
	if linha == num {
		count++
		return
	}

	for col := 0; col < num; col++ {
		if cols[col] || diag1[linha+col] || diag2[linha-col+num-1] {
			continue
		}
		cols[col] = true
		diag1[linha+col] = true
		diag2[linha-col+num-1] = true
		solve(linha+1, num, cols, diag1, diag2, count)
		cols[col] = false
		diag1[linha+col] = false
		diag2[linha-col+num-1] = false
	}
}

func totalNQueens(num int) int {
	cols := make([]bool, num)
	diag1 := make([]bool, 2*num-1)
	diag2 := make([]bool, 2*num-1)
	count := 0
	solve(0, num, cols, diag1, diag2, count)
	return count
}

func main() {
	var num int
	fmt.Println(totalNQueens(num))
}
