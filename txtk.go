package txtk

import (
	"fmt"
	"strings"
)

// EndsWith 检查字符串strA结尾是否是subStrA
func EndsWith(strA string, subStrA string) bool {

	return strings.HasSuffix(strA, subStrA)
}

// Pr 仅仅是封装了fmt.Print函数
func Pr(formatA string, argsA ...interface{}) {
	fmt.Printf(formatA, argsA...)
}

// Prl 类似Pr，但保证结尾有一个回车
func Prl(formatA string, argsA ...interface{}) {
	if EndsWith(formatA, "\n") {
		fmt.Printf(formatA, argsA...)
	} else {
		fmt.Printf(formatA+"\n", argsA...)
	}
}
