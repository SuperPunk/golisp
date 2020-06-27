package core

import (
	lexp "golisp/exp"
	"strconv"
)

func Eval(exp, env int) int {
	if lexp.IsSelfEvaluating(exp) {
		return exp
	} else if lexp.IsVariable(exp) {
		return lookupVariableValue(exp, env)
	} else if lexp.IsQuoted(exp) {
		return textOfQuotation(exp)
	} else if lexp.IsAssignment(exp) {
		return evalAssignment(exp, env)
	} else if lexp.IsDefinition(exp) {
		return evalDefinition(exp, env)
	} else if lexp.IsIf(exp) {
		return evalIf(exp, env)
	} else if lexp.IsLambda(exp) {
		return makeProcedure(lambdaParameters(exp), lambdaBody(exp), env)
	} else if lexp.IsBegin(exp) {
		return evalSequence(beginActions(exp), env)
	} else if lexp.IsCond(exp) {
		return Eval(cond2If(exp), env)
	} else if lexp.IsApplication(exp) {
		return apply(Eval(operator(exp), env), listOfValues(operands(exp), env))
	} else {
		panic("Unknown expression type -- EVAL " + strconv.Itoa(exp))
	}
}

func listOfValues(i int, env int) int {
	return 0
}

func operands(exp int) int {
	return 0
}

func operator(exp int) int {
	return 0
}

func cond2If(exp int) int {
	return 0
}

func evalSequence(actions int, env int) int {
	return 0
}

func beginActions(exp int) int {
	return 0
}

func lambdaBody(exp int) int {
	return 0
}

func lambdaParameters(exp int) int {
	return 0
}

func makeProcedure(parameters int, body int, env int) int {
	return 0
}

func evalIf(exp int, env int) int {
	return 0
}

func evalDefinition(exp int, env int) int {
	return 0
}

func evalAssignment(exp int, env int) int {
	return 0
}

func textOfQuotation(exp int) int {
	return 0
}

func lookupVariableValue(exp int, env int) int {
	return 0
}
