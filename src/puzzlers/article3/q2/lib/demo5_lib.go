// Package lib5 不同级目录可以使用不同的package
// 导入与使用逻辑见demo5
package lib5

import "fmt"

// Hello 名称的首字母为大写的程序实体才可以被当前包外的代码引用，否则它就只能被当前包内的其他代码引用。
func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
