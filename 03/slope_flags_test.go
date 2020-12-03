package main

import "testing"

func TestSlopeInputWithoutComma(t *testing.T) {
	sf := make(slopeFlags, 0)

	err := sf.Set("12")
	if err == nil {
		t.Error("expected an error")
	}
}
func TestSlopeInputWithFirstPartNonNumeric(t *testing.T) {
	sf := make(slopeFlags, 0)

	err := sf.Set("asdf,42")
	if err == nil {
		t.Error("expected an error")
	}
}
func TestSlopeInputWithSecondPartNonNumeric(t *testing.T) {
	sf := make(slopeFlags, 0)

	err := sf.Set("42,asdf")
	if err == nil {
		t.Error("expected an error")
	}
}

func TestSlopeInputWithMoreThanTwoIntegers(t *testing.T) {
	sf := make(slopeFlags, 0)

	err := sf.Set("42,42,42")
	if err == nil {
		t.Error("expected an error")
	}
}
func TestSlopeInputReturnsSlope(t *testing.T) {
	sf := make(slopeFlags, 0)

	err := sf.Set("4,2")
	if err != nil {
		t.Errorf("Not expected an error, %v", err)
	}

	if sf[0].Down != 4 && sf[0].Right != 2 {
		t.Errorf("Expecting %v, got %v", Slope{4, 2}, sf[0])
	}
}
func TestSlopeMultipleInputsAddedToArray(t *testing.T) {
	sf := make(slopeFlags, 0)

	err := sf.Set("4,2")
	if err != nil {
		t.Errorf("Not expected an error, %v", err)
	}

	if sf[0].Down != 4 && sf[0].Right != 2 {
		t.Errorf("Expecting %v, got %v", Slope{4, 2}, sf[0])
	}

	err = sf.Set("1,3")
	if err != nil {
		t.Errorf("Not expected an error, %v", err)
	}

	if sf[1].Down != 1 && sf[1].Right != 3 {
		t.Errorf("Expecting %v, got %v", Slope{4, 2}, sf[0])
	}

}
