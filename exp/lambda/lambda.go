package lambda

import "golisp/exp/common"

//  (lambda (⟨param⟩ … ⟨param⟩) ⟨body⟩))

func IsLambda(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "lambda")
}

func Parameters(exp *common.Pair) *common.Pair {
	return common.Cadr(exp).(*common.Pair)
}

func Body(exp *common.Pair) interface{} {
	return common.Cddr(exp)
}

func MakeLambda(parameters *common.Pair, body interface{}) *common.Pair {
	return common.Cons("lambda", common.Cons(parameters, body))
}
