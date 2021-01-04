package composite

func ExampleTreeNode() {
	root := NewTreeNode("root", NonLeaf)

	nl1 := NewTreeNode("nl1", NonLeaf)
	nl2 := NewTreeNode("nl2", NonLeaf)
	nl3 := NewTreeNode("nl3", NonLeaf)

	l1 := NewTreeNode("l1", Leaf)
	l2 := NewTreeNode("l2", Leaf)
	l3 := NewTreeNode("l3", Leaf)

	root.AddChild(nl1)
	root.AddChild(nl2)

	nl1.AddChild(nl3)
	nl3.AddChild(l1)

	nl2.AddChild(l1)
	nl2.AddChild(l2)
	nl2.AddChild(l3)

	l1.AddChild(l2)

	root.Print("")

	// Output:
	// Warning: leaf node cannot add child
	// +root
	//   +nl1
	//     +nl3
	//       -l1
	//   +nl2
	//     -l1
	//     -l2
	//     -l3
}
