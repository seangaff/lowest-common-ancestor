//Go implementation of LCA of two nodes on a binary tree.
//Code is adapted from compilethecode.com/
package lca

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func inorderRecursive(root *Node) {
	if root == nil {
		return
	}
	inorderRecursive(root.left)
	fmt.Printf("%d \n", root.key)
	inorderRecursive(root.right)
}

func leastCommonAncestorInBST(root *Node, n1 int, n2 int) int {
	if root == nil {
		return -1
	}
	if n1 > root.key && n2 > root.key {
		return leastCommonAncestorInBST(root.right, n1, n2)
	}
	if n1 < root.key && n2 < root.key {
		return leastCommonAncestorInBST(root.left, n1, n2)
	}
	return root.key
}

//test LCA functions - Original
/*
func main() {
	root := Node{100, nil, nil}
	root.left = &Node{50, nil, nil}
	root.right = &Node{150, nil, nil}
	root.left.left = &Node{20, nil, nil}
	root.left.left.left = &Node{10, nil, nil}
	root.left.left.right = &Node{30, nil, nil}
	root.left.right = &Node{80, nil, nil}
	root.right.left = &Node{110, nil, nil}
	root.right.right = &Node{200, nil, nil}

	fmt.Println("Inorder Traversal: ")
	inorderRecursive(&root)
	var result = leastCommonAncestorInBST(&root, 10, 80)
	fmt.Println("LAC is", result)

}
*/
