package core

import (
	"github.com/stretchr/testify/assert"
	"golisp/exp/application"
	"golisp/exp/common"
	"testing"
)

// (define (append x y)
//  (if (null? x)
//      y
//      (cons (car x) (append (cdr x) y))))
func TestEval(t *testing.T) {
	globalEnv := application.SetupEnvironment()

	// 定义过程：
	// (define (append x y)
	//  (if (null? x)
	//      y
	//      (cons (car x) (append (cdr x) y))))
	alternative := common.List("cons",
		common.List("car", "x"),
		common.List("append", common.List("cdr", "x"), "y"))
	body := common.List("if", common.List("null?", "x"), "y", alternative)
	exp := common.List("define", common.List("append", "x", "y"), body)
	Eval(exp, globalEnv)

	// 应用过程：
	// (append '(a b c) '(d e f))
	exp = common.List("append",
		common.List("cons", "'a",
			common.List("cons", "'b",
				common.List("cons", "'c", "nil"))),
		common.List("cons", "'d",
			common.List("cons", "'e",
				common.List("cons", "'f", "nil"))))
	value := Eval(exp, globalEnv).(*common.Pair)

	assert.Equal(t, "'a", common.Car(value))
	assert.Equal(t, "'b", common.Cadr(value))
	assert.Equal(t, "'c", common.Caddr(value))
	assert.Equal(t, "'d", common.Cadddr(value))
	assert.Equal(t, "'e", common.Caddddr(value))
	assert.Equal(t, "'f", common.Cadddddr(value))
}

func TestEval2(t *testing.T) {
	globalEnv := application.SetupEnvironment()

	// 求值：
	// (* 5 (+ 2 3))
	exp := common.List("*", 5, common.List("+", 2, 3))
	value := Eval(exp, globalEnv).(*common.Pair)

	assert.Equal(t, "'a", common.Car(value))
	assert.Equal(t, "'b", common.Cadr(value))
	assert.Equal(t, "'c", common.Caddr(value))
	assert.Equal(t, "'d", common.Cadddr(value))
	assert.Equal(t, "'e", common.Caddddr(value))
	assert.Equal(t, "'f", common.Cadddddr(value))
}
