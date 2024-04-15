// Package linkedlist is used for implementation  circular doubly linked list mainly for storing nodes in hash space
package linkedlist

type ListNode[T any] struct {
	hashedID int
	Val      *T
	Next     *ListNode[T]
	Prev     *ListNode[T]
}

type LinkedList[T any] struct {
	Head *ListNode[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) Insert(hashedId int, val T) {
	node := &ListNode[T]{
		hashedID: hashedId,
		Val:      &val,
	}

	if list.Head == nil {
		list.Head = node
		node.Next = node
		node.Prev = node
		list.size++
		return
	}

	current := list.Head
	for current.Next != list.Head && current.Next.hashedID < node.hashedID {
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
		if current.hashedID <= hashedID {
			closest = current
		}
		current = current.Next
	}

	if current.hashedID <= hashedID {
		closest = current
	}

	if closest.hashedID == hashedID {
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
		if current.hashedID == hashedID {
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

func (ln ListNode[T]) HashedID() int {
	return ln.hashedID
}

func (list *LinkedList[T]) Size() int {
	return list.size
}
