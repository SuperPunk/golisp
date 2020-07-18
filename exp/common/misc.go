package common

import "reflect"

type Proc func(interface{}) interface{}

// todo 判断interface是否为空，这串代码不知道是什么意思，有时间再看看go里面的nil吧。（nil==nil is false）
func IsNull(value interface{}) bool {
	vi := reflect.ValueOf(value)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func IsPair(value interface{}) bool {
	_, ok := value.(*Pair)
	return ok
}

func Map(proc Proc, list *Pair) *Pair {
	if list == nil {
		return nil
	}
	return &Pair{proc(Car(list)), Map(proc, Cdr(list).(*Pair))}
}
