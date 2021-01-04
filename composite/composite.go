package composite

import "fmt"

const (
	// Leaf node
	Leaf = iota
	// NonLeaf non-leaf node
	NonLeaf
)

// TreeNode ...
type TreeNode interface {
	SetParent(TreeNode)
	AddChild(TreeNode)
	GetName() string
	Print(string)
}

type treeNode struct {
	name       string
	parent     TreeNode
	children   []TreeNode
	isLeafNode bool
}

// NewTreeNode ...
func NewTreeNode(name string, nodeKind int) TreeNode {
	var isLeafNode bool
	if nodeKind == Leaf {
		isLeafNode = true
	}

	return &treeNode{
		name:       name,
		isLeafNode: isLeafNode,
	}
}

func (c *treeNode) SetParent(parent TreeNode) {
	c.parent = parent
}

func (c *treeNode) AddChild(child TreeNode) {
	if c.isLeafNode {
		fmt.Println("Warning: leaf node cannot add child")
		return
	}
	child.SetParent(c)
	c.children = append(c.children, child)
}

func (c *treeNode) GetName() string {
	return c.name
}

func (c *treeNode) Print(pre string) {
	if c.isLeafNode {
		fmt.Printf("%s-%s\n", pre, c.GetName())
		return
	}

	fmt.Printf("%s+%s\n", pre, c.GetName())
	pre += "  "
	for _, child := range c.children {
		child.Print(pre)
	}
}
