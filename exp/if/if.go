package _if

import "golisp/exp/common"

func If(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "if")
}

func IfPredicate(exp *common.Pair) string {
	return common.Cadr(exp).(string)
}

func IfConsequent(exp *common.Pair) interface{} {
	return common.Caddr(exp)
}

func IfAlternative(exp *common.Pair) interface{} {
	if common.Cdddr(exp) != nil {
		return common.Cdddr(exp)
	}
	return "false"
}

func MakeIf(predicate interface{}, consequent interface{}, alternative interface{}) *common.Pair {
	return common.List("if", predicate, consequent, alternative)
}

func True(x interface{}) bool {
	sv, ok := x.(string)
	return !(ok && sv == "false")
}
