package models

import "fmt"

//Node represents a binary tree node
type Node struct {
	Left  *Node
	Right *Node
	Data  int
	color color
}

//InsertData adds data as a node into the subtree
func (n *Node) InsertData(data int) {
	if data >= n.Data {
		if n.Right != nil {
			n.Right.InsertData(data)
		} else {
			n.Right = NewNode(data)
		}
	} else {
		if n.Left != nil {
			n.Left.InsertData(data)
		} else {
			n.Left = NewNode(data)
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
		Data:  data,
		Left:  nil,
		Right: nil,
		color: RED,
	}
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

func (n Node) PrettyDraw() string {
	return "not done"
}
