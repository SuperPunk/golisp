package quote

import (
	"golisp/exp/common"
)

// 引号表达式(字符串)：(quote ⟨text-of-quotation⟩)
func IsQuoted(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "quote")
}

func TextOfQuotation(exp *common.Pair) string {
	return common.Cadr(exp).(string)
}
