package main

import (
	"bufio"
	"io"
	"strings"
)

func ReadSource(source io.Reader, convertToPass func(string) Passport) []Passport {
	passports := []Passport{}
	bufferedSource := bufio.NewReader(source)

	var sb strings.Builder
	for {
		line, _, err := bufferedSource.ReadLine()
		if err != nil && err == io.EOF {
			passports = append(passports, convertToPass(sb.String()))
			sb.Reset()
			break
		}
		if len(strings.Trim(string(line), " ")) == 0 {
			passports = append(passports, convertToPass(sb.String()))
			sb.Reset()
		} else {
			sb.Write(line)
			sb.WriteString(" ")
		}
	}
	return passports
}
