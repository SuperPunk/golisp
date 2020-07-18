package exps

import (
	"golisp/exp/begin"
	"golisp/exp/common"
)

// 这样判断 common.Cdr(exps) == nil有坑
func IsLastExp(exps *common.Pair) bool {
	return common.IsNull(common.Cdr(exps))
}

func FirstExp(exps *common.Pair) interface{} {
	return common.Car(exps)
}

func RestExps(exps *common.Pair) *common.Pair {
	return common.Cdr(exps).(*common.Pair)
}

// 一系列表达式转化为begin表达式
func Sequence2Exp(seq *common.Pair) interface{} {
	if seq == nil {
		return seq
	}
	if IsLastExp(seq) {
		return FirstExp(seq)
	}
	return begin.MakeBegin(seq)
}
