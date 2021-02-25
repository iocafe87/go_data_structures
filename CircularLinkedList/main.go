package main

// CircularLinkNode 单个节点
type CircularLinkNode struct {
	PrevNode *CircularLinkNode // 前驱节点
	NextNode *CircularLinkNode // 后继节点
	Data     interface{}       // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (c *CircularLinkNode) init() *CircularLinkNode {
	c.NextNode = c
	c.PrevNode = c
	return c
}

func main() {
	r := new(CircularLinkNode)
	r.init()
}
