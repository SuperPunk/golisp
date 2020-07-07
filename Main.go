package golisp

import (
	"golisp/exp/application"
	"golisp/exp/common"
	"golisp/exp/procedure"
)

// 解释器最终把表达式规约到基本过程(语言原生提供的过程)的应用
// 设置一个全局环境：
//   1）提供基本过程名称-对象的映射
//   2）true和false的符号

var theGlobalEnvironment = setupEnvironment()

func setupEnvironment() *common.Pair {
	return func(initialEnv *common.Pair) *common.Pair {
		application.DefineVariable("true", true, initialEnv)
		application.DefineVariable("false", false, initialEnv)
		return initialEnv
	}(application.ExtendEnvironment(procedure.PrimitiveNames(), procedure.PrimitiveObjects(), application.TheEmptyEnvironment))
}
