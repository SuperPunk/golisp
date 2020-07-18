package main

import (
	"bufio"
	"fmt"
	"golisp/core"
	"golisp/exp/application"
	"golisp/exp/common"
	"golisp/exp/procedure"
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
	func(input string) {
		func(output interface{}) {
			announceOutput(outputPrompt)
			userPrint(output)
		}(core.Eval(input, theGlobalEnvironment))
	}(inputFromCMD())
	driverLoop()
}

func promptForInput(s string) {
	fmt.Printf("\n\n%s\n", s)
}

func announceOutput(s string) {
	fmt.Printf("\n%s\n", s)
}

func inputFromCMD() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func userPrint(object interface{}) {
	// todo check
	if v, ok := object.(*common.Pair); ok && procedure.IsCompound(v) {
		fmt.Print(common.List("compound-procedure", procedure.Parameters(v), procedure.Body(v)))
	}
	fmt.Print(object)
}
