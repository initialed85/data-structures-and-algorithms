package pkg

/*
- Linked List
  - Find the middle of singly-linked list (two pointer technique)
  - Reverse a singly-linked list
  - Rotate a singly-linked list
*/

type SinglyLinkedList struct {
	value int
	next  *SinglyLinkedList
}

func NewSinglyLinkedList(value int) *SinglyLinkedList {
	return &SinglyLinkedList{
		value: value,
		next:  nil,
	}
}

// Add adds the given value to the tail of the SinglyLinkedList
// and returns not the added value but the SingleLinkedList
// instance we called .Add() on (for use as a fluent interfacez)
func (l *SinglyLinkedList) Add(value int) *SinglyLinkedList {
	/* Iterating a singly linked list involves keeping a pointer
	to a current node and moving that to the next node until
	there are no more next nodes; so we do that to find the
	last node and then set the last nodes next as the node to
	be added

	Possible optimisation: keep a pointer to a pointer to the
	tail so adds go from O(n) to O(1)? */

	curr := l

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = NewSinglyLinkedList(value)

	return l
}

func FindMiddleOfSinglyLinkedList(head *SinglyLinkedList) int {
	/*
		Using the two pointer technique "fast and slow pointers",
		the fast pointer skips ahead by two nodes for as long
		as it can and the slow pointer skips ahead by one node-
		once the fast pointer can no longer skip ahead, the
		slow pointer will be at the middle-ish of the list
	*/

	fast := head
	slow := head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	return slow.value
}

func ReverseSinglyLinkedList(head *SinglyLinkedList) *SinglyLinkedList {
	/*
		Using three pointers to keep track of the previous,
		current and next nodes, we advance through the list,
		as we set current to current.next, we change current.next
		to to previous (which is naturally null for the head),
		finally we store current as previous before exiting
		the iteration

		It took me a little while to get my head around this
		one but it's actually simple; we step through like
		this:

		0 >> 1 >> 2 >> 3 >> ? (starting state)

		? << 0    1 >> 2 >> 3 >> ? (point head to null)

		? << 0 << 1    2 >> 3 >> ? (point 1 to head)

		? << 0 << 1 << 2    3 >> ? (point 2 to 1)

		? << 0 << 1 << 2 << 3      (point 3 to 2)

		Do note that during the interim states, the list is
		actually severed / partitioned in memory and it's
		only the fact that we're keeping the pointers to
		the fragments and reassembling as we go that causes
		us not to ruin it entirely
	*/

	var curr = head
	var next *SinglyLinkedList
	var prev *SinglyLinkedList

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func RotateSinglyLinkedList(head *SinglyLinkedList) *SinglyLinkedList {
	/*
		Similar to iterating the list, but we need track of
		the current node (which is eventually tail) as well
		as the previous node

		Once we've found the tail, we can point it to the
		head (which temporarily causes a cycle) until we
		point the previous node to null- finally, we return
		the tail (effectively making it the new head)

		So the three states look like this

		0 >> 1 >> 2 >> 3 >> ? (starting state)

		0 >> 1 >> 2 >> 3 >> 0 (temporary cycle as we point the tail to the head)

		0 >> 1 >> 2 >> ?      (new tail established by setting previous.next to null)
				               3 >> 0 (new head viewed in old order)

		3 >> 0 >> 1 >> 2 >> ? (list viewed from new head aka current)

		We can see that once we return current as the new head, the list
		has been rotated
	*/

	curr := head
	var prev *SinglyLinkedList

	for curr.next != nil {
		prev = curr
		curr = curr.next
	}

	curr.next = head
	prev.next = nil

	return curr
}
