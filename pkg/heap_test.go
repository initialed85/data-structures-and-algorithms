package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinHeap(t *testing.T) {
	h := new(MinHeap[int])

	_, err := h.Min()
	require.Error(t, err)

	h.Push(2)
	min, err := h.Min()
	require.NoError(t, err)
	require.Equal(t, 2, min)

	h.Push(3)
	min, err = h.Min()
	require.NoError(t, err)
	require.Equal(t, 2, min)

	h.Push(300)
	min, err = h.Min()
	require.NoError(t, err)
	require.Equal(t, 2, min)

	h.Push(1)
	min, err = h.Min()
	require.NoError(t, err)
	require.Equal(t, 1, min)

	h.Push(-100)
	min, err = h.Min()
	require.NoError(t, err)
	require.Equal(t, -100, min)

	min, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, -100, min)

	min, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 1, min)

	min, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 2, min)

	min, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 3, min)

	min, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 300, min)

	_, err = h.Pop()
	require.Error(t, err)
}

func TestMaxHeap(t *testing.T) {
	h := new(MaxHeap[int])

	_, err := h.Max()
	require.Error(t, err)

	h.Push(2)
	max, err := h.Max()
	require.NoError(t, err)
	require.Equal(t, 2, max)

	h.Push(3)
	max, err = h.Max()
	require.NoError(t, err)
	require.Equal(t, 3, max)

	h.Push(300)
	max, err = h.Max()
	require.NoError(t, err)
	require.Equal(t, 300, max)

	h.Push(1)
	max, err = h.Max()
	require.NoError(t, err)
	require.Equal(t, 300, max)

	h.Push(-100)
	max, err = h.Max()
	require.NoError(t, err)
	require.Equal(t, 300, max)

	max, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 300, max)

	max, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 3, max)

	max, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 2, max)

	max, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, 1, max)

	max, err = h.Pop()
	require.NoError(t, err)
	require.Equal(t, -100, max)

	_, err = h.Pop()
	require.Error(t, err)
}

func TestMinPriorityQueue(t *testing.T) {
	q := NewMinPriorityQueue[string, int]()

	q.Push("John", 4)
	q.Push("Luke", 3)
	q.Push("Mark", 2)
	q.Push("Matthew", 1)

	key, priority, err := q.Min()
	require.NoError(t, err)
	require.Equal(t, 1, priority)
	require.Equal(t, "Matthew", key)

	key, priority, err = q.Pop()
	require.NoError(t, err)
	require.Equal(t, 1, priority)
	require.Equal(t, "Matthew", key)

	key, priority, err = q.Min()
	require.NoError(t, err)
	require.Equal(t, 2, priority)
	require.Equal(t, "Mark", key)

	err = q.Adjust("Phil", 5)
	require.Error(t, err)

	err = q.Adjust("Mark", 5)
	require.NoError(t, err)

	key, priority, err = q.Pop()
	require.NoError(t, err)
	require.Equal(t, 3, priority)
	require.Equal(t, "Luke", key)

	key, priority, err = q.Pop()
	require.NoError(t, err)
	require.Equal(t, 4, priority)
	require.Equal(t, "John", key)

	q.Push("Phil", 6)
	err = q.Adjust("Phil", 1)
	require.NoError(t, err)

	key, priority, err = q.Pop()
	require.NoError(t, err)
	require.Equal(t, 1, priority)
	require.Equal(t, "Phil", key)

	key, priority, err = q.Pop()
	require.NoError(t, err)
	require.Equal(t, 5, priority)
	require.Equal(t, "Mark", key)

	_, _, err = q.Pop()
	require.Error(t, err)
}
