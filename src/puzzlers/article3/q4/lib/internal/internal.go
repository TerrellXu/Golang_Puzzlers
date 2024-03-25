// Package internal
// 模块级私有： 我们可以通过创建internal代码包让一些程序实体仅仅能被当前模块中的其他代码引用。
// internal代码包中声明的公开程序实体仅能被该代码包的直接父包及其子包中的代码引用。
// - 当然，引用前需要先导入这个internal包。
package internal

import (
	"fmt"
	"io"
)

func Hello(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s!\n", name)
}
