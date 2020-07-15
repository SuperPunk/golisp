package se

import (
	"strconv"
	"strings"
)

func SelfEvaluating(exp interface{}) bool {
	expString, ok := exp.(string)
	if !ok {
		return false
	}
	if isNumber(expString) {
		return true
	}
	if isString(expString) {
		return true
	}
	return false
}

// 字符串表示为：'apple 'cat 'eat
func isString(exp string) bool {
	return strings.HasPrefix(exp, "'")
}

// 能转化为一个数，就是数字
func isNumber(exp string) bool {
	_, err := strconv.Atoi(exp)
	return err == nil
}
