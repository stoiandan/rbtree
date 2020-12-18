package models

import (
	"fmt"
	"rbtree/utils"
)

//Node represents a binary tree node
type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Data   int
	color  color
}

//InsertData adds data as a node into the subtree
func (n *Node) InsertData(data int) {
	if data >= n.Data {
		if n.Right != nil {
			n.Right.InsertData(data)
		} else {
			n.Right = NewNode(data)
			n.Right.Parent = n
		}
	} else {
		if n.Left != nil {
			n.Left.InsertData(data)
		} else {
			n.Left = NewNode(data)
			n.Left.Parent = n
		}
	}
}

func (n *Node) InsertDataList(datalist []int) {
	for _, data := range datalist {
		n.InsertData(data)
	}
}

//NewNode creates a treee node
func NewNode(data int) *Node {
	return &Node{
		Data:   data,
		Left:   nil,
		Parent: nil,
		Right:  nil,
		color:  RED,
	}
}

func (n *Node) leftRotate() {
	pivot := n.Right
	if pivot == nil {
		return
	}

	pivot.Parent = n.Parent
	if n.Parent != nil {
		if n.Parent.Right == n {
			n.Parent.Right = pivot
		} else {
			n.Parent.Left = pivot
		}
	}

	n.Right = pivot.Left
	if pivot.Left != nil {
		pivot.Left.Parent = n
	}

	pivot.Left = n
	n.Parent = pivot
}

func (n *Node) rightRotate() {
	pivot := n.Left
	if pivot == nil {
		return
	}

	pivot.Parent = n.Parent
	if n.Parent != nil {
		if n.Parent.Right == n {
			n.Parent.Right = pivot
		} else {
			n.Parent.Left = pivot
		}
	}

	n.Left = pivot.Right
	if pivot.Right != nil {
		pivot.Right.Parent = n
	}

	n.Parent = pivot
	pivot.Right = n
}

func (n Node) String() string {
	str := ""
	if n.Left != nil {
		str = fmt.Sprintf("%v %v", n.Left.String(), n.Data)
	} else {
		str = fmt.Sprint(n.Data)
	}
	if n.Right != nil {
		str = fmt.Sprintf("%v %v", str, n.Right.String())
	}
	return str
}

// ArrayWithDepth helper method that returns the data sorted with the id number in paranthesis, root has 1, left has 2 and so on and so forth, missing nodes are still increasing the id
func (n Node) ArrayWithDepth() []string {
	return n.privateArrayWithDepth("1")
}

// privateArrayWithDepth private helper method used for ArrayWithDepth
func (n Node) privateArrayWithDepth(parentDepth string) []string {
	left := []string{}
	right := []string{}
	if n.Left != nil {
		left = n.Left.privateArrayWithDepth(parentDepth + "0")
	}
	if n.Right != nil {
		right = n.Right.privateArrayWithDepth(parentDepth + "1")
	}
	return append(append(left, fmt.Sprintf("%v(%v)", n.Data, utils.ConvertBinaryToDecimal(parentDepth))), right...)
}

func (n Node) PrettyDraw() string {
	n.ArrayWithDepth()
	return "not done"
}
