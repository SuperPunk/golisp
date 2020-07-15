package variable

import "golisp/exp/common"

func Variable(exp interface{}) bool {
	return common.IsSymbol(exp)
}
