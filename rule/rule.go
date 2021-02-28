package rule

import (
	"bytes"
	"fmt"
	"inference_engine_proof_of_work/utils"
	"io/ioutil"
)

type Rule struct {
	Name       string
	IsActive   bool
	Vars       []string
	Conclusion []string
}

func (r *Rule) NewRule(name string, vars, conclusion []string) *Rule {
	return &Rule{
		Name:       name,
		IsActive:   true,
		Vars:       vars,
		Conclusion: conclusion,
	}
}
func (r *Rule) Print() string {
	return fmt.Sprintf(" \t Rule Name : %s Rule Vars : %s Rule Conclusion : %s \n", r.Name, r.Vars, r.Conclusion)
}
func (r *Rule) DisActivate() {
	r.IsActive = false
}

func (r *Rule) RuleState() bool {
	return r.IsActive

}
func (r *Rule) RuleInFacts(facts []string) bool {
	found := 0
	for i := 0; i < len(facts); i++ {

		for j := 0; j < len(r.Vars); j++ {
			if facts[i] == r.Vars[j] {
				found++
			}
		}
	}
	if found == len(r.Vars) {
		return true
	} else {
		return false
	}

}
func CheckGoalInConclusions(rules []Rule, goal string) bool {
	ctr := 0
	for i := 0; i < len(rules); i++ {
		if utils.GoalInConclusions(rules[i].Conclusion, goal) {
			return true
		} else {
			ctr++
		}
	}
	if ctr == len(rules) {
		return false
	}
	return true
}

func GetRuleWithThisGoal(goal string, rules []Rule) (*Rule, bool) {
	for _, r := range rules {
		for i := 0; i < len(r.Conclusion); i++ {
			if r.Conclusion[i] == goal {
				return &r, true
			}
		}
	}
	return nil, false
}

func FileToRules(fileName string) *[]Rule {
	var rs []Rule
	var r Rule
	f, err := ioutil.ReadFile(fileName)
	utils.HandleErrFatal(err)
	for _, line := range bytes.Split(f, []byte("\n")) {
		lineString := string(line)
		ruleName := lineString[:3]
		exp := lineString[7:]
		vars := utils.ExpParse(exp)
		r := r.NewRule(ruleName, utils.StToSlice(vars[0]), utils.StToSlice(vars[1]))
		rs = append(rs, *r)
	}
	return &rs
}
