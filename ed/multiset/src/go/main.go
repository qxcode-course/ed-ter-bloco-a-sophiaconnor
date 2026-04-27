package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
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

func NewMultiSet(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Insert(value int) error {
	// Find the position to insert to keep sorted
	pos := 0
	for pos < v.size && v.data[pos] < value {
		pos++
	}
	if v.size == v.capacity {
		newCapacity := v.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		v.Reserve(newCapacity)
	}
	for i := v.size; i > pos; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[pos] = value
	v.size++
	return nil
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}
	newData := make([]int, newCapacity)
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}
	v.data = newData
	v.capacity = newCapacity
}

func (v *Vector) Contains(value int) bool {
	low := 0
	high := v.size - 1
	for low <= high {
		mid := (low + high) / 2
		if v.data[mid] == value {
			return true
		} else if v.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}
func (v *Vector) Erase(value int) error {
	idx := -1
	low := 0
	high := v.size - 1
	for low <= high {
		mid := (low + high) / 2
		if v.data[mid] == value {
			idx = mid
			break
		} else if v.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if idx == -1 {
		return errors.New("value not found")
	}
	for i := idx; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	return nil
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		parts := strings.Fields(line)
		fmt.Println(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				err := ms.Insert(value)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if err := ms.Erase(value); err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(ms.Contains(value))
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
