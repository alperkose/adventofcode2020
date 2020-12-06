package main

import "testing"

func TestCountingSingleAnswers(t *testing.T) {
	counts := NewGroup().Add("abc").CountAnswers()
	if counts != 3 {
		t.Errorf("Count should be %d but got %d", 3, counts)
	}
}
func TestCountingRepeatingAnswers(t *testing.T) {
	counts := NewGroup().Add("abc").Add("a").CountAnswers()
	if counts != 3 {
		t.Errorf("Count should be %d but got %d", 3, counts)
	}
}
func TestCountingNonRepeatingAnswers(t *testing.T) {
	counts := NewGroup().Add("abc").Add("dfz").CountAnswers()
	if counts != 6 {
		t.Errorf("Count should be %d but got %d", 6, counts)
	}
}
func TestCountingRepeatingAndNonRepeatingAnswers(t *testing.T) {
	counts := NewGroup().Add("abc").Add("ad").CountAnswers()
	if counts != 4 {
		t.Errorf("Count should be %d but got %d", 4, counts)
	}
}

func TestCommonAnswersWithASinglePerson(t *testing.T) {
	count := NewGroup().Add("abc").CommonAnswers()
	if count != 3 {
		t.Errorf("expecting 3 got %#+v\n", count)
	}
}
func TestCommonAnswersWithATwoPeople(t *testing.T) {
	count := NewGroup().Add("abc").Add("ab").CommonAnswers()
	if count != 2 {
		t.Errorf("expecting 2 got %#+v\n", count)
	}
}
func TestCommonAnswersWithATwoPeopleAndNoCommonAnswer(t *testing.T) {
	count := NewGroup().Add("abc").Add("zdf").CommonAnswers()
	if count != 0 {
		t.Errorf("expecting 0 got %#+v\n", count)
	}
}
