package main

import "fmt"

func main() {

	// 数组
	array := [5]int{}
	fmt.Println(array)

	array[0] = 8
	array[1] = 9
	array[2] = 7

	fmt.Println(array)
	fmt.Println(array[2])

	fmt.Println("------------------")

	// 数组链表
	// Array 占用空间小，查询快，直接通过索引获取数据，操作移动、删除效率低下，需要大量移动空间
	// LinkedList 移动、删除数据快，占用空间大，查询需要进行遍历，时间复杂度为O(n)

	// ArrayLink 数组链表
	type ArrayLink struct {
		Data      string
		NextIndex int64
	}

	var arrlink [5]ArrayLink
	arrlink[0] = ArrayLink{"I", 3}    // 下一个节点的下标为3
	arrlink[1] = ArrayLink{"Army", 4} // 下一个节点的下标为4
	arrlink[2] = ArrayLink{"You", 1}  // 下一个节点的下标为1
	arrlink[3] = ArrayLink{"Love", 2} // 下一个节点的下标为2
	arrlink[4] = ArrayLink{"!", -1}   // -1表示没有下一个节点

	dummynode := arrlink[0]

	for {
		fmt.Println(dummynode.Data)
		if dummynode.NextIndex == -1 {
			break
		}

		dummynode = arrlink[dummynode.NextIndex]
	}
}
