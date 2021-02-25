package main

import "fmt"

// LinkNode 单个节点
type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	// 新的节点
	node := new(LinkNode)
	node.Data = 2

	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	// 将 node1 链接到 node 节点上
	node.NextNode = node1

	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4

	// 将 node2 链接到 node1 节点上
	node1.NextNode = node2

	// 按顺序打印数据
	// 哑节点，无法确定 node 是否为空
	// 直接使用可能有空异常
	dummyNode := node

	for {
		if dummyNode != nil {
			// 打印节点值
			fmt.Println(dummyNode.Data)
			// 获取下一个节点
			dummyNode = dummyNode.NextNode
			continue
		}

		// 为空，则链表结束
		break
	}
}
