package common

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (p *Pair) String() string {
	return fmt.Sprintf("(%s)", p.convertToStr())
}

func (p *Pair) convertToStr() string {
	if p == nil {
		return ""
	}
	var str strings.Builder
	str.WriteString(conv2str(p.left))
	if IsNull(p.right) {
		return str.String()
	}
	str.WriteString(fmt.Sprintf(" %s", p.right.(*Pair).convertToStr()))
	return str.String()
}

func conv2str(v interface{}) string {
	if i, ok := v.(int); ok {
		return strconv.Itoa(i)
	}
	if p, ok := v.(*Pair); ok {
		return p.String()
	}
	return v.(string)
}
