// Merge potentially large amounts of possibly overlapping integer intervals.
package imerge

import "fmt"

// Node represents the top of a B-tree holding sorted non-overlapping integer intervals.
type Node struct {
	start int
	end   int
	left  *Node
	right *Node
}

// Create a new instance of Node holding a single integer interval.
func NewNode(start, end int) (*Node, error) {
	var err error
	if start > end {
		return nil, fmt.Errorf("Ignoring invalid input: %d > %d", start, end)
	}
	return &Node{start: start, end: end}, err
}

// Extract the interval hel by Node into a slice.
func (n *Node) toSlice() []int {
	return []int{n.start, n.end}
}

// Outputs the intervals
func (n *Node) Intervals() [][]int {
	l := [][]int{}
	if n.left != nil {
		l = append(l, n.left.Intervals()...)
	}
	l = append(l, n.toSlice())
	if n.right != nil {
		l = append(l, n.right.Intervals()...)
	}
	return l
}

// Inserts a new interval into the tree or merges it.
func (n *Node) Merge(start, end int) error {
	var err error
	if start > end {
		return fmt.Errorf("Ignoring invalid input: %d > %d", start, end)
	}
	if n.includes(start, end) {
		return err
	}

	if n.canMerge(start, end) {
		if start < n.start {
			n.left, n.start = newLeft(n.left, start)
		}
		if end > n.end {
			n.right, n.end = newRight(n.right, end)
		}
		return err
	}

	child := &n.left
	if start > n.start {
		child = &n.right
	}

	if *child != nil {
		(*child).Merge(start, end)
		return err
	}
	*child = &Node{start: start, end: end}
	return err
}

// Rebuilds the left side of a tree if merge occurred
func newLeft(n *Node, left int) (*Node, int) {
	if n == nil {
		return n, left
	}
	if left < n.start {
		return newLeft(n.left, left)
	}
	if left <= n.end {
		return n.left, n.start
	}
	n.right, left = newLeft(n.right, left)
	return n, left
}

// Rebuilds the right side of a tree if merge occurred
func newRight(n *Node, right int) (*Node, int) {
	if n == nil {
		return n, right
	}
	if right > n.end {
		return newRight(n.right, right)
	}
	if right >= n.start {
		return n.right, n.end
	}
	n.left, right = newRight(n.left, right)
	return n, right
}

func (n *Node) canMerge(start, end int) bool {
	if start > n.end || end < n.start {
		return false
	}
	return true
}

func (n *Node) includes(start, end int) bool {
	if start >= n.start && end <= n.end {
		return true
	}
	return false
}
