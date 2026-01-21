package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	require.Equal(t, []int{}, q.Items())

	q.Push(1)
	require.Equal(t, []int{1}, q.Items())

	q.Push(2)
	require.Equal(t, []int{1, 2}, q.Items())

	q.Push(3)
	require.Equal(t, []int{1, 2, 3}, q.Items())

	i1 := q.Pop()
	require.NotNil(t, i1)
	require.Equal(t, 1, *i1)
	require.Equal(t, []int{2, 3}, q.Items())

	i2 := q.Pop()
	require.NotNil(t, i2)
	require.Equal(t, 2, *i2)
	require.Equal(t, []int{3}, q.Items())

	i3 := q.Pop()
	require.NotNil(t, i3)
	require.Equal(t, 3, *i3)
	require.Equal(t, []int{}, q.Items())

	i4 := q.Pop()
	require.Nil(t, i4)
	require.Equal(t, []int{}, q.Items())
}
