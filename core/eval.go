package core

import (
	"golisp/exp/application"
	"golisp/exp/assign"
	"golisp/exp/begin"
	"golisp/exp/common"
	"golisp/exp/cond"
	"golisp/exp/define"
	"golisp/exp/exps"
	"golisp/exp/if"
	"golisp/exp/lambda"
	"golisp/exp/procedure"
	"golisp/exp/quote"
	"golisp/exp/se"
	"golisp/exp/variable"
	"log"
)

func Eval(exp interface{}, env *common.Pair) interface{} {
	if se.SelfEvaluating(exp) {
		return exp
	} else if variable.Variable(exp) {
		return application.LookupVariableValue(exp, env)
	} else if quote.IsQuoted(exp) {
		return quote.TextOfQuotation(exp.(*common.Pair))
	} else if assign.Assignment(exp) {
		evalAssignment(exp.(*common.Pair), env)
		return nil // todo check
	} else if define.Definition(exp) {
		evalDefinition(exp.(*common.Pair), env)
		return "ok"
	} else if _if.If(exp) {
		return evalIf(exp.(*common.Pair), env)
	} else if lambda.IsLambda(exp) {
		return procedure.MakeProcedure(lambda.Parameters(exp.(*common.Pair)), lambda.Body(exp.(*common.Pair)), env)
	} else if begin.IsBegin(exp) {
		return evalSequence(begin.Actions(exp.(*common.Pair)), env)
	} else if cond.IsCond(exp) {
		return Eval(cond.ToIf(exp.(*common.Pair)), env)
	} else if application.IsApplication(exp) {
		return apply(Eval(application.Operator(exp.(*common.Pair)), env), listOfValues(application.Operands(exp.(*common.Pair)), env))
	} else {
		log.Fatal("Unknown expression type -- EVAL ", exp)
		return nil
	}
}

// 生成实际参数表
func listOfValues(exps *common.Pair, env *common.Pair) *common.Pair {
	if application.NoOperands(exps) {
		return common.List()
	}
	return common.Cons(Eval(application.FirstOperand(exps), env), listOfValues(application.RestOperands(exps), env))
}

// 返回值是一串表达式的最后一个值
func evalSequence(expressions, env *common.Pair) interface{} {
	if exps.IsLastExp(expressions) {
		return Eval(exps.FirstExp(expressions), env)
	} else {
		Eval(exps.FirstExp(expressions), env)
		return evalSequence(exps.RestExps(expressions), env)
	}
}

func evalIf(exp, env *common.Pair) interface{} {
	if _if.True(Eval(_if.IfPredicate(exp), env)) {
		return Eval(_if.IfConsequent(exp), env)
	}
	return Eval(_if.IfAlternative(exp), env)
}

func evalDefinition(exp, env *common.Pair) {
	application.DefineVariable(define.DefinitionVariable(exp), Eval(define.DefinitionValue(exp), env), env)
}

func evalAssignment(exp, env *common.Pair) {
	application.SetVariableValue(assign.AssignmentVariable(exp), Eval(assign.AssignmentValue(exp), env), env)
}

func apply(proc interface{}, arguments *common.Pair) interface{} {
	if procedure.IsPrimitive(proc) {
		return procedure.ApplyPrimitiveProcedure(proc.(*common.Pair), arguments)
	} else if procedure.IsCompound(proc) {
		return evalSequence(
			procedure.Body(proc.(*common.Pair)),
			application.ExtendEnvironment(procedure.Parameters(proc.(*common.Pair)), arguments, procedure.Environment(proc.(*common.Pair))))
	} else {
		log.Fatal("Unknown procedure type -- APPLY ", proc)
		return nil
	}
}
