package main

const PoolSize = 1024

var (
	vm *VM
	PC int // 程序计数器
	SP int // 栈指针
	BP int // 基址指针
	AX int // 通用寄存器
)

func init() {
	vm = &VM{}

	BP = PoolSize
	SP = PoolSize // 指向栈的高地址，入栈时减小
}

// VM 虚拟机, 为了简化，没有堆，堆直接使用本解释器的堆
type VM struct {
	Stack [PoolSize]int
	Text  [PoolSize]int // 存放代码(指令)
	Data  [PoolSize]int // 只存字符串
}
