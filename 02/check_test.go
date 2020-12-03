package main

import (
	"fmt"
	"testing"
)

func TestInvalidInputToCheckWith(t *testing.T) {

	ok := CheckWith("messed up input", always(true))
	if ok {
		t.Error("should not accept messed up input")
	}
}
func TestCheckWithCheckerReturningTrue(t *testing.T) {
	ok := CheckWith("1-3 a: abcde", always(true))

	if !ok {
		t.Error("should have passed the check")
	}
}
func TestCheckWithCheckerReturningFalse(t *testing.T) {
	ok := CheckWith("1-3 a: abcde", always(false))
	if ok {
		t.Error("should not have passed the check")
	}

}
func always(b bool) Checker {
	return func(i, j int, c, p string) bool { return b }
}

var testCases = []struct {
	input     string
	first     int
	second    int
	checkChar string
	pwd       string
	ok        bool
}{
	{"1-3 a: abcde", 1, 3, "a", "abcde", true},
	{"1-3 b: cdefg", 1, 3, "b", "cdefg", true},
	{"2-9 c: ccccccccc", 2, 9, "c", "ccccccccc", true},
	{"asdf-9 f: dfas", 0, 0, "", "", false},
	{"fasd", 0, 0, "", "", false},
	{"1-2 fasd", 0, 0, "", "", false},
	{"1-2", 0, 0, "", "", false},
	{"1- fdsa: Fdas", 0, 0, "", "", false},
	{"1 fdsa: Fdas", 0, 0, "", "", false},
	{"1-2: Fdas", 0, 0, "", "", false},
}

func TestParsingPolicies(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			atLeast, atMost, checkChar, pwd, err := ParsePolicy(tc.input)
			if (err == nil) != tc.ok {
				t.Errorf("got error %v", err)
			}
			if atLeast != tc.first {
				t.Errorf("atLeast, got %v, want %v", atLeast, tc.first)
			}
			if atMost != tc.second {
				t.Errorf("atMost, got %v, want %v", atMost, tc.second)
			}
			if checkChar != tc.checkChar {
				t.Errorf("checkChar, got %v, want %v", checkChar, tc.checkChar)
			}
			if pwd != tc.pwd {
				t.Errorf("pwd, got %v, want %v", pwd, tc.pwd)
			}
		})
	}
}

var charCountTestCases = []struct {
	first     int
	second    int
	checkChar string
	pwd       string
	ok        bool
}{
	{1, 3, "a", "abcde", true},
	{1, 3, "b", "cdefg", false},
	{2, 9, "c", "ccccccccc", true},
}

func TestCharCount(t *testing.T) {
	for i, tc := range charCountTestCases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ok := CharCount(tc.first, tc.second, tc.checkChar, tc.pwd)
			if ok != tc.ok {
				t.Errorf("got %v, want %v", ok, tc.ok)
			}
		})
	}
}

var exactMatchTestCases = []struct {
	first     int
	second    int
	checkChar string
	pwd       string
	ok        bool
}{
	{1, 3, "a", "abcde", true},
	{1, 3, "b", "cdefg", false},
	{2, 9, "c", "ccccccccc", false},
}

func TestExactMatch(t *testing.T) {
	for i, tc := range exactMatchTestCases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ok := ExactlyOneMatchedIndex(tc.first, tc.second, tc.checkChar, tc.pwd)
			if ok != tc.ok {
				t.Errorf("got %v, want %v", ok, tc.ok)
			}
		})
	}
}
