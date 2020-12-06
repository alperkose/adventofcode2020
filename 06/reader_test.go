package main

import (
	"strings"
	"testing"
)

func TestReadingSinglePersonQuestions(t *testing.T) {
	source := strings.NewReader("abc")

	yesCollection := ReadSource(source)

	if len(yesCollection) != 1 {
		t.Errorf("expecting single collection but got %d", len(yesCollection))
	}
}

func TestReadingSingleGroupQuestions(t *testing.T) {
	source := strings.NewReader("a\nb\ncd")

	yesCollection := ReadSource(source)

	if len(yesCollection) != 1 {
		t.Errorf("expecting single collection but got %d", len(yesCollection))
	}
}

func TestReadingTwoGroupsQuestions(t *testing.T) {
	source := strings.NewReader("abc\n\nzd\nf")

	yesCollection := ReadSource(source)

	if len(yesCollection) != 2 {
		t.Errorf("expecting single collection but got %d", len(yesCollection))
		t.FailNow()
	}

}
