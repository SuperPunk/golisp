package common

type Pair struct {
	left  interface{}
	right interface{}
}

func Cons(left, right interface{}) *Pair {
	return &Pair{left, right}
}

func Car(p *Pair) interface{} {
	return p.left
}

func Cdr(p *Pair) interface{} {
	return p.right
}

func Cadr(p *Pair) interface{} {
	return Car(Cdr(p).(*Pair))
}

func SetCar(p *Pair, v interface{}) {
	p.left = v
}

func SetCdr(p *Pair, v interface{}) {
	p.right = v
}
