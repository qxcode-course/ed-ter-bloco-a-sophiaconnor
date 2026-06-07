package main
import "fmt"
type char byte
func main() {
    var num int
    fmt.Scan(&num)
    if num == 4 {
        //var matriz [4][4]char
        matriz := make([][]char, 4)
    }else {
        //var matriz [9][9]char
        matriz := make([][]char, 9)
    }
    for i := 0; i < num; i++ {
        //matriz[i] = make([]char, num)
        for j := 0; j < num; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }
        fmt.Println(matriz)
    }