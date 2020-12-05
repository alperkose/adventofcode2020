package main

import "strconv"

func ValidateBirthYear(inp string) bool {
	year, err := strconv.Atoi(inp)
	if err != nil || year < 1920 || year > 2002 {
		return false
	}
	return true
}

func ValidateIssueYear(inp string) bool {
	year, err := strconv.Atoi(inp)
	if err != nil || year < 2010 || year > 2020 {
		return false
	}
	return true
}
func ValidateExpirationYear(inp string) bool {
	year, err := strconv.Atoi(inp)
	if err != nil || year < 2020 || year > 2030 {
		return false
	}
	return true
}

func ValidateHeight(inp string) bool {
	inpLn := len(inp)
	if inpLn < 2 {
		return false
	}
	unit := inp[inpLn-2:]
	value, err := strconv.Atoi(inp[0 : inpLn-2])
	if err != nil {
		return false
	}
	switch unit {
	case "in":
		if value >= 59 && value <= 76 {
			return true
		}
	case "cm":
		if value >= 150 && value <= 193 {
			return true
		}
	}
	return false
}

func ValidateHairColor(inp string) bool {
	inpLn := len(inp)
	if inpLn != 7 {
		return false
	}
	if inp[0] != '#' {
		return false
	}
	for i := 1; i < inpLn; i++ {
		if (inp[i] < 'a' || inp[i] > 'f') && (inp[i] < '0' || inp[i] > '9') {
			return false
		}
	}
	return true
}

func ValidateEyeColor(inp string) bool {
	switch inp {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func ValidatePID(inp string) bool {
	inpLn := len(inp)

	if inpLn != 9 {
		return false
	}
	for i := 0; i < inpLn; i++ {
		if inp[i] < '0' || inp[i] > '9' {
			return false
		}
	}
	return true
}
