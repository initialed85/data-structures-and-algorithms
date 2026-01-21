package pkg

/*
The stack is a LIFO data structure with O(1) / constant time complexity to push or pop and O(n) / linear time
complexity to iterate
*/

type Stack[T any] struct {
	doublyLinkedList *DoublyLinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		doublyLinkedList: NewDoublyLinkedList[T](),
	}
}

func (s *Stack[T]) Push(value T) {
	s.doublyLinkedList.Push(value)
}

func (s *Stack[T]) Pop() *T {
	return s.doublyLinkedList.Pop()
}

func (s *Stack[T]) Length() int {
	return s.doublyLinkedList.Length()
}

func (s *Stack[T]) Items() []T {
	return s.doublyLinkedList.Items()
}
