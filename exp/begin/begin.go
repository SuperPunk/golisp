package begin

import "golisp/exp/common"

func IsBegin(exp []string) bool {
	return common.TaggedList(exp, "begin")
}

func BeginActions(exp []string) int {
	return exp[1]
}

func IsLastExp(seq []string) bool {
	return seq[1] == nil
}

func FirstExp(seq []string) string {
	return seq[0]
}

func RestExps(seq []string) []string {
	return seq[1]
}

func MakeBegin(seq []string) []string {
	return []string{"begin", seq[1]}
}

func Sequence2Exp(seq []string) []string {
	if seq == nil {
		return seq
	}
	if IsLastExp(seq) {
		return FirstExp(seq)
	}
	return MakeBegin(seq)
}
