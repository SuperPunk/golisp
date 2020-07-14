package procedure

import (
	"golisp/exp/common"
	"reflect"
)

func IsPrimitive(proc *common.Pair) bool {
	return common.TaggedList(proc, "primitive")
}

func PrimitiveImplementation(proc *common.Pair) interface{} {
	return common.Cadr(proc)
}

func ApplyPrimitiveProcedure(proc *common.Pair, args *common.Pair) interface{} {
	return applyInUnderlyingGolang(PrimitiveImplementation(proc), args)
}

func MakeProcedure(parameters interface{}, body interface{}, env interface{}) *common.Pair {
	return common.List("procedure", parameters, body, env)
}

func IsCompound(proc []string) bool {
	return common.TaggedList(proc, "procedure")
}

func Parameters(proc []string) string {
	return proc[1]
}

func Body(proc []string) string {
	return proc[2]
}

func Environment(proc []string) string {
	return proc[3]
}

func PrimitiveProcedures() *common.Pair {
	return common.List(
		common.List("car", common.Car),
		common.List("cdr", common.Cdr),
		common.List("cons", common.Cons),
		common.List("null?", common.IsNull),
		// add more primitive procedures here, if you wish...
	)
}

func PrimitiveNames() *common.Pair {
	return common.Map(common.MCar, PrimitiveProcedures())
}

func PrimitiveObjects() *common.Pair {
	return common.Map(func(proc interface{}) interface{} {
		return common.List("primitive", common.Cadr(proc.(*common.Pair)))
	}, PrimitiveProcedures())
}

// 此处反射执行函数
func applyInUnderlyingGolang(proc interface{}, args *common.Pair) interface{} {
	var castArgs func(*common.Pair) []reflect.Value
	castArgs = func(args *common.Pair) []reflect.Value {
		if args == nil {
			return nil
		}
		var s []reflect.Value
		s = append(s, reflect.ValueOf(common.Car(args)))
		s = append(s, castArgs(common.Cdr(args).(*common.Pair))...)
		return s
	}

	return reflect.ValueOf(proc).Call(castArgs(args))[0]
}
