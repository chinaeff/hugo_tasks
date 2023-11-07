package binary

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {
	return "graph TD\n" + toMermaid(t.Root)
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *Node) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func leftRotate(x *Node) *Node {
	y := x.Right
	x.Right = y.Left
	y.Left = x

	updateHeight(x)
	updateHeight(y)

	return y
}

func rightRotate(y *Node) *Node {
	x := y.Left
	y.Left = x.Right
	x.Right = y

	updateHeight(y)
	updateHeight(x)

	return x
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		// Duplicate key insertion is not allowed
		return node
	}

	updateHeight(node)

	balance := getBalance(node)

	if balance > 1 {
		if key < node.Left.Key {
			return rightRotate(node)
		}
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 {
		if key > node.Right.Key {
			return leftRotate(node)
		}
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func toMermaid(node *Node) string {
	if node == nil {
		return ""
	}

	mermaidStr := fmt.Sprintf("%d --> ", node.Key)
	if node.Left != nil {
		mermaidStr += fmt.Sprintf("%d\n", node.Left.Key)
	} else {
		mermaidStr += "null\n"
	}

	if node.Right != nil {
		mermaidStr += fmt.Sprintf("%d\n", node.Right.Key)
	} else {
		mermaidStr += "null\n"
	}

	return mermaidStr + toMermaid(node.Left) + toMermaid(node.Right)
}

func GenerateTree(count int) *AVLTree {
	rand.Seed(time.Now().UnixNano())
	tree := &AVLTree{}

	for i := 0; i < count; i++ {
		key := rand.Intn(1000)
		tree.Insert(key)
	}

	return tree
}

func main() {
	tree := GenerateTree(5)
	fmt.Println(tree.ToMermaid())

	go func() {
		for {
			time.Sleep(5 * time.Second)
			key := rand.Intn(1000)
			tree.Insert(key)
			fmt.Println(tree.ToMermaid())
		}
	}()

	select {}
}
