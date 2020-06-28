package core

import (
	lexp "golisp/exp"
	"golisp/exp/assign"
	"golisp/exp/define"
	"golisp/exp/exps"
	"golisp/exp/if"
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

// 生成实际参数表
func listOfValues(exps int, env int) int {
	if lexp.NoOperands(exps) {
		return 0
	}
	// todo 定义list结构
	return Eval(firstOperand(exps), env) + listOfValues(restOperands(exps), env)
}

func restOperands(exps int) int {
	return 0
}

func firstOperand(exps int) int {
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

func evalSequence(expressions int, env int) int {
	if exps.LastExp(expressions) {
		return Eval(exps.FirstExp(expressions), env)
	} else {
		Eval(exps.FirstExp(expressions), env)
		return evalSequence(exps.RestExps(expressions), env)
	}
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
	if _if.True(Eval(_if.Predicate(exp), env)) {
		return Eval(_if.Consequent(exp), env)
	}
	return Eval(_if.Alternative(exp), env)
}

func evalDefinition(exp int, env int) int {
	return define.DefineVariable(define.DefinitionVariable(exp), Eval(define.DefinitionValue(exp), env), env)
}

func evalAssignment(exp int, env int) int {
	return assign.SetVariableValue(assign.AssignmentVariable(exp), Eval(assign.AssignmentValue(exp), env), env)
}

func textOfQuotation(exp int) int {
	return 0
}

func lookupVariableValue(exp int, env int) int {
	return 0
}
