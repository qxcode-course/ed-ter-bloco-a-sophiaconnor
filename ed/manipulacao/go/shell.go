package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var result []int
	for _, v := range vet {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

func getCalmWomen(vet []int) []int {
	var result []int
	for _, v := range vet {
		if v > -10 && v < 0 {
			result = append(result, v)
		}
	}
	return result
}

func sortVet(vet []int) []int {
	result := make([]int, len(vet))
	copy(result, vet)
	sort.Ints(result)
	return result
}

func sortStress(vet []int) []int {
	result := make([]int, len(vet))
	copy(result, vet)
	
	sort.Slice(result, func(i, j int) bool {
		// Converte primeiro para float, depois divide, para não perder as casas decimais
		valI := math.Abs(float64(result[i]) / 10.0)
		valJ := math.Abs(float64(result[j]) / 10.0)
		return valI < valJ
	})
	
	return result
}


func reverse(vet []int) []int {
	result := make([]int, len(vet))
	for i, v := range vet {
		result[len(vet)-1-i] = v
	}
	return result
}

func unique(vet []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, v := range vet {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func repeated(vet []int) []int {
	count := make(map[int]int)
	for _, v := range vet {
		count[v]++
	}
	var keys []int
	for v, c := range count {
		if c > 1 {
			keys = append(keys, v)
		}
	}
	sort.Ints(keys)
	var result []int
	for _, v := range keys {
		for i := 0; i < count[v]-1; i++ {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
