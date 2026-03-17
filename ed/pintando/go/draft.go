package main

import (
	"fmt"
	"math"
)

func main() {
	var l1 float64
	var l2 float64
	var l3 float64
	var sub1 float64
	var sub2 float64
	var sub3 float64
	var soma float64
	var div float64
	var num float64
	raiz := 0.0
	fmt.Scanf("%f", &l1)
	fmt.Scanf("%f", &l2)
	fmt.Scanf("%f", &l3)
	soma = l1 + l2 + l3
	div = soma / 2
	sub1 = div - l1
	sub2 = div - l2
	sub3 = div - l3
	num = div * sub1 * sub2 * sub3
	raiz += math.Sqrt(num)
	fmt.Printf("%.2f\n", raiz)
}
