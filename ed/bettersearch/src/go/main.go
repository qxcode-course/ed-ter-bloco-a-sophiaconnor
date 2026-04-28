package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, size int, value int) (bool, int) {
	low, high := 0, size
	for low < high {
		mid := (low + high) / 2
		if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low < size && slice[low] == value, low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	i := 1
	for {
		if i >= len(parts)-1 {
			break
		}
		elem := parts[i]
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
		i++
	}
	size := i - 1
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, size, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
