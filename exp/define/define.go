package define

import "golisp/exp/common"

// ------变量定义
// (define ⟨var⟩ ⟨value⟩)

// ------标准过程定义
// (define (⟨var⟩ ⟨param⟩ … ⟨param⟩) ⟨body⟩)

// ------标准过程定义是以下包含lambda形式的语法糖
// (define ⟨var⟩
//  (lambda (⟨param⟩ … ⟨param⟩)
//    ⟨body⟩))

func Definition(exp []string) bool {
	return common.TaggedList(exp, "define")
}

func DefinitionVariable(exp []string) int {
	//  (if (symbol? (cadr exp))
	//      (cadr exp)
	//      (caadr exp)))
	if common.Symbol(exp[1]) {
		return exp[1]
	}
	panic("error occurs when retrieve definition variable") // todo caadr(exp)是啥意思
}

func DefinitionValue(exp int) int {
	//  (if (symbol? (cadr exp))
	//      (caddr exp)
	//      (make-lambda
	//       (cdadr exp)   ; formal parameters
	//       (cddr exp)))) ; body
	// todo 看不懂书中代码，exp的构造，回过来再看
	return 0
}
