package assign

import "golisp/exp/common"

// (set! ⟨var⟩ ⟨value⟩)
func Assignment(exp []string) bool {
	return common.TaggedList(exp, "set!")
}

func AssignmentVariable(exp []string) int {
	return exp[1]
}

func AssignmentValue(exp []string) int {
	return exp[2]
}
