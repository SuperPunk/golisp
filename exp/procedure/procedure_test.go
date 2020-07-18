package procedure

import (
	"fmt"
	"golisp/exp/common"
	"testing"
)

func Test_applyInUnderlyingGolang(t *testing.T) {
	v := applyInUnderlyingGolang(common.Cons, common.List("a", nil))
	fmt.Println(v)
}
