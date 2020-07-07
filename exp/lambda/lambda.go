package lambda

import "golisp/exp/common"

//  (lambda (⟨param⟩ … ⟨param⟩) ⟨body⟩))

func Lambda(exp []string) bool {
	return common.TaggedList(exp, "lambda")
}

func LambdaParameters(exp []string) int {
	return exp[1]
}

func LambdaBody(exp []string) int {
	return exp[2]
}

func MakeLambda(parameters string, body string) []string {
	return []string{"lambda", parameters, body}
}
