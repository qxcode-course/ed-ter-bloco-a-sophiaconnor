package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	info int		
	next *Node
	previous *Node
}
type LList struct {
	head *Node	
	tail *Node
	size int
}

func NewLList() *LList {
	return &LList{}
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) String() string {
	if l.head == nil {
		return "[]"
	}

	values := make([]string, 0, l.size)
	for node := l.head; node != nil; node = node.next {
		values = append(values, strconv.Itoa(node.info))
	}
	return "[" + strings.Join(values, ", ") + "]"
}
func (l *LList) PushBack(value int) {
	node := &Node{info: value, previous: l.tail}
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	if l.head == nil {
		l.head = node
	}
	l.size++
}
func (l *LList) PushFront(info int) {
    node := &Node{info: info, next: l.head}
    if l.head != nil {
        l.head.previous = node
    }
    l.head = node
    if l.tail == nil {
        l.tail = node
    }
    l.size++
}
func (l *LList) Front() *Node {
	return l.head
}
func (l *LList) Back() *Node {
	return l.tail
}
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.previous
}
func (n *Node) Value() int {
	return n.info
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
			num, _ := strconv.Atoi(v)
			ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
			num, _ := strconv.Atoi(v)
			ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "walk":
			 fmt.Print("[ ")
			 for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Print("]\n[ ")
			 for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Println("]")
		case "replace":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	node.Value = newvalue
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "insert":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Insert(node, newvalue)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
