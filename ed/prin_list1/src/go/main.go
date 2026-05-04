package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
// mostra a lista percorrendo linearmente do início ao fim
func ToStr(l *DList[int], sword *DNode[int]) string {
	var sb strings.Builder
	sb.WriteString("[ ")
	// Use l.root.next para percorrer a lista real sem entrar em loop infinito
	for it := l.root.next; it != l.root; it = it.next {
		if it == sword {
			sb.WriteString(fmt.Sprintf("%d> ", it.Value))
		} else {
			sb.WriteString(fmt.Sprintf("%d ", it.Value))
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// move para frente na lista de forma CIRCULAR
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil {
		return nil
	}
	next := it.next
	// Se o próximo for a sentinela, pula ela e volta para o primeiro
	if next == l.root {
		return l.root.next
	}
	return next
}


func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
