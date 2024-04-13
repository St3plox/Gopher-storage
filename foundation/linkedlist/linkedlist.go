// Package linkedlist is used for implementation  circular doubly linked list mainly for storing nodes in hash space
package linkedlist

type ListNode[T any] struct {
	HasheId int
	Val     T
	Next    *ListNode[T]
	Prev    *ListNode[T]
}

type LinkedList[T any] struct {
	Head *ListNode[T]
	size uint
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) Insert(hashedId int, val T) {
	node := &ListNode[T]{
		HasheId: hashedId,
		Val:     val,
	}

	if list.Head == nil {
		list.Head = node
		node.Next = node
		node.Prev = node
		list.size++
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
	list.size++
}

func (list *LinkedList[T]) FindClosestNode(hashedID int) *ListNode[T] {
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

	if current.HasheId <= hashedID {
		closest = current
	}

	if closest.HasheId == hashedID {
		return closest
	}

	return closest.Next
}

func (list *LinkedList[T]) RemovedNode(hashedID int) {
	if list.Head == nil {
		return
	}
	current := list.Head

	for current.Next != list.Head {
		if current.HasheId == hashedID {
			prev := current.Prev
			next := current.Next

			prev.Next = next
			next.Prev = prev

			list.size--
			return
		}
		current = current.Next
	}
}

func (list *LinkedList[T]) Values() []T {
	if list.Head == nil {
		return nil
	}

	values := make([]T, 0, list.size)
	current := list.Head
	for {
		values = append(values, current.Val)
		current = current.Next
		if current == list.Head {
			break
		}
	}

	return values
}