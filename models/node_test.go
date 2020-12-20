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

func TestLeftRotate(t *testing.T) {
	errorMessages := ""

	errorMessages += testLeftRotate(
		5,
		[]int{4, 3, 2},
		"[2(8) 3(4) 4(2) 5(1)]",
	)
	errorMessages += testLeftRotate(
		5,
		[]int{6, 7, 8},
		"[5(2) 6(1) 7(3) 8(7)]",
	)
	errorMessages += testLeftRotate(
		5,
		[]int{3, 8, -3, 4, 7, 10},
		"[-3(8) 3(4) 4(9) 5(2) 7(5) 8(1) 10(3)]",
	)
	errorMessages += testLeftRotateChild(
		5,
		[]int{4, 3, 2},
		[]bool{false},
		"[2(8) 3(4) 4(2) 5(1)]",
	)
	errorMessages += testLeftRotateChild(
		5,
		[]int{6, 7, 8},
		[]bool{true},
		"[5(1) 6(6) 7(3) 8(7)]",
	)
	errorMessages += testLeftRotateChild(
		5,
		[]int{3, 8, -3, 4, 7, 10},
		[]bool{false},
		"[-3(8) 3(4) 4(2) 5(1) 7(6) 8(3) 10(7)]",
	)
	errorMessages += testLeftRotateChild(
		5,
		[]int{3, 8, -3, 4, 7, 10},
		[]bool{true},
		"[-3(4) 3(2) 4(5) 5(1) 7(12) 8(6) 10(3)]",
	)

	if errorMessages != "" {
		t.Error("\n" + errorMessages)
	}
}

func testLeftRotate(rootData int, dataList []int, expected string) string {
	return testRotateChild(false, rootData, dataList, []bool{}, expected)
}

func testLeftRotateChild(rootData int, dataList []int, childDirections []bool, expected string) string {
	return testRotateChild(false, rootData, dataList, childDirections, expected)
}

func testRotateChild(rotateRight bool, rootData int, dataList []int, childDirections []bool, expected string) string {
	root := NewNode(rootData)
	root.InsertDataList(dataList)
	before := fmt.Sprintf("%v", root.ArrayWithDepth())
	temp := root
	var rotate string
	for _, childDirection := range childDirections {
		if childDirection {
			temp = root.Right
		} else {
			temp = root.Left
		}
	}
	if rotateRight {
		temp.rightRotate()
		rotate = "Right"
	} else {
		temp.leftRotate()
		rotate = "Left"
	}
	if root.Parent != nil {
		root = root.Parent
	}
	got := fmt.Sprintf("%v", root.ArrayWithDepth())
	if got != expected {
		if len(childDirections) > 0 {
			return fmt.Sprintf("%v Rotate of child failed!\nChild:   \t%v\nBefore:  \t%v\nExpected:\t%v\nGot:     \t%v\n\n", rotate, temp.Data, before, expected, got)
		} else {
			return fmt.Sprintf("%v Rotate failed!\nBefore:  \t%v\nExpected:\t%v\nGot:     \t%v\n\n", rotate, before, expected, got)
		}
	}
	return ""
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
