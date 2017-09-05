package bst

// todo: rewrite this in idiomatic go or just go and write everything in haskell
// todo: add delete def

type Node struct {
	Key int
	Value string
	Left *Node
	Right *Node
}

// the most interesting case is left for better times
// i'm of course talking about implementing the key equality aka key duplication
// there are three main solution
// 1. rebuild the tree - which might be expensive
// 2. make key field to be an array - which is pretty obvious solution
// 3. allow duplicates by making one of directions >= or >= - which would eliminate the search efficiency
// 4. maybe just replace the node with duplicate key with new one thereby erasing the portion of the tree -_-
// 5. do nothin' as i did -_- and return the node unchanged

func Insert(node *Node, key int, value string) *Node {
	if node == nil {
		return &Node{key, value, nil, nil}
	} else {
		if node.Key == key {
			return node
		} else {
			if node.Key > key {
				return &Node{node.Key,node.Value,Insert(node.Left, key, value), node.Right}
			} else { // node.Key < key
				return &Node{node.Key,node.Value,node.Left,Insert(node.Right, key, value)}
			}
		}
	}
}

// other names: search, find, contains
func Lookup(node *Node, key int) bool {
	if node == nil {
		return false
	} else {
		if node.Key == key {
			return true
		} else {
			if node.Key > key {
				return Lookup(node.Left, key)
			} else {
				return Lookup(node.Right, key)
			}
		}
	}
}

//func Delete(node *Node, key int) *Node {
//	if node.Key == key {
//		return nil
//	} else {
//		if node.Key > key {
//			return &Node{node.Key,node.Value,Delete(node.Left, key), node.Left }
//		} else {
//			return &Node{node.Key,node.Value,node.Left, Delete(node.Right, key) }
//		}
//	}
//}
