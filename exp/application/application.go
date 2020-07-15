package application

import (
	"golisp/exp/common"
	"log"
)

var TheEmptyEnvironment = common.List()

func IsApplication(exp interface{}) bool {
	return common.IsPair(exp)
}

func Operator(exp *common.Pair) interface{} {
	return common.Car(exp)
}

func Operands(exp *common.Pair) *common.Pair {
	return common.Cdr(exp).(*common.Pair)
}

func NoOperands(ops *common.Pair) bool {
	return ops == nil
}

func FirstOperand(ops *common.Pair) interface{} {
	return common.Car(ops)
}

func RestOperands(ops *common.Pair) *common.Pair {
	return common.Cdr(ops).(*common.Pair)
}

// 查询变量值
func LookupVariableValue(variable interface{}, env *common.Pair) interface{} {
	var envLoop func(*common.Pair) interface{}

	envLoop = func(env *common.Pair) interface{} {
		var scan func(*common.Pair, *common.Pair) interface{}

		scan = func(vars, vals *common.Pair) interface{} {
			if vars == nil {
				return envLoop(EnclosingEnvironment(env))
			}
			if variable == common.Car(vars) {
				return common.Car(vals)
			}
			return scan(common.Cdr(vars).(*common.Pair), common.Cdr(vals).(*common.Pair))
		}

		if env == TheEmptyEnvironment {
			log.Fatal("unbound variable", variable)
			return nil
		}

		return func(frame *common.Pair) interface{} {
			return scan(FrameVariables(frame), FrameValues(frame))
		}(FirstFrame(env))
	}
	return envLoop(env)
}

// 改变env中绑定的variable的值
func SetVariableValue(variable interface{}, value interface{}, env *common.Pair) {
	var envLoop func(*common.Pair)

	envLoop = func(env *common.Pair) {
		var scan func(*common.Pair, *common.Pair)

		scan = func(vars, vals *common.Pair) {
			if vars == nil {
				envLoop(EnclosingEnvironment(env))
			}
			if variable == common.Car(vars) {
				common.SetCar(vals, value)
			}
			scan(common.Cdr(vars).(*common.Pair), common.Cdr(vals).(*common.Pair))
		}

		if env == TheEmptyEnvironment {
			log.Fatal("unbound variable", variable)
		} else {
			func(frame *common.Pair) {
				scan(FrameVariables(frame), FrameValues(frame))
			}(FirstFrame(env))
		}
	}

	envLoop(env)
}

// 在env的第一个frame中绑定一个变量-值
func DefineVariable(variable interface{}, value interface{}, env *common.Pair) {
	func(frame *common.Pair) {
		var scan func(*common.Pair, *common.Pair)

		scan = func(vars *common.Pair, vals *common.Pair) {
			if vars == nil {
				AddBindingToFrame(variable, value, frame)
			} else if variable == common.Car(vars) {
				common.SetCar(vals, value)
			} else {
				scan(common.Cdr(vars).(*common.Pair), common.Cdr(vals).(*common.Pair))
			}
		}

		scan(FrameVariables(frame), FrameValues(frame))
	}(FirstFrame(env))
}

// 返回一个新的env，包含新创建的frame，该frame中将variables绑定于values，
func ExtendEnvironment(variables *common.Pair, values *common.Pair, baseEnv *common.Pair) *common.Pair {
	if common.Length(variables) == common.Length(values) {
		return common.Cons(MakeFrame(variables, values), baseEnv)
	}
	log.Fatal("变量和参数数量不一致")
	return nil
}

// 一个环境的外围环境，就是该环境的cdr
func EnclosingEnvironment(env *common.Pair) *common.Pair {
	return common.Cdr(env).(*common.Pair)
}

func FirstFrame(env *common.Pair) *common.Pair {
	return common.Car(env).(*common.Pair)
}

// 返回一个variables和values的pair
func MakeFrame(variables *common.Pair, values *common.Pair) *common.Pair {
	return common.Cons(variables, values)
}

// 变量序列
func FrameVariables(frame *common.Pair) *common.Pair {
	return common.Car(frame).(*common.Pair)
}

// 值序列
func FrameValues(frame *common.Pair) *common.Pair {
	return common.Cdr(frame).(*common.Pair)
}

func AddBindingToFrame(variable interface{}, value interface{}, frame *common.Pair) {
	common.SetCar(frame, common.Cons(variable, common.Car(frame)))
	common.SetCdr(frame, common.Cons(value, common.Cdr(frame)))
}
