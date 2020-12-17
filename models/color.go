package models

type color bool

const (
	//RED color of a node
	RED color = true
	//BLACK color of a node
	BLACK color = false
)

func (c color) String() string {
	switch c {
	case BLACK:
		return "Black"
	case RED:
		return "Red"
	default:
		return "Unkown"
	}
}
