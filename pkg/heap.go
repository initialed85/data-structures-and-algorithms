package pkg

import (
	"cmp"
	"fmt"
)

/*
A min heap is a somewhat ordered data structure with O(1) time complexity to get or pop the minimum value and between O(1)
and O(log n) time complexity to push a value (pending the size of the value in comparison to the existing values).
*/

type MinHeap[T cmp.Ordered] []T

func NewMinHeap[T cmp.Ordered]() *MinHeap[T] {
	h := make(MinHeap[T], 0)

	return &h
}

func (h *MinHeap[T]) Push(value T) {
	if len(*h) == 0 || value < (*h)[len(*h)-1] {
		(*h) = append((*h), value)
		return
	}

	(*h) = append((*h), value)

	for i := len(*h) - 1; i > 0; i-- {
		if (*h)[i] >= (*h)[i-1] {
			(*h)[i-1], (*h)[i] = (*h)[i], (*h)[i-1]
		}
	}
}

func (h *MinHeap[T]) Min() (T, error) {
	if len(*h) == 0 {
		return *new(T), fmt.Errorf("cannot get min for empty MinHeap")
	}

	return (*h)[len((*h))-1], nil
}

func (h *MinHeap[T]) Pop() (T, error) {
	if len(*h) == 0 {
		return *new(T), fmt.Errorf("cannot pop from empty MinHeap")
	}

	value := (*h)[len((*h))-1]

	*h = (*h)[:len((*h))-1]

	return value, nil
}

/*
A max heap is a somewhat ordered data structure with O(1) time complexity to get or pop the maximum value and between O(1)
and O(log n) time complexity to push a value (pending the size of the value in comparison to the existing values).
*/

type MaxHeap[T cmp.Ordered] []T

func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	h := make(MaxHeap[T], 0)

	return &h
}

func (h *MaxHeap[T]) Push(value T) {
	if len(*h) == 0 || value > (*h)[len(*h)-1] {
		(*h) = append((*h), value)
		return
	}

	(*h) = append((*h), value)

	for i := len(*h) - 1; i > 0; i-- {
		if (*h)[i] >= (*h)[i-1] {
			break
		}

		(*h)[i], (*h)[i-1] = (*h)[i-1], (*h)[i]
	}
}

func (h *MaxHeap[T]) Max() (T, error) {
	if len(*h) == 0 {
		return *new(T), fmt.Errorf("cannot get min for empty MaxHeap")
	}

	return (*h)[len((*h))-1], nil
}

func (h *MaxHeap[T]) Pop() (T, error) {
	if len(*h) == 0 {
		return *new(T), fmt.Errorf("cannot pop from empty MaxHeap")
	}

	value := (*h)[len((*h))-1]

	*h = (*h)[:len((*h))-1]

	return value, nil
}

/*
A min priority queue is very similar to a min heap, except the items have keys attached to them and once
in a min priority queue, an item can have its priority adjusted
*/

type PriorityItem[K comparable, P cmp.Ordered] struct {
	key      K
	priority P
	index    int
}

type MinPriorityQueue[K comparable, P cmp.Ordered] []PriorityItem[K, P]

func NewMinPriorityQueue[K comparable, P cmp.Ordered]() *MinPriorityQueue[K, P] {
	h := make(MinPriorityQueue[K, P], 0)

	return &h
}

func (h *MinPriorityQueue[K, P]) Push(key K, priority P) {
	if len(*h) == 0 || priority < (*h)[len(*h)-1].priority {
		(*h) = append((*h), PriorityItem[K, P]{
			priority: priority,
			key:      key,
		})

		(*h)[len(*h)-1].index = len(*h) - 1

		return
	}

	(*h) = append((*h), PriorityItem[K, P]{
		priority: priority,
		key:      key,
	})

	(*h)[len(*h)-1].index = len(*h) - 1

	for i := len(*h) - 1; i > 0; i-- {
		if (*h)[i].priority >= (*h)[i-1].priority {
			(*h)[i-1], (*h)[i] = (*h)[i], (*h)[i-1]
			(*h)[i-1].index += 1
			(*h)[i].index -= 1
		}
	}
}

func (h *MinPriorityQueue[K, P]) Min() (K, P, error) {
	if len(*h) == 0 {
		return *new(K), *new(P), fmt.Errorf("cannot get min for empty MinPriorityQueue")
	}

	item := (*h)[len((*h))-1]

	return item.key, item.priority, nil
}

func (h *MinPriorityQueue[K, P]) Pop() (K, P, error) {
	if len(*h) == 0 {
		return *new(K), *new(P), fmt.Errorf("cannot pop from empty MinHeap")
	}

	item := (*h)[len((*h))-1]

	*h = (*h)[:len((*h))-1]

	return item.key, item.priority, nil
}

func (h *MinPriorityQueue[K, P]) Adjust(key K, priority P) error {
	index := -1
	var item PriorityItem[K, P]

	for i, possibleItem := range *h {
		if possibleItem.key == key {
			index = i
			item = possibleItem
			break
		}
	}

	if index < 0 {
		return fmt.Errorf("cannot adjust priority for unknown key %#+v", key)
	}

	if priority < item.priority {
		(*h)[index].priority = priority

		for i := index; i < len(*h)-1; i++ {
			if (*h)[i].priority > (*h)[i+1].priority {
				break
			}

			(*h)[i], (*h)[i+1] = (*h)[i+1], (*h)[i]
			(*h)[i].index, (*h)[i+1].index = (*h)[i+1].index, (*h)[i].index
		}
	} else if priority > item.priority {
		(*h)[index].priority = priority

		for i := index; i > 0; i-- {
			if (*h)[i].priority < (*h)[i-1].priority {
				break
			}

			(*h)[i], (*h)[i-1] = (*h)[i-1], (*h)[i]
			(*h)[i].index, (*h)[i-1].index = (*h)[i-1].index, (*h)[i].index
		}
	}

	return nil
}
