package bst

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := &Node{2, "luffy",nil,nil}
	tree = Insert(tree, 8,"doflamingo")
	exist := Lookup(tree,4)
	if exist {
		t.Fail()
	}
	tree = Insert(tree, 3,"shanks")
	tree = Insert(tree, 4,"graybeard")
	tree = Insert(tree, 4,"blackbeard")
	exist = Lookup(tree,4)
	if !exist {
		t.Fail()
	}
}
