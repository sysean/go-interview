package main

// 指令集，统一都是64位

const (
	IMM = iota
	LC
	LI
	SC
	SI
	ADD
	SUB
	IMUL // 乘法
	IDIV // 除法
	EXIT // 退出
	ENT  // 进入函数，并为其在栈上分配空间
	LEA
	PUSH // 入栈
)
