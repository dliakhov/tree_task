package tree

type Tree struct {
	Root *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// IsSymetric returns true is the binary tree given is symetric
// an example of a symetric binary looks like:
//        5
//       / \
//      7   7
//     / \ / \
//    3  4 4  3
//   / \     / \
//  1   6   6   1
func IsSymetric(root *Tree) bool {
	return isNodeSymetric(root.Root.Left, root.Root.Right)
}

func isNodeSymetric(left, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if !isNodesEqual(left, right) {
		return false
	}
	isLeftSymetric := isNodeSymetric(left.Left, right.Right)
	isRightSymetric := isNodeSymetric(left.Right, right.Left)
	return isLeftSymetric && isRightSymetric
}

func isNodesEqual(n1, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}
	return n1.Val == n2.Val
}
