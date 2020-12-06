package main

import (
	"bufio"
	"io"
	"strings"
)

func ReadSource(source io.Reader) []*Group {
	bufferedSource := bufio.NewReader(source)

	var groups []*Group
	var group = NewGroup()
	for {
		line, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			groups = append(groups, group)
			break
		}
		if len(strings.Trim(string(line), " ")) == 0 {
			groups = append(groups, group)
			group = NewGroup()
		} else {
			group.Add(string(line))
		}
	}
	return groups
}
