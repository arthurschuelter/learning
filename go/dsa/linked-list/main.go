package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

func (n *Node) Link(next *Node) *Node {
	n.Next = next
	return next
}

func PrintList(n *Node) {
	cur := n
	fmt.Printf("%d -> ", cur.Data)
	for cur.Next != nil {
		cur = cur.Next
		fmt.Printf("%d -> ", cur.Data)
	}

	fmt.Println("||")

}

func ReverseList(n *Node) *Node {
	var prev *Node
	curr := n

	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

func main() {
	fmt.Println("Linked Lists ===")

	head := NewNode(0)

	head.Link(NewNode(1)).Link(NewNode(2)).Link(NewNode(3)).Link(NewNode(4)).Link(NewNode(5))

	PrintList(head)
	head = ReverseList(head)
	PrintList(head)

}
