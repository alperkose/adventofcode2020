package main

import (
	"fmt"
	"strings"
)

type Passport interface {
	Valid() bool
}

type ValidateField func(string) bool

// var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var requiredFields = map[string]ValidateField{
	"byr": ValidateBirthYear,
	"iyr": ValidateIssueYear,
	"eyr": ValidateExpirationYear,
	"hgt": ValidateHeight,
	"hcl": ValidateHairColor,
	"ecl": ValidateEyeColor,
	"pid": ValidatePID,
}

func PassportFromString(rawPassport string) Passport {
	for field, validateField := range requiredFields {
		fieldIndex := strings.Index(rawPassport, fmt.Sprintf("%s:", field))
		if fieldIndex < 0 {
			return passport{false}
		} else if fieldIndex != 0 && rawPassport[fieldIndex-1] != ' ' {
			return passport{false}
		}
		fieldValueBegin := fieldIndex + 4
		spaceIndex := strings.Index(rawPassport[fieldValueBegin:], " ")
		var fieldValue string
		if spaceIndex == -1 {
			fieldValue = rawPassport[fieldValueBegin:]
		} else {
			fieldValue = rawPassport[fieldValueBegin : fieldValueBegin+spaceIndex]
		}
		if !validateField(fieldValue) {
			return passport{false}
		}
	}
	return passport{true}
}

type passport struct {
	valid bool
}

func (p passport) Valid() bool {
	return p.valid
}
