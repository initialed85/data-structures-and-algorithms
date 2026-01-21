package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	require.Equal(t, []int{}, s.Items())

	s.Push(1)
	require.Equal(t, []int{1}, s.Items())

	s.Push(2)
	require.Equal(t, []int{1, 2}, s.Items())

	s.Push(3)
	require.Equal(t, []int{1, 2, 3}, s.Items())

	i1 := s.Pop()
	require.NotNil(t, i1)
	require.Equal(t, 3, *i1)
	require.Equal(t, []int{1, 2}, s.Items())

	i2 := s.Pop()
	require.NotNil(t, i2)
	require.Equal(t, 2, *i2)
	require.Equal(t, []int{1}, s.Items())

	i3 := s.Pop()
	require.NotNil(t, i3)
	require.Equal(t, 1, *i3)
	require.Equal(t, []int{}, s.Items())

	i4 := s.Pop()
	require.Nil(t, i4)
	require.Equal(t, []int{}, s.Items())
}
