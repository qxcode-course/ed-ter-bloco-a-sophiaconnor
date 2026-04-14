package main
import "fmt"
func main() {
    var qtd_animais int
    fmt.Scan(&qtd_animais)
    animais := make([]int, qtd_animais)
    for i := 0; i < qtd_animais; i++ {
        fmt.Scan(&animais[i])
    }
    qtd_casais := 0
    //casados := make([]int, 0)
    var repetido int
    for i := 0; i < qtd_animais; i++ {
        for j := i + 1; j < qtd_animais; j++ {
            if animais[i] == (-animais[j]) {
                qtd_casais++
                //casados = append(casados, animais[i])
                animais[i] = repetido
                break
            }
        }
    }
    fmt.Println(qtd_casais)
}