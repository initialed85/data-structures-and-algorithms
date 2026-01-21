package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
)

/*
The binary tree is a tree data structure wherein the left child of each node will be less than the value
of the node and the right child of each node will be greater than the value of each node; time complexity
to search a binary tree will be between O(log n) for balanced trees (created from sorted inputs) to O(n)
for unbalanced trees (created from unsorted inputs).
*/

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
	path  []int
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (t *BinaryTree) Add(value int) {
	var node = t

	for node.value != value {
		if value < node.value {
			if node.left == nil {
				node.left = &BinaryTree{
					value: value,
				}
			}

			node = node.left
		} else if value > node.value {
			if node.right == nil {
				node.right = &BinaryTree{
					value: value,
				}
			}

			node = node.right
		}
	}
}

func (t *BinaryTree) WriteToFile() error {
	data := "digraph {\n"

	var n = t

	stack := NewStack[*BinaryTree]()

	stack.Push(n)

	for n != nil {
		nPtr := stack.Pop()

		if nPtr == nil {
			break
		}

		n := *nPtr

		if n.left != nil {
			data += fmt.Sprintf("%d -> %d;\n", n.value, n.left.value)
			stack.Push(n.left)
		}

		if n.right != nil {
			data += fmt.Sprintf("%d -> %d;\n", n.value, n.right.value)
			stack.Push(n.right)
		}
	}

	data += "}\n"

	err := os.WriteFile("./binary_tree.dot", []byte(data), 0o777)
	if err != nil {
		return err
	}

	err = exec.Command("dot", "-Tpng", "binary_tree.dot", "-o", "binary_tree.png").Run()
	if err != nil {
		return err
	}

	return nil
}

func GetDeepestBinaryTreePath(array [1024]int) []int {
	t := NewBinaryTree(array[0])

	for i := 1; i < 1024; i++ {
		t.Add(array[i])
	}

	err := t.WriteToFile()
	if err != nil {
		panic(err)
	}

	var n = t

	stack := NewStack[*BinaryTree]()

	stack.Push(n)

	n.path = make([]int, 0)

	paths := make([][]int, 0)

	lastLength := 0

	for n != nil {
		nPtr := stack.Pop()

		if nPtr == nil {
			break
		}

		n := *nPtr

		n.path = append(n.path, n.value)

		if n.left != nil {
			stack.Push(n.left)
			n.left.path = slices.Clone(n.path)
		}

		if n.right != nil {
			stack.Push(n.right)
			n.right.path = slices.Clone(n.path)
		}

		if stack.Length() < lastLength {
			paths = append(paths, n.path)
		}

		lastLength = stack.Length()
	}

	// for _, path := range paths {
	// 	log.Printf("%#+v", path)
	// }

	slices.SortFunc(paths, func(a []int, b []int) int {
		if len(a) < len(b) {
			return -1
		} else if len(a) > len(b) {
			return 1
		} else {
			return 0
		}
	})

	return paths[len(paths)-1]
}

func GetBroadestBinaryTreeLevel(array [1024]int) []int {
	t := NewBinaryTree(array[0])

	for i := 1; i < 1024; i++ {
		t.Add(array[i])
	}

	var n = t

	queue := NewQueue[*BinaryTree]()

	queue.Push(n)

	levels := make([][]int, 0)
	level := make([]int, 0)

	for n != nil && queue.Length() > 0 {
		for range queue.Length() {
			nPtr := queue.Pop()

			if nPtr == nil {
				break
			}

			n := *nPtr

			level = append(level, n.value)

			if n.left != nil {
				queue.Push(n.left)
			}

			if n.right != nil {
				queue.Push(n.right)
			}
		}

		levels = append(levels, level)
		level = make([]int, 0)
	}

	// for _, level := range levels {
	// 	log.Printf("%#+v", level)
	// }

	slices.SortFunc(levels, func(a []int, b []int) int {
		if len(a) < len(b) {
			return -1
		} else if len(a) > len(b) {
			return 1
		} else {
			return 0
		}
	})

	return levels[len(levels)-1]
}

func InvertBinaryTree(array [1024]int) {
	t := NewBinaryTree(array[0])

	for i := 1; i < 1024; i++ {
		t.Add(array[i])
	}

	var n = t

	stack := NewStack[*BinaryTree]()

	stack.Push(n)

	for n != nil {
		nPtr := stack.Pop()

		if nPtr == nil {
			break
		}

		n := *nPtr

		n.left, n.right = n.right, n.left

		if n.left != nil {
			stack.Push(n.left)
		}

		if n.right != nil {
			stack.Push(n.right)
		}
	}

	t.WriteToFile()
}
