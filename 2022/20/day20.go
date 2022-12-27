package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	file := "input.txt"

	head, length := parse(file, 1)
	mix(head, length, 1)
	fmt.Println("part1:", answer(head))

	head, length = parse(file, 811589153)
	mix(head, length, 10)
	fmt.Println("part2:", answer(head))
}

func parse(filename string, decryptionKey int) (*node, int) {
	nums := util.IntLines(filename)

	var head, iter *node
	for i, num := range nums {
		node := newNode(i, num*decryptionKey)
		if head == nil {
			head = node
			iter = node
		} else {
			insertAfter(iter, node)
			iter = node
		}
	}

	return head, len(nums)
}

func mix(head *node, length int, times int) {
	for t := 0; t < times; t++ {
		for i := 0; i < length; i++ {
			curNode := head
			for curNode.id != i {
				curNode = curNode.next
			}

			advance := curNode.value % (length - 1)
			if advance < 0 {
				advance += length - 1
			}

			afterNode := curNode
			for j := 0; j < advance; j++ {
				afterNode = afterNode.next
			}

			if curNode != afterNode {
				// fmt.Println(curNode.value, "moves between", afterNode.value, "and", afterNode.next.value)
				insertAfter(afterNode, removeAfter(curNode.prev))
			}
		}
	}
}

func answer(head *node) int {
	zero := head
	for zero.value != 0 {
		zero = zero.next
	}

	var x, y, z int
	for i := 0; i < 3000; i++ {
		zero = zero.next
		if i == 999 {
			x = zero.value
		} else if i == 1999 {
			y = zero.value
		} else if i == 2999 {
			z = zero.value
		}
	}

	return x + y + z
}

type node struct {
	id    int
	value int
	next  *node
	prev  *node
}

func newNode(i, value int) *node {
	n := node{
		id:    i,
		value: value,
	}
	n.next = &n
	n.prev = &n
	return &n
}

func insertAfter(before, after *node) {
	after.next = before.next
	after.next.prev = after

	after.prev = before
	before.next = after
}

func removeAfter(before *node) *node {
	removed := before.next

	before.next = removed.next
	before.next.prev = before

	return removed
}

func printNodes(head *node) {
	node := head
	for {
		fmt.Printf("%d ", node.value)
		node = node.next
		if node == head {
			break
		}
	}
	fmt.Println()
}
