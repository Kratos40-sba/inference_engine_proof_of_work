package engine

import (
	"inference_engine_proof_of_work/rule"
	"inference_engine_proof_of_work/utils"
)

func Forward(rules []rule.Rule, facts []string, goal string) ([]string, bool) {
	found := false
	noRulesCanBeApplied := false
	rulesDeactivated := 0
	ctr := 0
	for i := 0; i < len(rules); i++ {
		if utils.GoalInConclusions(rules[i].Conclusion, goal) {
			break
		} else {
			ctr++
		}
	}
	if ctr == len(rules) {
		return facts, false
	}
	for !found && !noRulesCanBeApplied {
		for i := 0; i < len(rules); i++ {
			isActive := rules[i].RuleState()
			if rules[i].RuleInFacts(facts) && isActive {
				facts = append(facts, rules[i].Conclusion...)
				rules[i].DisActivate()
				rulesDeactivated++
				if utils.GoalInFacts(goal, facts) {
					found = true
					break
				}
			}
			if rulesDeactivated == len(rules) {
				noRulesCanBeApplied = true
				break
			}
		}
	}
	return facts, found
}
