package _if

import "golisp/exp/common"

// (if <predicate> <consequent> <alternative>)
func If(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "if")
}

func IfPredicate(exp *common.Pair) interface{} {
	return common.Cadr(exp)
}

func IfConsequent(exp *common.Pair) interface{} {
	return common.Caddr(exp)
}

func IfAlternative(exp *common.Pair) interface{} {
	if common.Cdddr(exp) != nil {
		return common.Cadddr(exp)
	}
	return false
}

func MakeIf(predicate interface{}, consequent interface{}, alternative interface{}) *common.Pair {
	return common.List("if", predicate, consequent, alternative)
}

func True(x interface{}) bool {
	v, ok := x.(bool)
	return !ok || v == true
}
