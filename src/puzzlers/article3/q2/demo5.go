package main

import (
	"flag"
	"puzzlers/article3/q2/lib"
)

//导入的路径是src下面的相对路径，
//而使用的时候是代码里声明的package名字。为了避免困惑，建议package的名字跟父级目录保持一致。

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

// 导入包时，import的是相对src的相对文件路径： puzzlers/article3/q2/lib
// 使用包内的函数时，其限定符是：包名.函数名(),压根与程序的文件名没有啥关系. lib5.Hello(name)
func main() {
	flag.Parse()
	lib5.Hello(name)
}
