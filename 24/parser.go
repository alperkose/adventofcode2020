package main

import (
	"bufio"
	"io"
)

func Parse(source io.Reader) []TileReference {
	buffer := bufio.NewReader(source)
	tiles := []TileReference{}
	for {
		line, _, err := buffer.ReadLine()
		if err == io.EOF {
			break
		}
		tiles = append(tiles, TileReferenceFromString(string(line)))

	}
	return tiles
}
