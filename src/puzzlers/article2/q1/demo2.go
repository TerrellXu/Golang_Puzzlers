package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	// 接受命令传参，给name变量传值
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
