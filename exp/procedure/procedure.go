package procedure

import "golisp/exp/common"

func IsPrimitive(proc *common.Pair) bool {
	return common.TaggedList(proc, "primitive")
}

func PrimitiveImplementation(proc *common.Pair) interface{} {
	return common.Cadr(proc)
}

func ApplyPrimitiveProcedure(proc int, args int) int {
	return 0
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
	return nil
}

func PrimitiveObjects() *common.Pair {
	return nil
}
