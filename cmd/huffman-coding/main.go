package main

import (
	"container/heap"
	"fmt"
)

// Node represents either an Internal Node or a Leaf Node
// Internal Node is a Node with children and a Leaf Node is childless
// In a Binary Tree, each Internal Node can have a maximum of 2 children (left and right).
// In a Min-Heap, which implements the Priority Queue ADT, the Binary Tree must be complete
// (meaning that it's filled left-to-right except maybe the last level), and the heap invariant
// where every parent node's Frequency is <= the Frequency of each of its children individually
// must be satisfied.
type Node struct {
	Frequency  int
	LeftNode   *Node
	RightNode  *Node
	Symbol     rune // nil for Internal Nodes
	Discovered bool
}

// MinHeap is a Data Structure (DT) which is an implementation of
// the priority queue ADT. One of the simplest ways (or canonical ways)
// to construct the Huffman Coding algorithm is to use a priority queue
// where the node with the lowest probability/frequency is given the
// highest priority.
type MinHeap []*Node

func (h MinHeap) Len() int {
	return len(h)
}

// Less returns whether h[i] < h[j]; this makes it a min-heap
func (h MinHeap) Less(i, j int) bool {
	return h[i].Frequency < h[j].Frequency
}

// Swap swaps elements
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds x to the heap (called by heap.Push)
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Node)) // Type assertion to *Node
}

// Pop removes and returns last element (called by heap.Pop)
func (h *MinHeap) Pop() any {
	old := *h
	length := len(old)
	lastNode := old[length-1]
	*h = old[0 : length-1]
	return lastNode
}

func main() {
	minHeap := &MinHeap{
		{
			Symbol:    'a',
			Frequency: 5,
		},
		{
			Symbol:    'b',
			Frequency: 9,
		},
		{
			Symbol:    'c',
			Frequency: 12,
		},
		{
			Symbol:    'd',
			Frequency: 13,
		},
		{
			Symbol:    'e',
			Frequency: 16,
		},
		{
			Symbol:    'f',
			Frequency: 45,
		},
	}

	// Establish the heap invariants and validate/sanitize the minHeap
	heap.Init(minHeap)

	constructHuffmanTree(minHeap)

	fmt.Println((*minHeap)[0].LeftNode)

	codeWords := make(map[string]string)

	dfs((*minHeap)[0], "", codeWords)

	fmt.Println(codeWords)
}

// dfs is a Depth-First Search function that traverses
// the built Huffman Tree and constructs a map of the symbols
// mapped to codewords. The codes (or codewords) in the Huffman
// Coding algorithm consist of binary digits (0 or 1, base 2). This is
// because compression is about minimizing entropy and redundancy
// on the bit level. After all, the goal of compression is to
// reduce the amount of bits.
func dfs(node *Node, code string, codeWords map[string]string) {

	if node == nil {
		return
	}

	// If it's a Leaf Node, add the constructed code word to codeWords map
	if node.LeftNode == nil && node.RightNode == nil {
		codeWords[string(node.Symbol)] = code
		return
	}

	// If the Internal Node has a left node, recursively call dfs() while appending "0" to `code`.
	dfs(node.LeftNode, code+"0", codeWords)

	// If the Internal Node has a right node, recursively call dfs() while appending "1" to `code`.
	dfs(node.RightNode, code+"1", codeWords)
}

// constructHuffmanTree is a function that extracts 2 lowest-frequency
// Nodes to build a binary tree from the bottom up (i.e.,
// the extracted Nodes are combined into a new Internal Node
// which is then reinserted into the min-heap). This process
// should be repeated until there is only 1 Node left, which
// becomes the root of the final Huffman Tree.
func constructHuffmanTree(nodes *MinHeap) {
	nodeLength := nodes.Len()

	for nodeLength > 1 {
		fmt.Printf("Length of nodes: %d \n", nodeLength)
		var newNode *Node
		if nodeLength >= 2 {
			// Remove two nodes with the highest priority (i.e., lowest frequency) from the min-heap.
			leftNode := heap.Pop(nodes).(*Node)
			rightNode := heap.Pop(nodes).(*Node)
			newNode = &Node{
				Frequency: leftNode.Frequency + rightNode.Frequency,
				LeftNode:  leftNode,
				RightNode: rightNode,
				Symbol:    0,
			}
		} else if nodeLength == 1 {
			leftNode := heap.Pop(nodes).(*Node)
			// The below should'nt be used. Why?
			// leftNode := heap.Remove(nodes, 0).(*Node)
			// heap.Init(nodes)
			newNode = &Node{
				Frequency: leftNode.Frequency,
				LeftNode:  leftNode,
				RightNode: nil,
				Symbol:    0,
			}
		} else {
			newNode = &Node{
				Frequency: 0,
				LeftNode:  nil,
				RightNode: nil,
				Symbol:    0,
			}
		}

		nodes.Push(newNode)

		nodeLength = nodes.Len()
	}
}
