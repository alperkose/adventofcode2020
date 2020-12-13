package main

import (
	"bufio"
	"io"
	"strings"
)

func ReadSource(source io.Reader) []string {
	bufferedSource := bufio.NewReader(source)
	rows := []string{}
	for {
		line, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		rows = append(rows, strings.TrimFunc(string(line), trimLine))
	}
	return rows
}

const trimChars = "\t "

func trimLine(r rune) bool {
	return strings.IndexRune(trimChars, r) >= 0
}
