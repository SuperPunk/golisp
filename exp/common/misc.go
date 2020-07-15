package common

type Proc func(interface{}) interface{}

func IsNull(value interface{}) bool {
	return value == nil
}

// todo check
func IsPair(value interface{}) bool {
	pair1, ok1 := value.(*Pair)
	_, ok2 := Cdr(pair1).(*Pair)
	return ok1 && !ok2
}

func Map(proc Proc, list *Pair) *Pair {
	if list == nil {
		return nil
	}
	return &Pair{proc(Car(list)), Map(proc, Cdr(list).(*Pair))}
}
