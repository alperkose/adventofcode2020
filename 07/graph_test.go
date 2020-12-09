package main

import "testing"

func TestAddingABagContainingNoOtherBagsToInitialRules(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("faded blue"))

	bags := rules.BagsCanContain("faded blue")
	if len(bags) != 0 {
		t.Errorf("expected no bags that can contain %#+v but got %#+v \n", "faded blue", bags)
	}
}
func TestAddingABagContainingTargetBag(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("vibrant plum", Content{"faded blue", 5}, Content{"dotted black", 6}))

	bags := rules.BagsCanContain("faded blue")
	if len(bags) != 1 {
		t.Errorf("expected one bag that can contain %#+v but got %#+v \n", "faded blue", bags)
		t.FailNow()
	}
	expected := []string{"vibrant plum"}
	oneOf(t, bags[0], expected)
}
func TestAddingBagsContainingTargetBagAndItsContainer(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("vibrant plum", Content{"faded blue", 5}, Content{"dotted black", 6}))

	rules.Add(NewBag("shiny gold", Content{"dark olive", 1}, Content{"vibrant plum", 2}))

	bags := rules.BagsCanContain("faded blue")
	if len(bags) != 2 {
		t.Errorf("expected two bags that can contain %#+v but got %#+v \n", "faded blue", bags)
		t.FailNow()
	}
	expected := []string{"vibrant plum", "shiny gold"}
	oneOf(t, bags[0], expected)
	oneOf(t, bags[1], expected)

	if bags[0] == bags[1] {
		t.Errorf("expected %#+v and  %#+v to contain %#+v but got %#+v \n", "vibrant plum", "shiny gold", "faded blue", bags)
	}

}

func TestAddingBagsContainingTargetBagAndTwoItsContainer(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("vibrant plum", Content{"faded blue", 5}, Content{"dotted black", 6}))

	rules.Add(NewBag("shiny gold", Content{"dark olive", 1}, Content{"vibrant plum", 2}))

	rules.Add(NewBag("dark olive", Content{"faded blue", 3}, Content{"dotted black", 4}))

	bags := rules.BagsCanContain("faded blue")
	if len(bags) != 3 {
		t.Errorf("expected two bags that can contain %#+v but got %#+v \n", "faded blue", bags)
		t.FailNow()
	}
	expected := []string{"vibrant plum", "dark olive", "shiny gold"}
	oneOf(t, bags[0], expected)
	oneOf(t, bags[1], expected)
	oneOf(t, bags[2], expected)
	if bags[0] == bags[1] || bags[1] == bags[2] || bags[2] == bags[0] {
		t.Errorf("expected a unique list of %#+v but got %#+v \n", expected, bags)
	}

}

func TestAddingBagsContainingTargetBagAndTwoItsContainerAndAnUnrelatedContainer(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("vibrant plum", Content{"faded blue", 5}, Content{"dotted black", 6}))

	rules.Add(NewBag("shiny gold", Content{"dark olive", 1}, Content{"vibrant plum", 2}))

	rules.Add(NewBag("dark olive", Content{"faded blue", 3}, Content{"dotted black", 4}))
	rules.Add(NewBag("dotted black"))

	bags := rules.BagsCanContain("faded blue")
	if len(bags) != 3 {
		t.Errorf("expected two bags that can contain %#+v but got %#+v \n", "faded blue", bags)
		t.FailNow()
	}
	expected := []string{"vibrant plum", "dark olive", "shiny gold"}
	oneOf(t, bags[0], expected)
	oneOf(t, bags[1], expected)
	oneOf(t, bags[2], expected)
	if bags[0] == bags[1] || bags[1] == bags[2] || bags[2] == bags[0] {
		t.Errorf("expected a unique list of %#+v but got %#+v \n", expected, bags)
	}

}

func TestCountBagsWithNoContentIsZero(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("faded blue"))

	cnt := rules.CountBagsWithin("faded blue")
	if cnt != 0 {
		t.Errorf("Expected 0 got %d", cnt)
	}
}

func TestCountBagsContainingABag(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("vibrant plum", Content{"faded blue", 5}))

	cnt := rules.CountBagsWithin("vibrant plum")
	if cnt != 5 {
		t.Errorf("Expected 5 got %d", cnt)
	}
}
func TestCountBagsWithTwoDifferentBags(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("vibrant plum", Content{"faded blue", 5}, Content{"dotted black", 6}))

	cnt := rules.CountBagsWithin("vibrant plum")
	if cnt != 11 {
		t.Errorf("Expected 11 got %d", cnt)
	}
}

func TestCountingBagsContainingALotOfBags(t *testing.T) {
	rules := NewRules()
	rules.Add(NewBag("vibrant plum", Content{"faded blue", 5}, Content{"dotted black", 6}))

	rules.Add(NewBag("shiny gold", Content{"dark olive", 1}, Content{"vibrant plum", 2}))

	rules.Add(NewBag("dark olive", Content{"faded blue", 3}, Content{"dotted black", 4}))
	rules.Add(NewBag("dotted black"))

	cnt := rules.CountBagsWithin("shiny gold")
	if cnt != 32 {
		t.Errorf("Expected 32 got %d", cnt)
	}
}

func oneOf(t *testing.T, v string, of []string) {
	for _, o := range of {
		if v == o {
			return
		}
	}
	t.Errorf("expected %#+v to be one of %#+v\n", v, of)

}
