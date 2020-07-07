package _if

import "golisp/exp/common"

func If(exp []string) bool {
	return common.TaggedList(exp, "if")
}

func IfPredicate(exp []string) int {
	return exp[1]
}

func IfConsequent(exp []string) int {
	return exp[2]
}

func IfAlternative(exp []string) int {
	if len(exp) > 3 {
		return exp[3]
	}
	return false
}

func MakeIf(predicate string, consequent string, alternative string) []string {
	return []string{"if", predicate, consequent, alternative}
}

func True(expValue int) bool {
	return true
}
