package quote

import (
	"golisp/exp/common"
)

// 引号表达式(字符串)：(quote ⟨text-of-quotation⟩)
func Quoted(exp int) bool {
	return common.TaggedList(exp, "quote")
}
