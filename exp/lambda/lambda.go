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

func Body(exp *common.Pair) string {
	return common.Cddr(exp).(string)
}

func MakeLambda(parameters *common.Pair, body string) *common.Pair {
	return common.Cons("lambda", common.Cons(parameters, body))
}
