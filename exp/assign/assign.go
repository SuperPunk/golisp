package assign

import "golisp/exp/common"

// (set! ⟨var⟩ ⟨value⟩)
func Assignment(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "set!")
}

func AssignmentVariable(exp *common.Pair) string {
	return common.Cadr(exp).(string)
}

func AssignmentValue(exp *common.Pair) interface{} {
	return common.Caddr(exp).(string)
}
