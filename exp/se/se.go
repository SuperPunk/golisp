package se

func SelfEvaluating(exp int) bool {
	if isNumber(exp) {
		return true
	}
	if isString(exp) {
		return true
	}
	return false
}

func isString(exp int) bool {
	return true
}

func isNumber(exp int) bool {
	return true
}
