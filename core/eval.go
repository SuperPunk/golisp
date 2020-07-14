package core

import (
	"golisp/exp/application"
	"golisp/exp/assign"
	"golisp/exp/begin"
	"golisp/exp/cond"
	"golisp/exp/define"
	"golisp/exp/exps"
	"golisp/exp/if"
	"golisp/exp/lambda"
	"golisp/exp/procedure"
	"golisp/exp/quote"
	"golisp/exp/se"
	"golisp/exp/variable"
	"strconv"
)

func Eval(exp, env int) interface{} {
	if se.SelfEvaluating(exp) {
		return exp
	} else if variable.Variable(exp) {
		return application.LookupVariableValue(exp, env)
	} else if quote.Quoted(exp) {
		return textOfQuotation(exp)
	} else if assign.Assignment(exp) {
		return evalAssignment(exp, env)
	} else if define.Definition(exp) {
		return evalDefinition(exp, env)
	} else if _if.If(exp) {
		return evalIf(exp, env)
	} else if lambda.Lambda(exp) {
		return procedure.MakeProcedure(lambda.LambdaParameters(exp), lambda.LambdaBody(exp), env)
	} else if begin.IsBegin(exp) {
		return evalSequence(begin.BeginActions(exp), env)
	} else if cond.IsCond(exp) {
		return Eval(cond.Cond2If(exp), env)
	} else if application.IsApplication(exp) {
		return apply(Eval(application.Operator(exp), env), listOfValues(application.Operands(exp), env))
	} else {
		panic("Unknown expression type -- EVAL " + strconv.Itoa(exp))
	}
}

// 生成实际参数表
func listOfValues(exps int, env int) int {
	if application.NoOperands(exps) {
		return 0
	}
	// todo 定义list结构
	return Eval(application.FirstOperand(exps), env) + listOfValues(application.RestOperands(exps), env)
}

func evalSequence(expressions int, env int) int {
	if exps.LastExp(expressions) {
		return Eval(exps.FirstExp(expressions), env)
	} else {
		Eval(exps.FirstExp(expressions), env)
		return evalSequence(exps.RestExps(expressions), env)
	}
}

func evalIf(exp int, env int) int {
	if _if.True(Eval(_if.IfPredicate(exp), env)) {
		return Eval(_if.IfConsequent(exp), env)
	}
	return Eval(_if.IfAlternative(exp), env)
}

func evalDefinition(exp int, env int) int {
	return application.DefineVariable(define.DefinitionVariable(exp), Eval(define.DefinitionValue(exp), env), env)
}

func evalAssignment(exp int, env int) int {
	return application.SetVariableValue(assign.AssignmentVariable(exp), Eval(assign.AssignmentValue(exp), env), env)
}

func textOfQuotation(exp int) int {
	// todo 回头再看，涉及到exp的表示
	return 0
}

func apply(proc int, arguments int) int {
	if procedure.IsPrimitive(proc) {
		return procedure.ApplyPrimitiveProcedure(proc, arguments)
	} else if procedure.IsCompound(proc) {
		return evalSequence(procedure.Body(proc), application.ExtendEnvironment(procedure.Parameters(proc), arguments, procedure.Environment(proc)))
	} else {
		panic("Unknown procedure type -- APPLY " + strconv.Itoa(proc))
	}
}
