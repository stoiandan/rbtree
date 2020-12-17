package models

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

//NewNode creates a treee node
func NewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
		color: RED,
	}
}
