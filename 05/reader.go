package main

import (
	"bufio"
	"io"
)

func ReadSource(source io.Reader) []string {
	bufferedSource := bufio.NewReader(source)
	var inputs []string
	for {
		line, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		inputs = append(inputs, string(line))
	}
	return inputs
}
