package main

import (
	"container/list"
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *list.List, sword *list.Element) string {
	var sb strings.Builder
	sb.WriteString("[ ")
	for it := l.Front(); it != nil; it = it.Next() {
		if it == sword {
			if it.Value.(int) > 0 {
				sb.WriteString(fmt.Sprintf("%d> ", it.Value))
			} else {
				sb.WriteString(fmt.Sprintf("<%d ", it.Value))
			}
		} else {
			sb.WriteString(fmt.Sprintf("%d ", it.Value))
		}
	}
	sb.WriteString("]")
	return sb.String()
}
// move para frente na lista circular
func Next(l *list.List, it *list.Element) *list.Element {
    if it.Next() != nil {
        return it.Next()
    }
    return l.Front()
}

// move para tras na lista circular
func Prev(l *list.List, it *list.Element) *list.Element {
    if it.Prev() != nil {
        return it.Prev()
    }
    return l.Back()
}

func main() {
	var qtd, chosen, fase int
	fmt.Scan(&qtd, &chosen, &fase)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		if sword.Value.(int) > 0 {
			l.Remove(Next(l, sword))
			sword = Next(l, sword)
		} else {
			l.Remove(Prev(l, sword))
			sword = Prev(l, sword)
		}
	}
	fmt.Println(ToStr(l, sword))
}
