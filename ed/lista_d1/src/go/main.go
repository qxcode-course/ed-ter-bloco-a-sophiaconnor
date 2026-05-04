package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LList struct {
	head *Node
	tail *Node
	size int
}

func NewLList() *LList {
	return &LList{}
}

func (ll *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	current := ll.head
	for current != nil {
		if current != ll.head {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", current.value))
		current = current.next
	}
	sb.WriteString("]")
	return sb.String()
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{value: value}
	if ll.size == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head.prev = newNode
		ll.head = newNode
	}
	ll.size++
}
func (ll *LList) PushBack(value int) {
	newNode := &Node{value: value}
	if ll.size == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}
	ll.size++
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
			//fmt.Println(ll.Size())
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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
