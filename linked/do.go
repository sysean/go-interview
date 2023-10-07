package linked

// ForTestLinked 创建一个链表
func ForTestLinked() {
	head := NewLinkedList([]int{1, 2, 3, 4, 5})
	head.Print()
}

// ForTestLinked2 递归翻转链表
func ForTestLinked2() {
	head := NewLinkedList([]int{1, 2, 3, 4, 5})
	head = reverseList(head)
	head.Print()
}

// ForTestLinked3 遍历翻转链表
func ForTestLinked3() {
	head := NewLinkedList([]int{1, 2, 3, 4, 5})
	head = reverseList2(head)
	head.Print()
}
