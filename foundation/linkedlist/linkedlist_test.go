// Package linkedlist is used for implementation  circular doubly linked list mainly for storing nodes in hash space

package linkedlist

import (
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	list := New()

	// Insert elements into the list
	list.Insert(1, "A")
	list.Insert(3, "C")
	list.Insert(2, "B")

	if list.Head.Next.HasheId != 2 {
		t.Errorf("Expected Head node HasheId to be 2, got %d", list.Head.Next.HasheId)
	}
}

func TestLinkedList_FindClosestNode(t *testing.T) {
	list := New()

	list.Insert(1, "1")
	list.Insert(3, "3")
	list.Insert(50, "50")
	list.Insert(101, "101")
	list.Insert(99, "99")

	type args struct {
		hashedId int
	}
	tests := []struct {
		name     string
		list     *LinkedList
		args     args
		expected int
	}{
		{"Test 1 get precise mid", list, args{50}, 50},
		{"Test 2 get closest mid", list, args{25}, 50},
		{"Test 3 get first precise", list, args{1},1},
		{"Test 4 get second precise", list, args{3}, 3},
		{"Test 5 get last precise", list, args{101}, 101},
		{"Test 6 get closest high", list, args{100}, 101},
		{"Test 7 get closest low", list, args{2}, 3},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.FindClosestNode(tt.args.hashedId)
			if got.HasheId != tt.expected {
				t.Errorf("Expected Head node HasheId to be %d, got %d", tt.expected, got.HasheId)
			}
		})
	}
}
