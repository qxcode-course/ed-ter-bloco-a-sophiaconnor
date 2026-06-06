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

type DIterator struct {
	l   *DList[int]
	cur *DNode[int]
}

func NewIterator(l *DList[int]) *DIterator {
	return &DIterator{l: l, cur: l.Front()}
}

func (it *DIterator) Curr() *DNode[int] {
	return it.cur
}

func (it *DIterator) Next() *DNode[int] {
	next := it.cur.Next()
	if next == it.l.End() {
		next = it.l.Front()
	}
	it.cur = next
	return it.cur
}

func (it *DIterator) PeekNext() *DNode[int] {
	next := it.cur.Next()
	if next == it.l.End() {
		next = it.l.Front()
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

	sword := NewIterator(l)
	for range chosen - 1 {
		sword.Next()
	}

	for range qtd - 1 {
		fmt.Println(ToStr(l, sword.Curr()))
		l.Erase(sword.PeekNext())
		sword.Next()
	}
	fmt.Println(ToStr(l, sword.Curr()))
}
