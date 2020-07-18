package define

import (
	"golisp/exp/common"
	"golisp/exp/lambda"
)

// ------变量定义
// (define ⟨var⟩ ⟨value⟩)

// ------标准过程定义
// (define (⟨var⟩ ⟨param⟩ … ⟨param⟩) ⟨body⟩)

// ------标准过程定义是以下包含lambda形式的语法糖, 转化为以下lambda表达式
// (define ⟨var⟩
//  (lambda (⟨param⟩ … ⟨param⟩)
//    ⟨body⟩))

func Definition(expression interface{}) bool {
	exp, ok := expression.(*common.Pair)
	return ok && common.TaggedList(exp, "define")
}

func DefinitionVariable(exp *common.Pair) string {
	if common.IsSymbol(common.Cadr(exp)) { // 变量定义
		return common.Cadr(exp).(string)
	}
	return common.Caadr(exp).(string) // 过程定义
}

func DefinitionValue(exp *common.Pair) interface{} {
	if common.IsSymbol(common.Cadr(exp)) { // 变量定义
		return common.Caddr(exp)
	}
	return lambda.MakeLambda(common.Cdadr(exp).(*common.Pair), common.Cddr(exp)) // 过程定义
}
