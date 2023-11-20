package main

import (
	"fmt"
	"os"

	"gopkg.in/olebedev/go-duktape.v3"
)

func main() {
	// 创建一个新的 Duktape 上下文
	ctx := duktape.New()

	// read JavaScript file
	jsFile := os.Args[1]
	jsCode, err := os.ReadFile(jsFile)
	if err != nil {
		fmt.Println("Error reading JavaScript file:", err)
		return
	}

	// exec JavaScript code
	if err := ctx.PevalString(string(jsCode)); err != nil {
		fmt.Println("Error executing JavaScript:", err)
		return
	}

	// get JavaScript output
	//result := ctx.SafeToString(-1)
	//fmt.Println("JavaScript output:", result)

	// close Duktape context
	ctx.DestroyHeap()
}
