package cond

import (
	"golisp/exp/begin"
	"golisp/exp/common"
	_if "golisp/exp/if"
	"log"
)

// cond是派生表达式，可转化为嵌套的if
func IsCond(exp []string) bool {
	return common.TaggedList(exp, "cond")
}

func CondClauses(exp []string) string {
	return exp[1]
}

func IsCondElseClause(clause []string) bool {
	return CondPredicate(clause) == "else"
}

func CondPredicate(clause []string) string {
	return clause[0]
}

func CondActions(clause []string) string {
	return clause[1]
}

func Cond2If(exp []string) []string {
	return ExpandClauses(CondClauses(exp))
}

func ExpandClauses(clauses string) []string {
	if clauses == nil {
		return nil
	}
	return func(first []string, rest []string) []string {
		if IsCondElseClause(first) {
			if rest == nil {
				return begin.Sequence2Exp(CondActions(first))
			}
			log.Fatal("ELSE clause isn't last: Cond2If", clauses)
		}
		return _if.MakeIf(CondPredicate(first), begin.Sequence2Exp(CondActions(first)), ExpandClauses(rest))
	}(clauses[0], clauses[1])
}
