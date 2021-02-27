package utils

import (
	"log"
	"strings"
)

/*
todo : checker with regex
*/
func GoalInFacts(goal string, facts []string) bool {
	for _, v := range facts {
		if v == goal {
			return true
		}
	}
	return false
}
func GoalInConclusions(cs []string, goal string) bool {
	counter := 0
	for i := 0; i < len(cs); i++ {
		if strings.Contains(strings.Join(cs, ""), goal) {
			return true

		} else {
			counter++
		}
	}
	if counter == len(cs) {
		return false
	}
	return true
}
func ExpParse(exp string) []string {
	exp = strings.ReplaceAll(exp, "ET", "")
	exp = strings.ReplaceAll(exp, "ALORS", ".")
	exp = strings.ReplaceAll(exp, " ", "")
	vars := strings.Split(exp, ".")
	return vars
}
func HandleErrFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func StToSlice(s string) []string {
	ss := make([]string, 0)
	s = strings.ReplaceAll(s, " ", "")
	for _, sss := range s {
		ss = append(ss, string(sss))
	}
	return ss
}
