package common

func TaggedList(exp *Pair, tag string) bool {
	return Car(exp) == tag
}

func IsSymbol(expression interface{}) bool {
	exp, ok := expression.(string)
	if !ok {
		return false
	}
	if exp == "" {
		return true
	}
	return isLetter(exp[0]) && IsSymbol(exp[1:])
}

func isLetter(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '?' || c == '+' || c == '-' || c == '*' || c == '/'
}
