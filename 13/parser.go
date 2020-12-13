package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Parse(source io.Reader) (int, []Bus) {
	bufferedReader := bufio.NewReader(source)

	line, _, _ := bufferedReader.ReadLine()
	earliestTime, _ := strconv.Atoi(string(line))

	line, _, _ = bufferedReader.ReadLine()
	parts := strings.Split(string(line), ",")
	buses := []Bus{}
	for _, busStr := range parts {
		if busStr == "x" {
			buses = append(buses, IgnoredBus)
		} else {
			bus, _ := strconv.Atoi(busStr)
			buses = append(buses, Bus(bus))
		}
	}
	return earliestTime, buses
}
