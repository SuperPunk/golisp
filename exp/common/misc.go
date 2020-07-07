package common

type Proc func(interface{}) interface{}

func IsNull(value interface{}) bool {
	return value == nil
}

func Map(proc Proc, list *Pair) *Pair {
	if list == nil {
		return nil
	}
	return &Pair{proc(Car(list)), Map(proc, Cdr(list).(*Pair))}
}
