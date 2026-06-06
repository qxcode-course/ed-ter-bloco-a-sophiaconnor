package main
import "fmt"
func subsetSum(nums []int, resultado_soma int) bool {
    qtd := len(nums)
    eh_possivel := make([][]bool, qtd+1)
    for i := range eh_possivel {
        eh_possivel[i] = make([]bool, resultado_soma+1)
        eh_possivel[i][0] = true
    }
    for i := 1; i <= qtd; i++ {
        for j := 1; j <= resultado_soma; j++ {
            if nums[i-1] > j {
                eh_possivel[i][j] = eh_possivel[i-1][j]
            } else {
                eh_possivel[i][j] = eh_possivel[i-1][j] || eh_possivel[i-1][j-nums[i-1]]
            }
        }
    }
    return eh_possivel[qtd][resultado_soma]
}
func main() {
    var qtd_num int
    fmt.Scan(&qtd_num)
    var soma int
    fmt.Scan(&soma)
    vet := make([]int, qtd_num)
    for i := 0; i < qtd_num; i++ {
        fmt.Scan(&vet[i])
    }
    result := subsetSum(vet, soma)
    if result {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
