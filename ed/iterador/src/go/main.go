package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)-1
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) {
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1
	return i.data[i.index]
}
type ReverseIterator = Iterator
// Cria o iterador com uma cópia invertida dos dados
func (l *MyList) ReverseIterator() *ReverseIterator { 
	reversed := make([]int, len(l.data))
	for i, v := range l.data {
		reversed[len(l.data)-1-i] = v
	}
	return &Iterator{data: reversed, index: -1}
}

// Apenas uma struct para rastrear a posição atual
type CyclicIterator struct {
	data  []int
	index int
}

// 1. A função geradora (adicionada na sua MyList)
func (l *MyList) CyclicIterator() *CyclicIterator {
	return &CyclicIterator{data: l.data, index: -1}
}

// 2. O único método necessário para o laço da sua main funcionar
func (i *CyclicIterator) Next() int {
	if len(i.data) == 0 {
		return 0 // Proteção caso a lista esteja vazia
	}
	// Avança o índice e usa o resto da divisão (%) para voltar ao início (0) ciclicamente
	i.index = (i.index + 1) % len(i.data)
	return i.data[i.index]
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			 fmt.Print("[ ")
			 for it := mylist.ReverseIterator(); it.HasNext(); {
			 	fmt.Printf("%v ", it.Next())
			 }
			 fmt.Println("]")
		case "cyclic":
			 qtd, _ := strconv.Atoi(args[1])
			 fmt.Print("[ ")
			 it := mylist.CyclicIterator()
			 for range qtd {
			 	fmt.Printf("%v ", it.Next())
			 }
			 fmt.Println("]")
		}
	}

}
