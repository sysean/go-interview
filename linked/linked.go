package linked

import "fmt"

type Node struct {
	Val  int
	next *Node
}

func NewLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.next = &Node{Val: arr[i]}
		cur = cur.next
	}

	return head
}

func (n *Node) Print() {
	fmt.Printf("Print linked list: ")
	for cur := n; cur != nil; cur = cur.next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Println()
}

// 1 ->    2 -> 3 -> 4 -> 5
// 1 -> <- 2 <- 3 <- 4 <- 5
// 递归翻转
func reverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	tmp := reverseList(head.next)
	head.next.next = head
	head.next = nil
	return tmp
}

// 迭代翻转
// 1 -> 2 -> 3 -> 4 -> 5
//      |
//      p
// p.next = head

// 1 -><- 2   3 -> 4 -> 5
//        |
//        p

func reverseList2(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
