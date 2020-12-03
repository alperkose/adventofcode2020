package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	expenses := pipeIn()
	fmt.Println(expenses)
	n1, n2 := findAPair(expenses, 2020)
	fmt.Println(n1, n2, (n1 * n2))
	n1, n2, n3 := findATriple(expenses, 2020)
	fmt.Println(n1, n2, n3, (n1 * n2 * n3))
}

func findAPair(expenses []int, target int) (int, int) {
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == target {
				return expenses[i], expenses[j]
			}
		}
	}
	return 2020, 0
}
func findAPairWithSort(expenses []int, target int) (int, int) {
	sort.Ints(expenses)
	return findAPairFromSorted(expenses, target)
}

func findAPairFromSorted(expenses []int, target int) (int, int) {
	for i := 0; i < len(expenses); i++ {
		for j := len(expenses) - 1; j > i; j-- {
			sum := expenses[i] + expenses[j]
			if sum == target {
				return expenses[i], expenses[j]
			} else if sum < target {
				break
			}
		}
	}
	return 2020, 0
}

func findATriple(expenses []int, target int) (int, int, int) {
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == target {
					return expenses[i], expenses[j], expenses[k]
				}
			}
		}
	}
	return 2020, 0, 0
}

func findATripleWithSort(expenses []int, target int) (int, int, int) {
	sort.Ints(expenses)
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			for k := len(expenses) - 1; k > j; k-- {
				sum := expenses[i] + expenses[j] + expenses[k]
				if sum == target {
					return expenses[i], expenses[j], expenses[k]
				} else if sum < target {
					break
				}
			}
		}
	}
	return 2020, 0, 0
}

func pipeIn() []int {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		log.Fatal("The command is intended to work with pipes")
	}

	reader := bufio.NewReader(os.Stdin)
	var output []int

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		i, err := strconv.Atoi(string(input))
		if err != nil {
			log.Fatalf("the line doesn't contain number: %s", input)
		}
		output = append(output, i)
	}
	return output
}
