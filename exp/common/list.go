package common

// 构建一个List
func List(values ...interface{}) *Pair {
	if len(values) == 0 {
		return nil
	}
	return Cons(values[0], List(values[1:]...))
}

func Length(p *Pair) int {
	if p == nil {
		return 0
	}
	return 1 + Length(Cdr(p).(*Pair))
}
