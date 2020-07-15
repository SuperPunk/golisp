package cond

import (
	"golisp/exp/common"
	"golisp/exp/exps"
	_if "golisp/exp/if"
	"log"
)

// cond是派生表达式，可转化为嵌套的if
func IsCond(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "cond")
}

func Clauses(exp *common.Pair) *common.Pair {
	return common.Cdr(exp).(*common.Pair)
}

func IsCondElseClause(clause *common.Pair) bool {
	return Predicate(clause) == "else"
}

func Predicate(clause *common.Pair) string {
	return common.Car(clause).(string)
}

func Actions(clause *common.Pair) *common.Pair {
	return common.Cdr(clause).(*common.Pair)
}

func ToIf(exp *common.Pair) interface{} {
	return ExpandClauses(Clauses(exp))
}

func ExpandClauses(clauses *common.Pair) interface{} {
	if clauses == nil {
		return nil
	}
	return func(first, rest *common.Pair) interface{} {
		if IsCondElseClause(first) {
			if rest == nil {
				return exps.Sequence2Exp(Actions(first))
			}
			log.Fatal("ELSE clause isn't last: ToIf", clauses)
		}
		return _if.MakeIf(Predicate(first), exps.Sequence2Exp(Actions(first)), ExpandClauses(rest))
	}(common.Car(clauses).(*common.Pair), common.Cdr(clauses).(*common.Pair))
}
