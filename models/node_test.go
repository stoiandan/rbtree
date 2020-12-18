package models

import (
	"fmt"
	"testing"
)

func TestInsertingToRight(t *testing.T) {
	root := NewNode(5)
	root.InsertData(6)

	if root.Right == nil || root.Right.Data != 6 {
		t.Error("Inserting into tree fails")
	}
}

func TestStringOfNode(t *testing.T) {
	root := NewNode(5)
	root.InsertData(1)
	root.InsertData(12)
	root.InsertData(3)
	root.InsertData(10)
	root.InsertData(6)
	root.InsertData(-6)
	expected := "-6 1 3 5 6 10 12"
	if root.String() != expected {
		t.Errorf("The print was not correct!\nExpected:\t%v\nGot:     \t%v", expected, root.String())
	}
}

func TestStringOfNodeAndInsertDataList(t *testing.T) {
	root := NewNode(5)
	root.InsertDataList([]int{-6, 5, 3, 9, 8, 6, -9, 3, 2, 5, 8})

	expected := "-9 -6 2 3 3 5 5 5 6 8 8 9"
	if root.String() != expected {
		t.Errorf("The print was not correct!\nExpected:\t%v\nGot:     \t%v", expected, root.String())
	}
}
func TestArrayWithDepthOfNodeAndInsertDataList(t *testing.T) {
	root := NewNode(5)
	root.InsertDataList([]int{-6, 5, 3, 9, 8, 6, -9, 3, 2, 5, 8})

	expected := "[-9(4) -6(2) 2(10) 3(5) 3(11) 5(1) 5(3) 5(56) 6(28) 8(14) 8(29) 9(7)]"
	got := root.ArrayWithDepth()
	if fmt.Sprintf("%v", got) != expected {
		t.Errorf("The bubble print was not correct!\nExpected:\t%v\nGot:     \t%v", expected, got)
	}
}

func TestPrettyDrawOfNodeAndInsertDataList(t *testing.T) {
	root := NewNode(5)
	root.InsertDataList([]int{-6, 5, 3, 9, 8, 6, -9, 3, 2, 5, 8})
	//          5
	//   /             \
	//  -6	              5
	//  / \               \
	// -9  3              9
	//    /\             /
	//   2  3           8
	//                 / \
	// 	           6   8
	// 	          /
	// 	         5
}
