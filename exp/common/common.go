package common

func TaggedList(exp *Pair, tag string) bool {
	return Car(exp) == tag
}

// todo 变量由符号表示
func Symbol(exp int) bool {
	return true
}
