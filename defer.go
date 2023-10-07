package main

import "fmt"

func ForTestOnReturnAdd() {
	a := onReturnAdd()
	fmt.Printf("[ForTestOnReturnAdd] onReturnAdd : a = %d", a)
}

func onReturnAdd() int {
	a := 1
	defer func() {
		a += 1
	}()
	return a
}

func ForTestOnReturnAdd2() {
	a := onReturnAdd2()
	fmt.Printf("[ForTestOnReturnAdd2] onReturnAdd2 : a = %d", a)
}

func onReturnAdd2() (a int) {
	a = 1
	defer func() {
		a += 1
	}()
	return a
}

func deferTestOne() {
	defer fmt.Println(func() string {
		fmt.Println("defer func 01")
		return "defer func 02"
	}())
}
