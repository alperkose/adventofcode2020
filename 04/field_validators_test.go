package main

import "testing"

func TestValidateBirthYear(t *testing.T) {
	if !ValidateBirthYear("1920") {
		t.Error("1920")
		t.FailNow()
	}
	if !ValidateBirthYear("1930") {
		t.Error("1930")
		t.FailNow()
	}
	if !ValidateBirthYear("2002") {
		t.Error("2002")
		t.FailNow()
	}
	if ValidateBirthYear("2003") {
		t.Error("2003")
		t.FailNow()
	}
	if ValidateBirthYear("2020") {
		t.Error("2020")
		t.FailNow()
	}
	if ValidateBirthYear("1919") {
		t.Error("1919")
		t.FailNow()
	}
}
func TestValidateIssueYear(t *testing.T) {
	if !ValidateIssueYear("2010") {
		t.Error("2010")
		t.FailNow()
	}
	if !ValidateIssueYear("2015") {
		t.Error("2015")
		t.FailNow()
	}
	if !ValidateIssueYear("2020") {
		t.Error("2020")
		t.FailNow()
	}
	if ValidateIssueYear("2009") {
		t.Error("2009")
		t.FailNow()
	}
	if ValidateIssueYear("2021") {
		t.Error("2021")
		t.FailNow()
	}
}
func TestValidateExpirationYear(t *testing.T) {
	if !ValidateExpirationYear("2020") {
		t.Error("2020")
		t.FailNow()
	}
	if !ValidateExpirationYear("2025") {
		t.Error("2025")
		t.FailNow()
	}
	if !ValidateExpirationYear("2030") {
		t.Error("2030")
		t.FailNow()
	}
	if ValidateExpirationYear("2019") {
		t.Error("2019")
		t.FailNow()
	}
	if ValidateExpirationYear("2031") {
		t.Error("2031")
		t.FailNow()
	}
}
func TestValidateHeight(t *testing.T) {
	if !ValidateHeight("150cm") {
		t.Error("150cm")
		t.FailNow()
	}
	if !ValidateHeight("193cm") {
		t.Error("193cm")
		t.FailNow()
	}
	if ValidateHeight("149cm") {
		t.Error("149cm")
		t.FailNow()
	}
	if ValidateHeight("194cm") {
		t.Error("194cm")
		t.FailNow()
	}

	if !ValidateHeight("59in") {
		t.Error("59in")
		t.FailNow()
	}
	if !ValidateHeight("76in") {
		t.Error("76in")
		t.FailNow()
	}
	if ValidateHeight("58in") {
		t.Error("58in")
		t.FailNow()
	}
	if ValidateHeight("77in") {
		t.Error("77in")
		t.FailNow()
	}

	if ValidateHeight("77ok") {
		t.Error("77ok")
		t.FailNow()
	}
	if ValidateHeight("sdfcm") {
		t.Error("sdfcm")
		t.FailNow()
	}
	if ValidateHeight("fdsain") {
		t.Error("fdsain")
		t.FailNow()
	}
	if ValidateHeight("") {
		t.Error("")
		t.FailNow()
	}

}
func TestValidateHairColor(t *testing.T) {
	if !ValidateHairColor("#abf059") {
		t.Error("#abf059")
		t.FailNow()
	}
	if !ValidateHairColor("#ffffff") {
		t.Error("#ffffff")
		t.FailNow()
	}
	if !ValidateHairColor("#000000") {
		t.Error("#000000")
		t.FailNow()
	}
	if ValidateHairColor("#adz432") {
		t.Error("#adz432")
		t.FailNow()
	}

	if ValidateHairColor("#abc1234") {
		t.Error("#abc1234")
		t.FailNow()
	}
	if ValidateHairColor("#abc12") {
		t.Error("#abc12")
		t.FailNow()
	}
	if ValidateHairColor("123fdc") {
		t.Error("123fdc")
		t.FailNow()
	}
	if ValidateHairColor("1234fdc") {
		t.Error("123fdc")
		t.FailNow()
	}
	if ValidateHairColor("") {
		t.Error("")
		t.FailNow()
	}

}
func TestValidateEyeColor(t *testing.T) {
	if !ValidateEyeColor("amb") {
		t.Error("amb")
		t.FailNow()
	}
	if !ValidateEyeColor("blu") {
		t.Error("blu")
		t.FailNow()
	}
	if !ValidateEyeColor("brn") {
		t.Error("brn")
		t.FailNow()
	}
	if !ValidateEyeColor("gry") {
		t.Error("gry")
		t.FailNow()
	}
	if !ValidateEyeColor("grn") {
		t.Error("grn")
		t.FailNow()
	}
	if !ValidateEyeColor("hzl") {
		t.Error("hzl")
		t.FailNow()
	}
	if !ValidateEyeColor("oth") {
		t.Error("oth")
		t.FailNow()
	}

	if ValidateEyeColor("asm") {
		t.Error("asm")
		t.FailNow()
	}

	if ValidateEyeColor("blue") {
		t.Error("blue")
		t.FailNow()
	}

	if ValidateEyeColor("") {
		t.Error("")
		t.FailNow()
	}

}
func TestValidatePID(t *testing.T) {
	if !ValidatePID("012345678") {
		t.Error("012345678")
		t.FailNow()
	}
	if !ValidatePID("000000001") {
		t.Error("000000001")
		t.FailNow()
	}
	if !ValidatePID("999999999") {
		t.Error("999999999")
		t.FailNow()
	}
	if ValidatePID("0123456789") {
		t.Error("0123456789")
		t.FailNow()
	}
	if ValidatePID("O12345678") {
		t.Error("O12345678")
		t.FailNow()
	}
	if ValidatePID("") {
		t.Error("")
		t.FailNow()
	}
}
