package main

import (
	"strings"
	"testing"
)

func TestReadMap(t *testing.T) {
	testSource := strings.NewReader(".#..\n#...")
	rows := ReadMap(testSource)
	if len(rows) != 2 {
		t.Errorf("Should've parsed two lines, parsed %d lines", len(rows))
	}
	if rows[0] != ".#.." {
		t.Errorf("First line, expected \".#..\" got %s", rows[0])
	}
	if rows[1] != "#..." {
		t.Errorf("Second line, expected \"#...\" got %s", rows[1])
	}
}

func TestTreeCountingOnOpenGround(t *testing.T) {
	rows := []string{"....", "...."}
	actualTrees := CountTrees(rows, 1, 3)
	if actualTrees != 0 {
		t.Errorf("TreeCount, expected 0, got %d", actualTrees)
	}
}
func TestTreeCountingOnAllTrees(t *testing.T) {
	rows := []string{"####", "####"}
	actualTrees := CountTrees(rows, 1, 3)
	if actualTrees != 2 {
		t.Errorf("TreeCount, expected 2, got %d", actualTrees)
	}
}
func TestTreeCountingOnPartialTrees(t *testing.T) {
	rows := []string{".#.#", "...#"}
	actualTrees := CountTrees(rows, 1, 3)
	if actualTrees != 1 {
		t.Errorf("TreeCount, expected 1, got %d", actualTrees)
	}
}

func TestTreeCountingWithSlidingMap(t *testing.T) {
	rows := []string{".#.#", "...#", "..#."}
	actualTrees := CountTrees(rows, 1, 3)
	if actualTrees != 2 {
		t.Errorf("TreeCount, expected 2, got %d", actualTrees)
	}
}

func TestTreeCountingWithSkippingRow(t *testing.T) {
	rows := []string{".#.#", "####", "...#"}
	actualTrees := CountTrees(rows, 2, 3)
	if actualTrees != 1 {
		t.Errorf("TreeCount, expected 1, got %d", actualTrees)
	}
}
