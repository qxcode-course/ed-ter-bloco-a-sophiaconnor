package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), 
		size:     0,
		capacity: capacity,
	}
}


func (v *Vector) Clear() {
	v.size = 0
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	} 
	newData := make([]int, newCapacity, newCapacity) 
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}
	v.data = newData
	v.capacity = newCapacity
}

func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index > v.size {
		return errors.New("index out of range")
	}
	if v.size == v.capacity {
		newCapacity := v.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		v.Reserve(newCapacity)
	}
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.size++
	return nil
}

func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1
}


func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return i
		}
	}
	return -1
}

func (v *Vector) String() string {
	if v.size == 0 {
		return "[]"
	}
	var builder strings.Builder
	builder.WriteString("[")
	for i := 0; i < v.size; i++ {
		if i > 0 {
			builder.WriteString(", ")
		}
		fmt.Fprintf(&builder, "%d", v.data[i])
	}
	builder.WriteString("]")
	return builder.String()
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	v := NewSet(0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		
		fmt.Printf("$%s\n", line)

		parts := strings.Fields(line)
		cmd := parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			capacity, _ := strconv.Atoi(parts[1])
			v = NewSet(capacity)
		case "insert":
			for i := 1; i < len(parts); i++ {
				value, _ := strconv.Atoi(parts[i])
				if !v.Contains(value) {
					// Lógica para manter ordenado
					pos := 0
					for pos < v.size && v.data[pos] < value {
						pos++
					}
					v.Insert(pos, value)
				}
			}
		case "show":
			fmt.Println(v.String())
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		/*case "erase":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			if index != -1 {
				v.Erase(index)
			}*/
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
