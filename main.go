package main

import (
	"fmt"
	"inference_engine_proof_of_work/engine"
	"inference_engine_proof_of_work/rule"
)

func main() {

	rs := rule.FileToRules("rules.txt")
	facts := []string{"A", "B", "C", "D", "E"}
	goal := "Z"

	newFacts, found := engine.Forward(*rs, facts, goal)
	newFacts2, found2 := engine.BackWard(*rs, facts, goal)

	fmt.Println("Forward Channing :")
	fmt.Println(newFacts, found)
	fmt.Println("Backward Channing :")
	fmt.Println(newFacts2, found2)

}
