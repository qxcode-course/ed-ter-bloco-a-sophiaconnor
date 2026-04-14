package main
import "fmt"
func main() {
	var qtd_pessoas_fila int
	fmt.Scan(&qtd_pessoas_fila)
	fila := make ([]int, qtd_pessoas_fila)
	for i := 0; i < qtd_pessoas_fila; i++ {
		fmt.Scan(&fila[i])
	}
	var qtd_sairam_fila int
	fmt.Scan(&qtd_sairam_fila)
	sairam := make([]int, qtd_sairam_fila)
	for i:=0; i<qtd_sairam_fila; i++{
		fmt.Scan(&sairam[i])
	}
	map 
}
