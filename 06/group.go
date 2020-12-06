package main

import "strings"

type Group struct {
	answers []string
}

func NewGroup() *Group {
	return &Group{[]string{}}
}

func (g *Group) Add(answers string) *Group {
	g.answers = append(g.answers, answers)
	return g
}

func (g *Group) CountAnswers() int {
	lenAns := len(g.answers)
	if lenAns == 0 {
		return 0
	}
	uniqueQuestions := map[rune]int{}
	for i := 0; i < lenAns; i++ {
		answerGroup := g.answers[i]
		for _, q := range answerGroup {
			uniqueQuestions[q] = 1
		}
	}
	return len(uniqueQuestions)

}

func (g *Group) CommonAnswers() int {
	lenAns := len(g.answers)
	if lenAns == 0 {
		return 0
	}
	uniqueQuestions := map[rune]int{}
	firstAnswerGroup := g.answers[0]
	for _, ans := range firstAnswerGroup {
		uniqueQuestions[ans] = 1
	}

	for i := 1; i < lenAns; i++ {
		answerGroup := g.answers[i]
		for key, _ := range uniqueQuestions {
			if strings.Index(answerGroup, string(key)) < 0 {
				delete(uniqueQuestions, key)
			}
		}
	}
	return len(uniqueQuestions)
}
