package core

import (
	"golisp/exp"
	"strconv"
)

func apply(procedure int, arguments int) int {
	if exp.IsPrimitiveProcedure(procedure) {
		return applyPrimitiveProcedure(procedure, arguments)
	} else if exp.IsCompoundProcedure(procedure) {
		return evalSequence(procedureBody(procedure), extendEnvironment(procedureParameters(procedure), arguments, procedureEnvironment(procedure)))
	} else {
		panic("Unknown procedure type -- APPLY " + strconv.Itoa(procedure))
	}
}

func procedureBody(procedure int) int {
	return 0
}

func extendEnvironment(parameters int, arguments int, environment int) int {
	return 0
}

func procedureEnvironment(procedure int) int {
	return 0
}

func procedureParameters(procedure int) int {
	return 0
}

func applyPrimitiveProcedure(procedure int, arguments int) int {
	return 0
}
