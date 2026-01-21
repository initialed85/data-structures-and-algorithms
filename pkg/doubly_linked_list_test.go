package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	require.Equal(t, []int{}, l.Items())

	l.Push(1)
	require.Equal(t, []int{1}, l.Items())

	l.Push(2)
	require.Equal(t, []int{1, 2}, l.Items())

	l.Push(3)
	require.Equal(t, []int{1, 2, 3}, l.Items())

	i1 := l.Pop()
	require.NotNil(t, i1)
	require.Equal(t, 3, *i1)
	require.Equal(t, []int{1, 2}, l.Items())

	i2 := l.Pop()
	require.NotNil(t, i2)
	require.Equal(t, 2, *i2)
	require.Equal(t, []int{1}, l.Items())

	i3 := l.Pop()
	require.NotNil(t, i3)
	require.Equal(t, 1, *i3)
	require.Equal(t, []int{}, l.Items())

	i4 := l.Pop()
	require.Nil(t, i4)
	require.Equal(t, []int{}, l.Items())

	l.PushLeft(1)
	require.Equal(t, []int{1}, l.Items())

	l.PushLeft(2)
	require.Equal(t, []int{2, 1}, l.Items())

	l.PushLeft(3)
	require.Equal(t, []int{3, 2, 1}, l.Items())

	i5 := l.PopLeft()
	require.NotNil(t, i5)
	require.Equal(t, 3, *i5)
	require.Equal(t, []int{2, 1}, l.Items())

	i6 := l.PopLeft()
	require.NotNil(t, i6)
	require.Equal(t, 2, *i6)
	require.Equal(t, []int{1}, l.Items())

	i7 := l.PopLeft()
	require.NotNil(t, i7)
	require.Equal(t, 1, *i7)
	require.Equal(t, []int{}, l.Items())

	i8 := l.PopLeft()
	require.Nil(t, i8)
	require.Equal(t, []int{}, l.Items())
}
