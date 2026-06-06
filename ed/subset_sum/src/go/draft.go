package main
import "fmt"
func main() {
    var qtd_num int
    fmt.Scanf("%d", &qtd_num)
    vet := make([]int, qtd_num)
    for i := 0; i < qtd_num; i++ {
        fmt.Scanf("%d", &vet[i])
    }
    for i, v := range vet {
        fmt.Printf("vet[%d] = %d\n", i, v)
    }
    fmt.Println("Slice:", vet)
}
