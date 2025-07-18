package main

import (
	"fmt"
	"slices"
)

// An internal node is a node who has children and so is opposite to a leaf node.
// In a min-heap, it can have a maximum of 2 children.
type Node struct {
	Frequency int16
	LeftNode  *Node
	RightNode *Node
	Symbol    string
}

// type LeafNode struct {
// 	Symbol    string
// 	Frequency int16
// }

func main() {
	// Min-heap LOL
	// There are a bunch of rules
	// in a min heap parents can't be larger than any of their children.
	// Every parent can have a maximum of 2 children.
	// levels must be filled from left to right. There can NOT be any gaps in the middle.
	// Min-heap is a binary tree.
	// The below is an acceptable min-heap because the binary tree looks like this:
	//         1         (0)
	//       /   \
	//      2     3      (1,2)
	//     / \   /
	//    4   5 6        (3,4,5)

	// Min-heap priority queue sorted based on the frequencies of the symbols:
	nodes := []*Node{
		{
			Symbol:    "a",
			Frequency: 5,
		},
		{
			Symbol:    "b",
			Frequency: 9,
		},
		{
			Symbol:    "c",
			Frequency: 12,
		},
		{
			Symbol:    "d",
			Frequency: 13,
		},
		{
			Symbol:    "e",
			Frequency: 16,
		},
		{
			Symbol:    "f",
			Frequency: 45,
		},
	}

	fmt.Println(nodes)

	nodes = extract(nodes)

	fmt.Println(nodes)

	nodes = extract(nodes)

	fmt.Println(nodes)

	nodes = extract(nodes)

	fmt.Println(nodes)

	nodes = extract(nodes)

	fmt.Println(nodes)

	nodes = extract(nodes)

	fmt.Println(*nodes[0])
	fmt.Println(*nodes[0].LeftNode)
	fmt.Println(*nodes[0].RightNode)

	// leftNode := nodes[0].LeftNode
	// rightNode := nodes[0].RightNode
	// for {
	// 	fmt.Println(leftNode)
	// 	fmt.Println(rightNode)
	// 	if leftNode == nil || rightNode == nil {
	// 		break
	// 	}
	// 	leftNode = leftNode.LeftNode
	// 	rightNode = leftNode.RightNode
	// }
}

// Modifying the length of a slice in Go requires us to return a new slice (i.e., a new slice header)
// void functions can't be used
func extract(nodes []*Node) []*Node {
	nodeLength := len(nodes)
	fmt.Printf("Length of nodes: %d \n", nodeLength)

	var newNode *Node
	if nodeLength >= 2 {
		leftNode := nodes[0]
		rightNode := nodes[1]
		newNode = &Node{
			Frequency: nodes[0].Frequency + nodes[1].Frequency,
			LeftNode:  leftNode,
			RightNode: rightNode,
			Symbol:    "",
		}
		// NOT LIKE THIS: WHY?
		// LeftNode: &Node{
		// 		Symbol:    nodes[0].Symbol,
		// 		Frequency: nodes[0].Frequency,
		// 	},
		// 	RightNode: &Node{
		// 		Symbol:    nodes[1].Symbol,
		// 		Frequency: nodes[1].Frequency,
		// 	},
	} else if nodeLength == 1 {
		leftNode := nodes[0]
		newNode = &Node{
			Frequency: nodes[0].Frequency,
			LeftNode:  leftNode,
			RightNode: nil,
			Symbol:    "",
		}
	} else {
		newNode = &Node{
			Frequency: 0,
			LeftNode:  nil,
			RightNode: nil,
			Symbol:    "",
		}
	}

	fmt.Printf("New node: %v \n", *newNode)
	fmt.Printf("Left node: %v \n", *newNode.LeftNode)
	fmt.Printf("Right node: %v \n", *newNode.RightNode)

	newNodes := nodes[2:]
	// newNodes := slices.Delete(nodes, 0, 2)

	if len(newNodes) > 0 {
		for i, node := range newNodes {
			fmt.Printf("Index: %d \n", i)
			if node.Frequency > newNode.Frequency {
				// firstHalf := newNodes[:i]
				// otherHalf := newNodes[i:]
				// newNodes = slices.Concat(firstHalf, []*Node{newNode}, otherHalf)
				newNodes = slices.Insert(newNodes, i, newNode)
				break
			}

			if (i == (len(newNodes) - 1)) && newNode.Frequency > node.Frequency {
				// firstHalf := newNodes[:i]
				// otherHalf := newNodes[i:]
				// newNodes = slices.Concat(newNodes, []*Node{newNode})
				newNodes = slices.Insert(newNodes, i, newNode)
				// newNodes = []Node{newNode}
				break
			}
		}
	} else {
		newNodes = []*Node{newNode}
	}
	return newNodes
}
