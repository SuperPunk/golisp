package exp

// 自求值表达式，返回自身
func IsSelfEvaluating(exp int) bool {
	return true
}

func IsVariable(exp int) bool {
	return true
}

func IsQuoted(exp int) bool {
	return true
}

func IsAssignment(exp int) bool {
	return true
}

func IsDefinition(exp int) bool {
	return true
}

func IsIf(exp int) bool {
	return true
}

func IsLambda(exp int) bool {
	return true
}

func IsBegin(exp int) bool {
	return true
}

func IsCond(exp int) bool {
	return true
}

func IsApplication(exp int) bool {
	return true
}

func IsPrimitiveProcedure(procedure int) bool {
	return true
}

func IsCompoundProcedure(procedure int) bool {
	return true
}

func NoOperands(exps int) bool {
	return true
}
