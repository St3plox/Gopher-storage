// Package linkedlist is used for implementation  circular doubly linked list mainly for storing nodes in hash space 
package linkedlist

type ListNode struct {
	HasheId int
	Val     any
	Next    *ListNode
	Prev    *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (list *LinkedList) Insert(hashedId int, val any) {
	
	node := &ListNode{
		HasheId: hashedId,
		Val: val,
	}
	
	if list.Head == nil {
		list.Head = node
		node.Next = node
		node.Prev = node
		return
	}
	
	current := list.Head
	for current.Next != list.Head && current.Next.HasheId < node.HasheId {
		current = current.Next
	}
	
	node.Next = current.Next
	node.Prev = current
	current.Next.Prev = node
	current.Next = node
}

func (list *LinkedList) FindClosestNode(hashedID int) *ListNode {
	
	if list.Head == nil {
		return nil
	}
	
	current := list.Head
	closest := current
	
	for current.Next != list.Head {
		if current.HasheId <= hashedID {
			closest = current
		}
		current = current.Next
	}
	
	// Check the last node
	if current.HasheId <= hashedID {
		closest = current
	}
	
	return closest
}
