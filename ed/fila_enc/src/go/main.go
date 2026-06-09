package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Fila struct {
	itens []int
}

func (q *Queue[T]) Enqueue(value T){
	q.size++
	newNode := &Node[T]{Value: value}
	if q.tail != nil {
		q.tail.next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
}
func (q *Queue[T]) Dequeue() (T, bool){
	if q.size == 0 {
		var zero T
		return zero, false
	}
	value := q.head.Value
	q.head = q.head.next
	q.size--
	if q.size == 0 {
		q.tail = nil
	}
	return value, true
}
func (q *Queue[T]) Peek() (T, bool){
	if q.size == 0 {
		var zero T
		return zero, false
	}
	return q.head.Value, true
}
func (q *Queue[T]) Size() int {
	return q.size
}
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
// func (q *Queue[T]) Clear()
func (f *Fila) Push(value int) {
	f.itens = append(f.itens, value)
}
func (f *Fila) Pop() (int, bool) {
	if len(f.itens) == 0 {
		return 0, false
	}
	value := f.itens[0]
	f.itens = f.itens[1:]
	return value, true
}
type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			for _, arg := range args[1:] {
			value, _ := strconv.Atoi(arg)
			queue.Enqueue(value)
			}
		case "pop":
			// if _, ok := queue.Dequeue(); !ok {
			// 	fmt.Println("falha: fila vazia")
			// }
		case "peek":
			if value, ok := queue.Peek(); ok {
			fmt.Println(value)
			} else {
			fmt.Println("falha: fila vazia")
			}
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
