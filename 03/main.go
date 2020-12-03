package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Slope struct {
	Down, Right int
}

func main() {
	var slopes slopeFlags

	flag.Var(&slopes, "s", "Slope, comma separated integers: -s 1,3")
	flag.Parse()

	if len(slopes) == 0 {
		fmt.Fprint(os.Stderr, "No slope provided as argument (ex: -s 1,3)\n")
		os.Exit(1)
	}

	rows := ReadMap(os.Stdin)
	arborealMultiplier := 1
	for _, c := range slopes {
		arborealMultiplier *= CountTrees(rows, c.Down, c.Right)
	}

	fmt.Println(arborealMultiplier)
}

func CountTrees(rows []string, downSlope, rightSlope int) int {
	mapWidth := len(rows[0])
	treeCount := 0
	currentLoc := 0

	for i := 0; i < len(rows); i += downSlope {
		// fmt.Println(rows[i], string(rows[i][currentLoc]), currentLoc)
		if rows[i][currentLoc] == '#' {
			treeCount++
		}
		currentLoc = (currentLoc + rightSlope) % mapWidth
	}
	return treeCount
}

func ReadMap(source io.Reader) []string {
	var output []string
	bufferedSource := bufio.NewReader(source)
	for {
		input, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, string(input))
	}
	return output
}

type slopeFlags []Slope

func (i *slopeFlags) String() string {
	return "my string representation"
}

func (i *slopeFlags) Set(value string) error {
	parts := strings.Split(value, ",")
	if len(parts) != 2 {
		return fmt.Errorf("Expected comma separated integers")
	}

	down, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("Expected comma separated integers")
	}

	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("Expected comma separated integers")
	}
	*i = append(*i, Slope{down, right})
	return nil
}
