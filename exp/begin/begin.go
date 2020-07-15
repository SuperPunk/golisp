package begin

import (
	"golisp/exp/common"
)

func IsBegin(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "begin")
}

func Actions(exp *common.Pair) *common.Pair {
	return common.Cdr(exp).(*common.Pair)
}

func MakeBegin(seq *common.Pair) *common.Pair {
	return common.Cons("begin", seq)
}
