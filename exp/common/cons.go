package common

type Pair struct {
	left  interface{}
	right interface{}
}

func Cons(left, right interface{}) *Pair {
	return &Pair{left, right}
}

// todo 用于map，不知道有没有什么更好的方式
func MCar(p interface{}) interface{} {
	return Car(p.(*Pair))
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

func Caddr(p *Pair) interface{} {
	return Car(Cdr(Cdr(p).(*Pair)).(*Pair))
}

func Caadr(p *Pair) interface{} {
	return Car(Car(Cdr(p).(*Pair)).(*Pair))
}

func Cdadr(p *Pair) interface{} {
	return Cdr(Car(Cdr(p).(*Pair)).(*Pair))
}

func Cddr(p *Pair) interface{} {
	return Cdr(Cdr(p).(*Pair))
}

func Cdddr(p *Pair) interface{} {
	return Cdr(Cdr(Cdr(p).(*Pair)).(*Pair))
}

func Cadddr(p *Pair) interface{} {
	return Car(Cdr(Cdr(Cdr(p).(*Pair)).(*Pair)).(*Pair))
}

func Caddddr(p *Pair) interface{} {
	return Car(Cdr(Cdr(Cdr(Cdr(p).(*Pair)).(*Pair)).(*Pair)).(*Pair))
}

func Cadddddr(p *Pair) interface{} {
	return Car(Cdr(Cdr(Cdr(Cdr(Cdr(p).(*Pair)).(*Pair)).(*Pair)).(*Pair)).(*Pair))
}

func SetCar(p *Pair, v interface{}) {
	p.left = v
}

func SetCdr(p *Pair, v interface{}) {
	p.right = v
}
