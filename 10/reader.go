package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ReadSource(source io.Reader) []int {
	bufferedSource := bufio.NewReader(source)
	inputs := []int{}

	for {
		line, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		if len(strings.Trim(string(line), " ")) == 0 {
			continue
		}
		num, _ := strconv.Atoi(string(line))
		inputs = append(inputs, num)
	}
	return inputs
}
