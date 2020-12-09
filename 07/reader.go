package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const MARKER_BAGS_CONTAIN = " bags contain "
const MARKER_NO_CONTENT = "no other bags."

func ReadSource(source io.Reader) *Rules {
	bufferedSource := bufio.NewReader(source)

	var rules = NewRules()
	for {
		line, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		statement := string(line)
		index := strings.Index(statement, MARKER_BAGS_CONTAIN)
		bagColor := statement[:index]
		contentStr := statement[index+len(MARKER_BAGS_CONTAIN):]
		if contentStr == MARKER_NO_CONTENT {
			rules.Add(NewBag(bagColor))
			continue
		}
		pieces := strings.Split(contentStr[:len(contentStr)-1], ", ")
		nBags := make([]Content, 0, len(pieces))
		for _, piece := range pieces {
			i1stSp := strings.Index(piece, " ")
			amount, _ := strconv.Atoi(piece[:i1stSp])
			color := piece[i1stSp+1 : strings.Index(piece, "bag")-1]
			nBags = append(nBags, Content{color, amount})
		}
		rules.Add(NewBag(bagColor, nBags...))
	}
	return rules
}
