package main

import (
	"strings"
	"testing"
)

func TestSingleLine(t *testing.T) {
	source := strings.NewReader("faded blue bags contain no other bags.")

	rules := ReadSource(source)

	if len(rules.bags) != 1 {
		t.Errorf("expecting single bag but got %#+v", rules.bags)
	}
	if len(rules.containedIn) != 0 {
		t.Errorf("expected an empty map, but got %#+v", rules.containedIn)
	}
}
func TestTwoLines(t *testing.T) {
	source := strings.NewReader("vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.")

	rules := ReadSource(source)

	if len(rules.bags) != 2 {
		t.Errorf("expecting 2 bags but got %#+v", rules.bags)
	}
	if len(rules.containedIn) != 2 {
		t.Errorf("expected an map with two entries, but got %#+v", rules.containedIn)
	}
}
