package models

import "testing"

func TestInsertingToRight(t *testing.T) {
	root := NewNode(5)
	root.InsertData(6)

	if root.Right == nil || root.Right.Data != 6 {
		t.Error("Inserting into tree fails")
	}
}
