package imerge

import (
	"testing"
)

func TestTree(t *testing.T) {
	n := makeTree(input)
	got := n.Intervals()
	if !isIntervalsEqual(inputOrdered, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", input, got)
	}
}

func TestInvalidIntervalInit(t *testing.T) {
	_, err := NewNode(2, 1)
	if err == nil {
		t.Errorf("Should produce error for invalid interval")
	}
}

func TestInvalidIntervalInsert(t *testing.T) {
	n, _ := NewNode(1, 2)
	err := n.Merge(2, 1)
	if err == nil {
		t.Errorf("Should produce error for invalid interval")
	}
}

func TestNoMergeMerge(t *testing.T) {
	n := makeTree(input)
	n.Merge(352, 354)
	expect := append([][]int{}, inputOrdered[:20]...)
	expect = append(expect, []int{352, 354})
	expect = append(expect, inputOrdered[20:]...)
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func TestFullySwallowed(t *testing.T) {
	n := makeTree(input)
	n.Merge(28, 440)
	expect := [][]int{{28, 440}}
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func TestLeftLeafUpdateRight(t *testing.T) {
	n := makeTree(input)
	n.Merge(31, 34)
	expect := append([][]int{}, []int{30, 34})
	expect = append(expect, inputOrdered[1:]...)
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func TestLeftLeafUpdateBoth(t *testing.T) {
	n := makeTree(input)
	n.Merge(29, 34)
	expect := append([][]int{}, []int{29, 34})
	expect = append(expect, inputOrdered[1:]...)
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func TestOuterMiddle(t *testing.T) {
	n := makeTree(input)
	n.Merge(34, 225)
	expect := append([][]int{}, inputOrdered[0])
	expect = append(expect, []int{34, 225})
	expect = append(expect, inputOrdered[16:]...)
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func TestInnerMiddle(t *testing.T) {
	n := makeTree(input)
	n.Merge(61, 141)
	expect := append([][]int{}, inputOrdered[:4]...)
	expect = append(expect, []int{60, 144})
	expect = append(expect, inputOrdered[10:]...)
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func TestInnerMiddleBoundary(t *testing.T) {
	n := makeTree(input)
	n.Merge(73, 150)
	expect := append([][]int{}, inputOrdered[:5]...)
	expect = append(expect, []int{70, 155})
	expect = append(expect, inputOrdered[12:]...)
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func TestOuterMiddleBoundary(t *testing.T) {
	n := makeTree(input)
	n.Merge(40, 242)
	expect := append([][]int{}, inputOrdered[:1]...)
	expect = append(expect, []int{40, 242})
	expect = append(expect, inputOrdered[17:]...)
	got := n.Intervals()
	if !isIntervalsEqual(expect, got) {
		t.Errorf("Expected:\n%v, Got:\n%v", expect, got)
	}
}

func makeTree([][]int) *Node {
	n, err := NewNode(input[0][0], input[0][1])
	if err != nil {
		panic(err)
	}
	for i := 1; i < len(input); i++ {
		err := n.Merge(input[i][0], input[i][1])
		if err != nil {
			panic(err)
		}
	}
	return n
}

func isIntervalsEqual(a, b [][]int) bool {
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] == nil || b[i] == nil || len(a[i]) != 2 || len(b[i]) != 2 {
			return false
		}
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}
