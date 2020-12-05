package main

import (
	"strings"
	"testing"
)

func TestReadingSinglePassword(t *testing.T) {
	source := strings.NewReader("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm")

	passports := ReadSource(source, PassportFromString)

	if len(passports) != 1 {
		t.Errorf("expecting single passport but got %d", len(passports))
	} else if !passports[0].Valid() {
		t.Errorf("unable to parse passport")
	}
}

func TestReadingSinglePasswordInTwoLines(t *testing.T) {
	source := strings.NewReader("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm")

	passports := ReadSource(source, PassportFromString)

	if len(passports) != 1 {
		t.Errorf("expecting single passport but got %d", len(passports))
	} else if !passports[0].Valid() {
		t.Errorf("unable to parse passport")
	}

}
func TestReadingMultiplePasswords(t *testing.T) {
	source := strings.NewReader("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\nhcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm")

	passports := ReadSource(source, PassportFromString)

	if len(passports) != 2 {
		t.Errorf("expecting single passport but got %d", len(passports))
		t.FailNow()
	}
	if !passports[0].Valid() {
		t.Errorf("unable to parse passport")
	}
	if !passports[1].Valid() {
		t.Errorf("unable to parse passport")
	}

}
