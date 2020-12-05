package main

import "testing"

func TestPassportValidity(t *testing.T) {
	var cases = []struct {
		testName, input  string
		expectedValidity bool
	}{
		{"Valid pass", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"Birth year missing", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd iyr:2017 cid:147 hgt:183cm", false},
		{"Issue year missing", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 cid:147 hgt:183cm", false},
		{"Expiration year missing", "ecl:gry pid:860033327 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", false},
		{"Height missing", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147", false},
		{"Hair color missing", "ecl:gry pid:860033327 eyr:2020 byr:1937 iyr:2017 cid:147 hgt:183cm", false},
		{"Eye color missing", "pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", false},
		{"Passport ID missing", "ecl:gry eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", false},
		{"Country ID missing", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 hgt:183cm", true},
		{"Fields are separated by space", "ecl:grypid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", false},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			pass := PassportFromString(c.input)
			if pass.Valid() != c.expectedValidity {
				t.Errorf("Expected %v, got %v", validityStr(c.expectedValidity), validityStr(pass.Valid()))
			}
		})
	}
}

func validityStr(valid bool) string {
	if valid {
		return "Valid"
	}
	return "Invalid"
}
