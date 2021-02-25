package main

import "fmt"

// LinkNode 单个节点
type LinkNode struct {
	Data     int
	PrevNode *LinkNode
	NextNode *LinkNode
}

func main() {
	// 创建 node1
	node1 := new(LinkNode)
	node1.Data = 1
	node1.PrevNode = nil

	// 创建 node2
	node2 := new(LinkNode)
	node2.Data = 2
	node2.PrevNode = node1

	// 把 node2 链接到 node1
	node1.NextNode = node2

	// 创建 node3
	node3 := new(LinkNode)
	node3.Data = 3
	node3.PrevNode = node2

	// 把 node3 链接到 node2
	node2.NextNode = node3

	fmt.Println(node3.PrevNode.Data)
}
