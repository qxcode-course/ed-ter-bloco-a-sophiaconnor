package main

import "fmt"

func totalNQueens(n int) int {
    count := 0
    var solve func(row, cols, diag1, diag2 int)
    solve = func(row, cols, diag1, diag2 int) {
        if row == n {
            count++
            return
        }

        available := ((1 << n) - 1) & ^(cols | diag1 | diag2)
        for available != 0 {
            bit := available & -available
            available -= bit
            solve(row+1, cols|bit, (diag1|bit)<<1, (diag2|bit)>>1)
        }
    }

    solve(0, 0, 0, 0)
    return count
}

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }
    fmt.Println(totalNQueens(n))
}
