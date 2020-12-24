package main

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Parse(source io.Reader) ([]Rule, []string) {
	buffer := bufio.NewReader(source)
	rules := []Rule{}
	samples := []string{}
	parseRules := true
	for {
		line, _, err := buffer.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		if len(strings.Trim(string(line), " ")) == 0 {
			parseRules = false
			continue
		}
		if parseRules {
			rule := ParseLine(string(line))
			rules = append(rules, rule)
		} else {
			samples = append(samples, string(line))
		}

	}
	sort.Slice(rules, func(a, b int) bool { return rules[a].order < rules[b].order })
	return rules, samples
}

func ParseLine(line string) Rule {
	ind := strings.Index(line, ": ")
	order, _ := strconv.Atoi(line[:ind])
	if line[ind+2] == '"' {
		return ValueRule(order, rune(line[ind+3]))
	}
	parts := strings.Split(line[ind+2:], " | ")
	path := strings.Split(parts[0], " ")
	pathIndices := make([]int, len(path))
	for i, p := range path {
		pathIndices[i], _ = strconv.Atoi(p)
	}
	if len(parts) > 1 {
		path = strings.Split(parts[1], " ")
		pathIndices2 := make([]int, len(path))
		for i, p := range path {
			pathIndices2[i], _ = strconv.Atoi(p)
		}
		return DoublePathRule(order, pathIndices, pathIndices2)
	}
	return SinglePathRule(order, pathIndices)
}
