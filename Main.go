package main

import (
	"bufio"
	"fmt"
	"golisp/core"
	"golisp/exp/application"
	"golisp/exp/common"
	"os"
	"strings"
)

// 解释器最终把表达式规约到基本过程(语言原生提供的过程)的应用
// 设置一个全局环境：
//   1）提供基本过程名称-对象的映射
//   2）true和false的符号

func main() {
	driverLoop()
}

var inputPrompt = ";;; M-Eval input:"
var outputPrompt = ";;; M-Eval value:"
var theGlobalEnvironment = application.SetupEnvironment()

// R-E-P Loop
func driverLoop() {
	promptForInput(inputPrompt)
	expression := inputFromCMD()
	tokens := tokenize(expression)
	ast := buildAST(tokens)
	value := core.Eval(ast, theGlobalEnvironment)
	announceOutput(outputPrompt, value)
	driverLoop()
}

func promptForInput(s string) {
	fmt.Printf("\n\n%s\n", s)
}

func announceOutput(prompt string, output interface{}) {
	fmt.Printf("\n%s\n", prompt)
	fmt.Print(output)
}

func inputFromCMD() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString(';')
	return strings.TrimSuffix(input, ";")
}

func tokenize(exp string) []string {
	return strings.Fields(
		strings.Replace(strings.Replace(exp, "(", " ( ", -1), ")", " ) ", -1),
	)
}

func buildAST(tokens []string) *common.Pair {
	var stack []interface{}
	for _, token := range tokens {
		if token == ")" {
			var sub *common.Pair
			for st := stack[len(stack)-1]; st != "("; st = stack[len(stack)-1] {
				sub = common.Cons(st, sub)
				stack = stack[:len(stack)-1]
			}
			stack[len(stack)-1] = sub
		} else {
			stack = append(stack, token)
		}
	}
	return stack[0].(*common.Pair)
}
