package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var checkerFlag string
	flag.StringVar(&checkerFlag, "checker", "singleMatch", "help message for flag n")
	flag.Parse()

	var checker Checker
	if checkerFlag == "count" {
		checker = CharCount
	} else {
		checker = ExactlyOneMatchedIndex
	}

	policies := pipeIn()
	validCount := 0
	for _, policy := range policies {
		if CheckWith(policy, checker) {
			validCount++
		}
	}
	fmt.Println("Valid policies:", validCount)
}

type Checker func(int, int, string, string) bool

func CheckWith(policy string, fCheck Checker) bool {
	atLeast, atMost, checkChar, pwd, err := ParsePolicy(policy)
	if err != nil {
		return false
	}
	return fCheck(atLeast, atMost, checkChar, pwd)
}

func CharCount(atLeast, atMost int, checkChar, pwd string) bool {
	count := strings.Count(pwd, checkChar)
	return atLeast <= count && count <= atMost
}

func ExactlyOneMatchedIndex(firstIndex, secondIndex int, checkChar, pwd string) bool {
	firstOccurence := pwd[firstIndex-1] == checkChar[0]
	secondOccurence := pwd[secondIndex-1] == checkChar[0]
	return !(firstOccurence && secondOccurence) && (firstOccurence || secondOccurence)
}

func ParsePolicy(policy string) (int, int, string, string, error) {
	parts := strings.Split(policy, ": ")
	if len(parts) < 2 {
		return 0, 0, "", "", fmt.Errorf("Invalid policy: %s", policy)
	}

	rules := strings.Split(parts[0], " ")
	if len(rules) < 2 {
		return 0, 0, "", "", fmt.Errorf("Invalid policy: %s", policy)
	}

	counts := strings.Split(rules[0], "-")
	if len(counts) < 2 {
		return 0, 0, "", "", fmt.Errorf("Invalid policy: %s", policy)
	}

	atLeast, err := strconv.Atoi(string(counts[0]))
	if err != nil {
		return 0, 0, "", "", err
	}

	atMost, err := strconv.Atoi(string(counts[1]))
	if err != nil {
		return 0, 0, "", "", err
	}
	return atLeast, atMost, rules[1], parts[1], nil
}

func pipeIn() []string {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		log.Fatal("The command is intended to work with pipes")
	}

	reader := bufio.NewReader(os.Stdin)
	var output []string

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, string(input))
	}
	return output
}
