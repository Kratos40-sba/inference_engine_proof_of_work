package engine

import (
	"container/list"
	"inference_engine_proof_of_work/rule"
	stack2 "inference_engine_proof_of_work/stack"
	"inference_engine_proof_of_work/utils"
)

func BackWard(rules []rule.Rule, facts []string, goal string) ([]string, bool) {

	ctr := 0
	var stack = stack2.MyStack{Stack: list.New()}
	for i := 0; i < len(rules); i++ {
		if utils.GoalInConclusions(rules[i].Conclusion, goal) {
			break
		} else {
			ctr++
		}
	}
	if ctr == len(rules) {
		return facts, false
	} else {
		stack.Push(goal)
	}
	for stack.Size() > 0 {
		if utils.GoalInFacts(goal, facts) {
			goal = stack.Pop()
		} else if rule.CheckGoalInConclusions(rules, goal) {

			r, found := rule.GetRuleWithThisGoal(goal, rules)
			if !found {
				return facts, false
			} else {
				for i := 0; i < len(r.Vars); i++ {
					stack.Push(r.Vars[i])
				}
				facts = append(facts, goal)
			}
		} else {
			return facts, false
		}
	}
	return facts, true
}
