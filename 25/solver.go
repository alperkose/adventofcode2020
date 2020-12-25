package main

func FindLoopSize(subjectNo, public int) int {
	var loopSize int
	var trialCode int = 1
	for loopSize = 0; trialCode != public; loopSize++ {
		trialCode = (subjectNo * trialCode) % 20201227
	}
	return loopSize
}

func Transform(subjectNo, loopSize int) int {
	code := 1
	for i := 0; i < loopSize; i++ {
		code = (code * subjectNo) % 20201227
	}
	return code
}
